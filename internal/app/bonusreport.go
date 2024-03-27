package app

import (
	"washbonus/internal/entities"
)

type (
	BonusReportController interface {
		List(ctx Ctx, auth AdminAuth, filter entities.BonusReportFilter) (entities.Page[entities.BonusReport], error)
	}

	BonusReportService interface {
		List(ctx Ctx, filter entities.BonusReportFilter) (entities.Page[entities.BonusReport], error)
	}

	BonusReportRepo interface {
		List(ctx Ctx, filter entities.BonusReportFilter) (entities.Page[entities.BonusReport], error)
	}
)
