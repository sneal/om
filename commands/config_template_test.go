package commands_test

import (
	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConfigTemplate", func() {
	var (
		logger *fakes.Logger
	)

	BeforeEach(func() {
		logger = &fakes.Logger{}
	})

	It("exports the installation", func() {
		command := commands.NewConfigTemplate(logger)

		err := command.Execute([]string{
			"--output-file", "/path/to/output.zip",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	FDescribe("Usage", func() {
		It("returns usage information for the command", func() {
			command := commands.NewExportInstallation(nil, nil)
			Expect(command.Usage()).To(Equal(jhanda.Usage{
				Description:      "This command generates a configuration template that can be passed in to om configure-product",
				ShortDescription: "generates a config template for the product",
				Flags:            command.Options,
			}))
		})
	})
})
