package reportusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

type reportStorage interface {
	Create(ctx context.Context, data *entities.Report) error
	GetByID(ctx context.Context, id int) (*entities.Report, error)
	UpdateByID(ctx context.Context, id int, data *entities.Report) error
	ListReport(ctx context.Context, paging *common.Paging, filter *reportiomodel.Filter) ([]*entities.Report, error)
}

type reportUseCase struct {
	reportSto    reportStorage
	accountCache accountCache
}

type accountCache interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

func NewReportUseCase(reportSto reportStorage, accountCache accountCache) *reportUseCase {
	return &reportUseCase{reportSto: reportSto, accountCache: accountCache}
}
