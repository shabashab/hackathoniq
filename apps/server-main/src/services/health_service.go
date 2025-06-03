package services

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthService struct {
	logger *zap.Logger
	db     *gorm.DB
}

// HealthStatus represents the health check response.
// swagger:model
type HealthStatus struct {
	// Whether the application is alive.
	Alive bool `json:"alive"`

	// Database connectivity status.
	Database bool `json:"database"`
}

func NewHealthService(logger *zap.Logger, db *gorm.DB) *HealthService {
	return &HealthService{
		logger: logger,
		db:     db,
	}
}

func (healthService HealthService) GetHealthStatus() HealthStatus {
	healthService.logger.Info("Health status requested")

	dbHealthStatus := healthService.getDatabaseHealthStatus()

	// for now, just return the database health status as the overall health status
	// in the future, we can add more checks

	aliveStatus := dbHealthStatus

	return HealthStatus{
		Alive:    aliveStatus,
		Database: dbHealthStatus,
	}
}

func (healthService HealthService) getDatabaseHealthStatus() bool {
	db, err := healthService.db.DB()

	if err != nil {
		healthService.logger.Error("database healthcheck failed", zap.Error(err))
		fmt.Println(err)
		return false
	}

	err = db.Ping()

	if err != nil {
		healthService.logger.Error("database healthcheck failed", zap.Error(err))
		return false
	}

	return true
}
