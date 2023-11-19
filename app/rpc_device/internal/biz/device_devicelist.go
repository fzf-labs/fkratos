package biz

import (
	"context"
	"fkratos/api/paginator"
	"fkratos/internal/dto"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_device/v1"

	"github.com/fzf-labs/fpkg/orm"
)

// DeviceList 设备表-列表数据查询
func (d *DeviceUseCase) DeviceList(ctx context.Context, req *pb.DeviceListReq) (*pb.DeviceListReply, error) {
	resp := &pb.DeviceListReply{
		Paginator: &paginator.PaginatorReply{},
		List:      make([]*pb.DeviceInfo, 0),
	}
	paginatorReq := &orm.PaginatorReq{}
	err := dto.Copy(paginatorReq, req.GetPaginator())
	if err != nil {
		return nil, err
	}
	result, p, err := d.deviceRepo.FindMultiByPaginator(ctx, paginatorReq)
	if err != nil {
		return nil, errorx.DataSQLErr.WithError(err).Err()
	}
	err = dto.Copy(&resp.List, result)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	err = dto.Copy(&resp.Paginator, p)
	if err != nil {
		return nil, errorx.DataFormattingError.WithError(err).Err()
	}
	return resp, nil
}
