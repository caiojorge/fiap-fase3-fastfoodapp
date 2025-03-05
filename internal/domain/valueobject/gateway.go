package valueobject

type Gateway struct {
	GatewayName          string
	GatewayToken         string
	GatewayTransactionID string
}

func NewGateway(gatewayName string, gatewayToken string) Gateway {
	return Gateway{
		GatewayName:  gatewayName,
		GatewayToken: gatewayToken,
	}
}
