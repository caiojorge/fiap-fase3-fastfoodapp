package robot

import (
	"context"

	"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/server"
	"go.uber.org/zap"
)

type INotifierRobot interface {
	Notify(ctx context.Context) error
}

type NotifierRobot struct {
	server *server.GinServer
	logger *zap.Logger
}

func NewNotifierRobot(server *server.GinServer, logger *zap.Logger) *NotifierRobot {
	return &NotifierRobot{
		server: server,
		logger: logger,
	}
}

func (n *NotifierRobot) Notify(ctx context.Context) error {
	output, err := n.server.GetContainer().NotifierKitchenUseCase.Notify(context.Background())
	if err != nil {
		n.logger.Error("Failed to notify kitchen", zap.Error(err))
		return err
	}

	n.logger.Info("Kitchen notified", zap.Any("output", output))

	return nil
}
