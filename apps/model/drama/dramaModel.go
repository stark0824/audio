package drama

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DramaModel = (*customDramaModel)(nil)

type (
	// DramaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDramaModel.
	DramaModel interface {
		dramaModel
		RowBuilder() squirrel.SelectBuilder
		FindListByIds(ctx context.Context, rowBuilder squirrel.SelectBuilder, IndexIds string) ([]*Drama, error)
	}

	customDramaModel struct {
		*defaultDramaModel
	}
)

// NewDramaModel returns a model for the database table.
func NewDramaModel(conn sqlx.SqlConn, c cache.CacheConf) DramaModel {
	return &customDramaModel{
		defaultDramaModel: newDramaModel(conn, c),
	}
}

func (m *defaultDramaModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(dramaRows).From(m.table)
}

func (m *defaultDramaModel) FindPageAdminListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Drama, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Drama
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultDramaModel) FindListByIds(ctx context.Context, rowBuilder squirrel.SelectBuilder, IndexIds string) ([]*Drama, error) {

	rowBuilder = rowBuilder.Where(" id in  ( " + IndexIds + " ) ")
	query, values, err := rowBuilder.Where("status = 1").OrderBy("id ASC").ToSql()
	if err != nil {
		logx.Info(err)
		return nil, err
	}

	var resp []*Drama
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
