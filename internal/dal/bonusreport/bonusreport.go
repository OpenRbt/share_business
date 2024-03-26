package bonusreport

import (
	"context"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"

	"github.com/gocraft/dbr/v2"
)

func (r *repo) List(ctx context.Context, filter entities.BonusReportFilter) (entities.Page[entities.BonusReport], error) {
	query := r.db.NewSession(nil).
		Select("count(*)").
		From(dbr.I("balance_events").As("be")).
		Join(dbr.I("wallets").As("w"), "be.wallet_id = w.id").
		Join(dbr.I("organizations").As("o"), "w.organization_id = o.id")

	transactionFilterBuild(query, filter)

	var count int64
	err := query.LoadOneContext(ctx, &count)
	if err != nil {
		return entities.Page[entities.BonusReport]{}, err
	}

	query = r.db.NewSession(nil).
		Select(
			"be.id",
			"be.\"user\" as user_id",
			"ABS(be.new_amount - be.old_amount) as amount",
			"CASE WHEN be.new_amount - be.old_amount >= 0 THEN 'deposit' ELSE 'withdrawal' END as operation_type",
			"be.date",
			"o.id as org_id",
			"o.name as org_name",
			"o.deleted as org_deleted").
		From(dbr.I("balance_events").As("be")).
		Join(dbr.I("wallets").As("w"), "be.wallet_id = w.id").
		Join(dbr.I("organizations").As("o"), "w.organization_id = o.id").
		OrderDesc("be.date").
		Paginate(uint64(filter.Page()), uint64(filter.PageSize()))

	transactionFilterBuild(query, filter)

	var bonusReports []dbmodels.BonusReport
	_, err = query.LoadContext(ctx, &bonusReports)
	if err != nil {
		return entities.Page[entities.BonusReport]{}, err
	}

	return entities.NewPage(conversions.BonusReportsFromDB(bonusReports), filter.Filter, count), nil
}

func transactionFilterBuild(query *dbr.SelectStmt, filter entities.BonusReportFilter) {
	if filter.OrganizationID != nil {
		query.Where("organization_id = ?", filter.OrganizationID)
	}
	if filter.BalanceOperationType != nil {
		switch *filter.BalanceOperationType {
		case entities.DepositOperationType:
			query.Where("new_amount - old_amount >= 0")
		case entities.WithdrawalOperationType:
			query.Where("new_amount - old_amount < 0")
		default:
			panic("Unknown app operation: " + *filter.BalanceOperationType)
		}
	}
}
