package bonusreport

import (
	"context"
	"washbonus/internal/entities"
)

func (s *bonusReportService) List(ctx context.Context, filter entities.BonusReportFilter) (entities.Page[entities.BonusReport], error) {
	return s.bonusReportRepo.List(ctx, filter)
}
