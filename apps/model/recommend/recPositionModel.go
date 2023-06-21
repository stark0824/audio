package recommend

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RecPositionModel = (*customRecPositionModel)(nil)
var cacheCpAudioRecPositionAllPrefix = "cache:cpAudio:recPosition:all:"

type (

	// RecPositionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecPositionModel.
	RecPositionModel interface {
		recPositionModel
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RecPosition, error)
		FindAllCache(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RecPosition, error)
	}

	customRecPositionModel struct {
		*defaultRecPositionModel
	}
)

// NewRecPositionModel returns a model for the database table.
func NewRecPositionModel(conn sqlx.SqlConn, c cache.CacheConf) RecPositionModel {
	return &customRecPositionModel{
		defaultRecPositionModel: newRecPositionModel(conn, c),
	}
}

func (m *defaultRecPositionModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(recPositionRows).From(m.table)
}

func (m *defaultRecPositionModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RecPosition, error) {

	rowBuilder = rowBuilder.OrderBy("sort ASC")

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	//todo add cache
	var resp []*RecPosition
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRecPositionModel) FindAllCache(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RecPosition, error) {
	var resp []*RecPosition

	query := fmt.Sprintf("select %s from %s", recPositionRows, m.table)

	err := m.QueryRowCtx(ctx, &resp, "cpAudioRecPositionIdKey:all", func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowsCtx(ctx, v, query)
	})

	if err != nil {
		logx.Errorf("%s", err.Error())
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return resp, nil
}
