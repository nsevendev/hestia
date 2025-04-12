package closedperiod

import (
	"context"
	"fmt"
	"hestia/internal/models"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

func (s *closedPeriodService) Delete(ctx context.Context, uuidClosurePeriod string) error {
	uid, err := uuid.Parse(uuidClosurePeriod)
	if err != nil {
		return fmt.Errorf("UUID invalide : %w", err)
	}

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&models.ClosurePeriod{}, "uuid = ?", uid)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("aucune période de fermeture trouvée avec cet UUID")
		}
		return nil
	}); err != nil {
		return fmt.Errorf("impossible de supprimer la période de fermeture : %w", err)
	}

	return nil
}