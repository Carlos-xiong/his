package mdmbrandservicelogic

import (
	"context"

	"his/mdm/rpc/internal/svc"
	"his/mdm/rpc/mdm"

	"github.com/zeromicro/go-zero/core/logx"
)

type MdmBrandDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMdmBrandDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MdmBrandDetailsLogic {
	return &MdmBrandDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌信息
func (l *MdmBrandDetailsLogic) MdmBrandDetails(in *mdm.MdmBrandInfoRequest) (*mdm.MdmBrandInfo, error) {
	// todo: add your logic here and delete this line
	brand, err := l.svcCtx.MdmBrandModel.Get(context.Background(), int(in.BrandId))
	if err != nil {
		return nil, err
	}

	return &mdm.MdmBrandInfo{BrandId: int64(brand.ID), BrandName: brand.BrandName}, nil
}
