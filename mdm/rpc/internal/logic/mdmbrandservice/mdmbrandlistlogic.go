package mdmbrandservicelogic

import (
	"context"

	"his/mdm/rpc/internal/svc"
	"his/mdm/rpc/mdm"

	"github.com/zeromicro/go-zero/core/logx"
)

type MdmBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMdmBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MdmBrandListLogic {
	return &MdmBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌列表
func (l *MdmBrandListLogic) MdmBrandList(in *mdm.MdmBrandListRequest) (*mdm.MdmBrandListResponse, error) {
	// todo: add your logic here and delete this line

	return &mdm.MdmBrandListResponse{}, nil
}
