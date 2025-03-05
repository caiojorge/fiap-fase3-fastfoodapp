package usecase

import (
	"context"
)

type FindOrderByParamsUseCase interface {
	FindOrdersByParams(ctx context.Context, params map[string]interface{}) ([]*OrderFindByParamOutputDTO, error)
}
