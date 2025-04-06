package closedperiod

import (
	"context"
	"fmt"
	"hestia/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *closedPeriodService) Create(ctx context.Context, title string, startDate string, endDate string) error {
	startTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return err
	}

	endTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return err
	}

	if !endTime.After(startTime) {
		return fmt.Errorf("la date de fin doit être supérieure à la date de début")
	}

	var existing models.ClosurePeriod
	err = s.db.WithContext(ctx).
		Where("NOT (? > end_date OR ? < start_date)", startTime, endTime).
		Take(&existing).Error

	if err == nil {
		return fmt.Errorf("une période de fermeture existe déjà entre %s et %s", existing.StartDate.Format("2006-01-02"), existing.EndDate.Format("2006-01-02"))
	}

	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("erreur lors de la vérification des périodes existantes : %w", err)
	}

	closurePeriod := models.ClosurePeriod{
		UUID:      uuid.New(),
		Title:     title,
		StartDate: startTime,
		EndDate:   endTime,
	}

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&closurePeriod).Error
	}); err != nil {
		return fmt.Errorf("impossible de creer closure period: %w", err)
	}
	
	return nil
}