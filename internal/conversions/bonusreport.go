package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/models"
	"washbonus/openapi/admin/restapi/operations/reports"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func OperationTypeFromRest(operation *models.BalanceOperationType) *entities.BalanceOperationType {
	if operation == nil {
		return nil
	}

	switch *operation {
	case models.BalanceOperationTypeDeposit:
		o := entities.DepositOperationType
		return &o
	case models.BalanceOperationTypeWithdrawal:
		o := entities.WithdrawalOperationType
		return &o
	default:
		panic("Unknown app operation: " + *operation)
	}
}

func OperationTypeToRest(operation entities.BalanceOperationType) models.BalanceOperationType {
	switch operation {
	case entities.DepositOperationType:
		return models.BalanceOperationTypeDeposit
	case entities.WithdrawalOperationType:
		return models.BalanceOperationTypeWithdrawal
	default:
		panic("Unknown api operation: " + operation)
	}
}

func OperationTypeFromDB(operation dbmodels.BalanceOperationType) entities.BalanceOperationType {
	switch operation {
	case dbmodels.DepositOperationType:
		return entities.DepositOperationType
	case dbmodels.WithdrawalOperationType:
		return entities.WithdrawalOperationType
	default:
		panic("Unknown db operation: " + operation)
	}
}

func BonusReportFilterFromRest(params reports.GetBonusReportsParams) (entities.BonusReportFilter, error) {
	var organizationID *uuid.UUID
	if params.OrganizationID != nil {
		id, err := uuid.FromString(params.OrganizationID.String())
		if err != nil {
			return entities.BonusReportFilter{}, err
		}
		organizationID = &id
	}

	return entities.BonusReportFilter{
		Filter:               entities.NewFilter(*params.Page, *params.PageSize),
		OrganizationID:       organizationID,
		BalanceOperationType: OperationTypeFromRest((*models.BalanceOperationType)(params.Operation)),
	}, nil
}

func BonusReportToRest(report entities.BonusReport) models.Report {
	id := strfmt.UUID(report.ID.String())
	date := strfmt.DateTime(report.Date)
	operation := OperationTypeToRest(report.OperationType)
	orgId := strfmt.UUID(report.Organization.ID.String())
	amount := report.Amount.Shift(-report.Amount.Exponent()).IntPart()

	return models.Report{
		ID:        &id,
		Date:      &date,
		Operation: &operation,
		UserID:    &report.UserID,
		Amount:    &amount,
		Organization: &models.SimpleOrganization{
			ID:      &orgId,
			Name:    &report.Organization.Name,
			Deleted: &report.Organization.Deleted,
		},
	}
}

func BonusReportPageToRest(list entities.Page[entities.BonusReport]) *models.ReportPage {
	items := []*models.Report{}
	for _, v := range list.Items {
		i := BonusReportToRest(v)
		items = append(items, &i)
	}

	return &models.ReportPage{
		Page:       &list.Page,
		PageSize:   &list.PageSize,
		TotalPages: &list.TotalPages,
		TotalItems: &list.TotalItems,
		Items:      items,
	}
}

func BonusReportFromDB(report dbmodels.BonusReport) entities.BonusReport {
	return entities.BonusReport{
		ID:            report.ID,
		Amount:        report.Amount,
		Date:          report.Date,
		UserID:        report.UserID,
		OperationType: OperationTypeFromDB(report.OperationType),
		Organization: entities.SimleOrganization{
			ID:      report.Organization.ID,
			Name:    report.Organization.Name,
			Deleted: report.Organization.Deleted,
		},
	}
}

func BonusReportsFromDB(report []dbmodels.BonusReport) []entities.BonusReport {
	items := []entities.BonusReport{}
	for _, v := range report {
		items = append(items, BonusReportFromDB(v))
	}
	return items
}
