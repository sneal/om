package commands

import (
	"fmt"
	"github.com/pivotal-cf/go-pivnet"
	pivnetlog "github.com/pivotal-cf/go-pivnet/logger"
	"github.com/pivotal-cf/pivnet-cli/filter"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func NewPivnetClient(
	logger pivnetlog.Logger,
	progressWriter io.Writer,
	factory PivnetFactory,
	token string,
) *pivnetClient {
	return &pivnetClient{
		filter:         filter.NewFilter(logger),
		progressWriter: progressWriter,
		downloader: factory(
			pivnet.ClientConfig{
				Host:              pivnet.DefaultHost,
				Token:             token,
				UserAgent:         fmt.Sprintf("om-download-product"),
				SkipSSLValidation: false,
			},
			logger,
		),
	}
}

type pivnetClient struct {
	downloader     PivnetDownloader
	filter         *filter.Filter
	progressWriter io.Writer
}

type fileArtifact struct {
	name          string
	sha256        string
	slug          string
	releaseID     int
	productFileID int
}

type stemcell struct {
	slug    string
	version string
}

func (p *pivnetClient) GetAllProductVersions(slug string) ([]string, error) {
	releases, err := p.downloader.ReleasesForProductSlug(slug)
	if err != nil {
		return nil, err
	}

	var versions []string
	for _, release := range releases {
		versions = append(versions, release.Version)
	}
	return versions, nil
}

func (p *pivnetClient) GetLatestProductFile(slug, version, glob string) (*fileArtifact, error) {
	// 1. Check the release for given version / slug
	release, err := p.downloader.ReleaseForVersion(slug, version)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the release for %s %s: %s", slug, version, err)
	}

	// 2. Get filename from pivnet
	productFiles, err := p.downloader.ProductFilesForRelease(slug, release.ID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the product files for %s %s: %s", slug, version, err)
	}

	productFiles, err = p.filter.ProductFileKeysByGlobs(productFiles, []string{glob})
	if err != nil {
		return nil, fmt.Errorf("could not glob product files: %s", err)
	}

	if err := p.checkForSingleProductFile(glob, productFiles); err != nil {
		return nil, err
	}

	return &fileArtifact{
		name:          productFiles[0].AWSObjectKey,
		sha256:        productFiles[0].SHA256,
		releaseID:     release.ID,
		slug:          slug,
		productFileID: productFiles[0].ID,
	}, nil
}

func (p *pivnetClient) checkForSingleProductFile(glob string, productFiles []pivnet.ProductFile) error {
	if len(productFiles) > 1 {
		var productFileNames []string
		for _, productFile := range productFiles {
			productFileNames = append(productFileNames, path.Base(productFile.AWSObjectKey))
		}
		return fmt.Errorf("the glob '%s' matches multiple files. Write your glob to match exactly one of the following:\n  %s", glob, strings.Join(productFileNames, "\n  "))
	} else if len(productFiles) == 0 {
		return fmt.Errorf("the glob '%s' matchs no file", glob)
	}

	return nil
}

func (p *pivnetClient) DownloadProductToFile(fa *fileArtifact, file *os.File) error {
	err := p.downloader.DownloadProductFile(file, fa.slug, fa.releaseID, fa.productFileID, p.progressWriter)
	if err != nil {
		return fmt.Errorf("could not download product file %s: %s", fa.slug, err)
	}
	return nil
}

func (p *pivnetClient) DownloadProductStemcell(fa *fileArtifact) (*stemcell, error) {
	dependencies, err := p.downloader.ReleaseDependencies(fa.slug, fa.releaseID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch stemcell dependency for %s: %s", fa.slug, err)
	}

	stemcellSlug, stemcellVersion, err := p.getLatestStemcell(dependencies)
	if err != nil {
		return nil, fmt.Errorf("could not sort stemcell dependency: %s", err)
	}

	return &stemcell{slug: stemcellSlug, version: stemcellVersion}, nil
}

func (p *pivnetClient) getLatestStemcell(dependencies []pivnet.ReleaseDependency) (string, string, error) {
	const errorForVersion = "versioning of stemcell dependency in unexpected format. the following version could not be parsed: %s"

	var stemcellSlug string
	var stemcellVersion string
	var stemcellVersionMajor int
	var stemcellVersionMinor int

	for _, dependency := range dependencies {
		if strings.Contains(dependency.Release.Product.Slug, "stemcells") {
			versionString := dependency.Release.Version
			splitVersions := strings.Split(versionString, ".")
			if len(splitVersions) == 1 {
				splitVersions = []string{splitVersions[0], "0"}
			}
			if len(splitVersions) != 2 {
				return stemcellSlug, stemcellVersion, fmt.Errorf(errorForVersion, versionString)
			}
			major, err := strconv.Atoi(splitVersions[0])
			if err != nil {
				return stemcellSlug, stemcellVersion, fmt.Errorf(errorForVersion, versionString)
			}
			minor, err := strconv.Atoi(splitVersions[1])
			if err != nil {
				return stemcellSlug, stemcellVersion, fmt.Errorf(errorForVersion, versionString)
			}
			if major > stemcellVersionMajor {
				stemcellVersionMajor = major
				stemcellVersionMinor = minor
				stemcellVersion = versionString
				stemcellSlug = dependency.Release.Product.Slug
			} else if major == stemcellVersionMajor && minor > stemcellVersionMinor {
				stemcellVersionMinor = minor
				stemcellVersion = versionString
				stemcellSlug = dependency.Release.Product.Slug
			}
		}
	}

	return stemcellSlug, stemcellVersion, nil
}
