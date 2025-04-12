package closedperiod

import (
	"context"
	"fmt"
	"hestia/internal/models"
)

func (s *closedPeriodService) List(ctx context.Context) ([]models.ClosurePeriod, error) {
	var periods []models.ClosurePeriod

	err := s.db.WithContext(ctx).
		Order("start_date ASC").
		Find(&periods).Error

	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des périodes de fermeture : %w", err)
	}

	return periods, nil
}