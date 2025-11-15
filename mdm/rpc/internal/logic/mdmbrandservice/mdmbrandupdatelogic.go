package mdmbrandservicelogic

import (
	"context"

	"his/mdm/rpc/internal/svc"
	"his/mdm/rpc/mdm"

	"github.com/zeromicro/go-zero/core/logx"
)

type MdmBrandUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMdmBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MdmBrandUpdateLogic {
	return &MdmBrandUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌修改
func (l *MdmBrandUpdateLogic) MdmBrandUpdate(in *mdm.MdmBrandInfo) (*mdm.EmptyResponse, error) {
	// todo: add your logic here and delete this line

	return &mdm.EmptyResponse{}, nil
}
