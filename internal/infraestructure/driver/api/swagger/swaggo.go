package swagger

import (
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/docs"
	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/server"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Swagger interface {
	Execute() string
}

type Swaggo struct {
	hostname string
	hostport string
	server   *server.GinServer
}

func NewSwaggo(hostname, hostport string, server *server.GinServer) Swagger {
	return &Swaggo{hostname: hostname, hostport: hostport, server: server}
}

func (s *Swaggo) Execute() string {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", s.hostname, s.hostport)
	docs.SwaggerInfo.BasePath = "/kitchencontrol/api/v1"

	swaggerURL := fmt.Sprintf("http://%s:%s/kitchencontrol/api/v1/docs/doc.json", s.hostname, s.hostport)
	s.server.GetRouter().GET("/kitchencontrol/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return swaggerURL
}
