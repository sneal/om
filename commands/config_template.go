package commands

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/pivotal-cf/jhanda"
	yaml "gopkg.in/yaml.v2"
)

type ConfigTemplate struct {
	logger  logger
	Options struct {
		Product string `long:"product"  short:"p"  required:"true" description:"path to product to generate config template for"`
	}
}

func NewConfigTemplate(logger logger) ConfigTemplate {
	return ConfigTemplate{
		logger: logger,
	}
}

func (ct ConfigTemplate) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "This command will export the current installation of the target Ops Manager.",
		ShortDescription: "exports the installation of the target Ops Manager",
		Flags:            ct.Options,
	}
}

func (ct ConfigTemplate) Execute(args []string) error {
	if _, err := jhanda.Parse(&ct.Options, args); err != nil {
		panic(err)
		// return fmt.Errorf("could not parse config-template flags flags: %s", err)
	}

	zipReader, err := zip.OpenReader(ct.Options.Product)
	if err != nil {
		panic(err)
	}

	defer zipReader.Close()

	for _, file := range zipReader.File {
		matched, _ := regexp.MatchString("metadata/.*\\.yml", file.Name)

		if matched {
			metadataFile, err := file.Open()
			if err != nil {
				panic(err)
			}

			contents, err := ioutil.ReadAll(metadataFile)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(contents))

			type metadata struct {
				FormTypes          []map[string]string      `yaml:"form_types"`
				PropertyBlueprints []map[string]interface{} `yaml:"property_blueprints"`
				JobTypes           []map[string]interface{} `yaml:"job_types"`
			}
			metadataStruct := metadata{}
			err = yaml.Unmarshal(contents, &metadataStruct)
			if err != nil {
				panic(err)
			}

			fmt.Println(contents)
		}
	}

	// unzip the file from the path (ct.Options.Product)

	// read the metadata into a struct/interface

	// find all the properties

	ct.logger.Printf("finished generating config template")

	return nil
}
