package closedperiod

import (
	"context"
	"fmt"
	"hestia/internal/models"
	"time"

	"gorm.io/gorm"
)

func (s *closedPeriodService) Active(ctx context.Context) (*models.ClosurePeriod, error) {
	now := time.Now().Truncate(24 * time.Hour)

	var active models.ClosurePeriod
	err := s.db.WithContext(ctx).
		Where("? BETWEEN start_date AND end_date", now).
		Take(&active).Error

	if err == gorm.ErrRecordNotFound {
		// pas de period pas d'erreur
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération de la période de fermeture active : %w", err)
	}

	return &active, nil
}