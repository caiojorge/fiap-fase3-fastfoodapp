package sharedurl

import "fmt"

type URL struct {
	hostname string
	hostport string
}

func NewURL(hostname, hostport string) *URL {
	return &URL{
		hostname: hostname,
		hostport: hostport,
	}
}

func (u *URL) GetPaymentURL(collectorID, posID string) string {
	return fmt.Sprintf("http://%s:%s/instore/orders/qr/seller/collectors/%s/pos/%s/qrs", u.hostname, u.hostport, collectorID, posID)
}

func (u *URL) GetWebhookURL() string {
	return fmt.Sprintf("http://%s:%s/kitchencontrol/api/v1/checkouts/confirmation/payment", u.hostname, u.hostport)
}
