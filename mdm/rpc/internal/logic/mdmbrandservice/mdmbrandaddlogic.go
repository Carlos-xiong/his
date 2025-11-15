package mdmbrandservicelogic

import (
	"context"

	"his/mdm/rpc/internal/svc"
	"his/mdm/rpc/mdm"

	"github.com/zeromicro/go-zero/core/logx"
)

type MdmBrandAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMdmBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MdmBrandAddLogic {
	return &MdmBrandAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌增加
func (l *MdmBrandAddLogic) MdmBrandAdd(in *mdm.MdmBrandInfo) (*mdm.EmptyResponse, error) {
	// todo: add your logic here and delete this line

	return &mdm.EmptyResponse{}, nil
}
