package commands

import (
	"github.com/pivotal-cf/jhanda/commands"
)

type GenerateCertificate struct {
	service certificateGenerator
	logger  logger
}

//go:generate counterfeiter -o ./fakes/certificate_generator.go --fake-name CertificateGenerator . certificateGenerator
type certificateGenerator interface {
	Generate(string) (string, error)
}

func NewGenerateCertificate(service certificateGenerator, logger logger) GenerateCertificate {
	return GenerateCertificate{service: service, logger: logger}
}

func (g GenerateCertificate) Execute(args []string) error {
	output, err := g.service.Generate("blah")
	if err != nil {
		return err
	}

	g.logger.Printf(output)
	return nil
}

func (g GenerateCertificate) Usage() commands.Usage {
	return commands.Usage{
		Description:      "This authenticated command generates a certificate authority on the Ops Manager",
		ShortDescription: "generates a certificate authority on the Opsman",
	}
}
