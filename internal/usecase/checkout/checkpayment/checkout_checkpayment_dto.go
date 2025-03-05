package usecase

type CheckPaymentInputDTO struct{}
type CheckPaymentOutputDTO struct {
	OrderID              string
	Status               string
	GatewayTransactionID string
	PaymentApproved      bool
}
