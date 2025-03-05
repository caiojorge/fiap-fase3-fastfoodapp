package converter

import (
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/domain/valueobject"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driven/model"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
)

// TODO: voltar aqui para usar metodos ao invés de func: type CustomerConverter struct{}

func FromEntity(entity *entity.Customer) *model.Customer {

	cpfWithoutNonDigits := formatter.RemoveMaskFromCPF(entity.GetCPF().Value)

	return &model.Customer{
		CPF:   cpfWithoutNonDigits,
		Name:  entity.GetName(),
		Email: entity.GetEmail(),
	}
}

// ToEntity converte um model.Customer para um entity.Customer
// Não verifica se o CPF é válido
func ToEntity(model *model.Customer) *entity.Customer {
	// coloco a mascara no cpf qdo crio a entidade
	modelCPF, err := formatter.PutMaskOnCPF(model.CPF)
	if err != nil {
		modelCPF = model.CPF // se der erro, deixa sem mascara mesmo
	}

	cpf := valueobject.CPF{
		Value: modelCPF,
	}

	return &entity.Customer{
		CPF:   cpf,
		Name:  model.Name,
		Email: model.Email,
	}

}
