package branch

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MdmBranchModel = (*customMdmBranchModel)(nil)

type (
	// MdmBranchModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMdmBranchModel.
	MdmBranchModel interface {
		mdmBranchModel
		withSession(session sqlx.Session) MdmBranchModel
	}

	customMdmBranchModel struct {
		*defaultMdmBranchModel
	}
)

// NewMdmBranchModel returns a model for the database table.
func NewMdmBranchModel(conn sqlx.SqlConn) MdmBranchModel {
	return &customMdmBranchModel{
		defaultMdmBranchModel: newMdmBranchModel(conn),
	}
}

func (m *customMdmBranchModel) withSession(session sqlx.Session) MdmBranchModel {
	return NewMdmBranchModel(sqlx.NewSqlConnFromSession(session))
}
