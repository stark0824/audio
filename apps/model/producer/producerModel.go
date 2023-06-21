package producer

import (
	"audio/common/globalkey"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProducerModel = (*customProducerModel)(nil)

type (
	// ProducerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProducerModel.
	ProducerModel interface {
		producerModel
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy int) ([]*Producer, error)
		FindAllByNormal(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Producer, error)
	}

	customProducerModel struct {
		*defaultProducerModel
	}
)

// NewProducerModel returns a model for the database table.
func NewProducerModel(conn sqlx.SqlConn, c cache.CacheConf) ProducerModel {
	return &customProducerModel{
		defaultProducerModel: newProducerModel(conn, c),
	}
}

func (m *defaultProducerModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy int) ([]*Producer, error) {

	if orderBy == 1 {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy("id ASC")
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Producer
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultProducerModel) FindAllByNormal(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Producer, error) {

	rowBuilder = rowBuilder.OrderBy("id DESC")
	query, values, err := rowBuilder.Where("status = ?", globalkey.StateNormal).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Producer
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultProducerModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(producerRows).From(m.table)
}
