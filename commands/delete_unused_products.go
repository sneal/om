package commands

import (
	"github.com/pivotal-cf/om/api"
)

type DeleteUnusedProducts struct {
	service deleteUnusedProductsService
	logger  logger
}

//counterfeiter:generate -o ./fakes/delete_unused_products_service.go --fake-name DeleteUnusedProductsService . deleteUnusedProductsService
type deleteUnusedProductsService interface {
	DeleteAvailableProducts(input api.DeleteAvailableProductsInput) error
}

func NewDeleteUnusedProducts(service deleteUnusedProductsService, logger logger) *DeleteUnusedProducts {
	return &DeleteUnusedProducts{
		service: service,
		logger:  logger,
	}
}

func (dup DeleteUnusedProducts) Execute(args []string) error {
	dup.logger.Printf("trashing unused products")

	err := dup.service.DeleteAvailableProducts(api.DeleteAvailableProductsInput{
		ShouldDeleteAllProducts: true,
	})
	if err != nil {
		return err
	}

	dup.logger.Printf("done")

	return nil
}
