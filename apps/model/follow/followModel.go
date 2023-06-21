package follow

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		RowBuilder() squirrel.SelectBuilder
		FindOneByUidAndDid(ctx context.Context, rowBuilder squirrel.SelectBuilder, uid int64, did int64) ([]*Follow, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn, c cache.CacheConf) FollowModel {
	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn, c),
	}
}

func (m *defaultFollowModel) FindOneByUidAndDid(ctx context.Context, rowBuilder squirrel.SelectBuilder, uid int64, did int64) ([]*Follow, error) {

	query, values, err := rowBuilder.Where(" uid = ? And did = ? limit 1 ", uid, did).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Follow
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(followRows).From(m.table)
}
