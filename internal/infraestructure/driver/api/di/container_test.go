package di

import (
	"testing"

	//"github.com/caiojorge/fiap-challenge-ddd/internal/infraestructure/driver/api/di"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestContainer_Validate(t *testing.T) {
	db := &gorm.DB{}
	logger, _ := zap.NewProduction()

	t.Run("should work properly", func(t *testing.T) {
		container := NewContainer(db, logger)

		assert.NotNil(t, container)

		err := container.Validate()
		assert.Nil(t, err)
	})

	t.Run("should return error when db is nil", func(t *testing.T) {

		container := NewContainer(nil, logger)

		assert.NotNil(t, container)

		err := container.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, "db is not set", err.Error())

	})

	t.Run("should return error when logger is nil", func(t *testing.T) {
		container := NewContainer(db, nil)

		assert.NotNil(t, container)

		err := container.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, "logger is not set", err.Error())
	})

	t.Run("should return error when some object is nil", func(t *testing.T) {
		container := NewContainer(db, logger)
		container.Logger = nil

		assert.NotNil(t, container)

		err := container.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, "logger is not set", err.Error())
	})

}
