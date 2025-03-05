package usecase

import (
	"context"
	"fmt"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type ProductDeleteUseCase struct {
	repository portsrepository.ProductRepository
}

func NewProductDelete(repository portsrepository.ProductRepository) *ProductDeleteUseCase {
	return &ProductDeleteUseCase{
		repository: repository,
	}
}

// DeleteProduct remove um novo produto.
func (cr *ProductDeleteUseCase) DeleteProduct(ctx context.Context, id string) error {

	_, err := cr.repository.Find(ctx, id)
	if err != nil && err.Error() != "product not found" {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	err = cr.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
