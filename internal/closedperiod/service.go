package closedperiod

import (
	"context"
	"hestia/internal/models"

	"gorm.io/gorm"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type closedPeriodService struct {
	db *gorm.DB
}	

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type ClosedPeriodService interface {
	Create(ctx context.Context, title string, startDate string, endDate string) error
	Delete(ctx context.Context, uuidClosurePeriod string) error
	Active(ctx context.Context) (*models.ClosurePeriod, error)
	List(ctx context.Context) ([]models.ClosurePeriod, error)
}

func NewClosedPeriodService(db *gorm.DB) ClosedPeriodService {
	return &closedPeriodService{
		db,
	}
}