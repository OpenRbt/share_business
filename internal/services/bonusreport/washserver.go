package bonusreport

import (
	"washbonus/internal/app"

	"go.uber.org/zap"
)

type bonusReportService struct {
	logger          *zap.SugaredLogger
	bonusReportRepo app.BonusReportRepo
}

func New(l *zap.SugaredLogger, bonusReportRepo app.BonusReportRepo) app.BonusReportService {
	return &bonusReportService{
		logger:          l,
		bonusReportRepo: bonusReportRepo,
	}
}
