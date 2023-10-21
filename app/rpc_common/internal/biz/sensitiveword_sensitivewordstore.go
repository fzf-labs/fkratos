package biz

import (
	"context"
	"fkratos/app/rpc_common/internal/data/gorm/fkratos_common_model"
	"fkratos/internal/errorx"

	pb "fkratos/api/rpc_common/v1"

	"github.com/fzf-labs/fpkg/util/jsonutil"
)

// SensitiveWordStore 敏感词-保存
func (s *SensitiveWordUseCase) SensitiveWordStore(ctx context.Context, req *pb.SensitiveWordStoreReq) (*pb.SensitiveWordStoreReply, error) {
	resp := &pb.SensitiveWordStoreReply{}
	labs, err := jsonutil.Marshal(req.GetLabs())
	if err != nil {
		return nil, errorx.DataFormattingError.WithCause(err).WithMetadata(errorx.SetErrMetadata(err))
	}
	_, err = s.sensitiveWordRepo.SensitiveWordStore(ctx, &fkratos_common_model.SensitiveWord{
		ID:   req.GetId(),
		Word: req.GetWord(),
		Labs: labs,
		Desc: req.GetDesc(),
	})
	if err != nil {
		return nil, err
	}
	err = s.sensitiveWordRepo.SensitiveWordsCacheDel(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
