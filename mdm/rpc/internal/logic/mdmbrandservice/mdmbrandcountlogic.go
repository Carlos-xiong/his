package mdmbrandservicelogic

import (
	"context"

	"his/mdm/rpc/internal/svc"
	"his/mdm/rpc/mdm"

	"github.com/zeromicro/go-zero/core/logx"
)

type MdmBrandCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMdmBrandCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MdmBrandCountLogic {
	return &MdmBrandCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌数量
func (l *MdmBrandCountLogic) MdmBrandCount(in *mdm.MdmBrandListRequest) (*mdm.TotalResponse, error) {
	// todo: add your logic here and delete this line

	return &mdm.TotalResponse{}, nil
}
