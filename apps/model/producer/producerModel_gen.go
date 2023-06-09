// Code generated by goctl. DO NOT EDIT.

package producer

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	producerFieldNames          = builder.RawFieldNames(&Producer{})
	producerRows                = strings.Join(producerFieldNames, ",")
	producerRowsExpectAutoSet   = strings.Join(stringx.Remove(producerFieldNames, "`id`"), ",")
	producerRowsWithPlaceHolder = strings.Join(stringx.Remove(producerFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCpAudioProducerIdPrefix = "cache:cpAudio:producer:id:"
)

type (
	producerModel interface {
		Insert(ctx context.Context, data *Producer) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Producer, error)
		Update(ctx context.Context, data *Producer) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProducerModel struct {
		sqlc.CachedConn
		table string
	}

	Producer struct {
		Id        int64  `db:"id"`
		Name      string `db:"name"`
		Uid       int64  `db:"uid"`
		Status    int64  `db:"status"` // 0不展示 1展示
		Mark      string `db:"mark"`
		CreatedAt int64  `db:"created_at"`
	}
)

func newProducerModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProducerModel {
	return &defaultProducerModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`producer`",
	}
}

func (m *defaultProducerModel) Delete(ctx context.Context, id int64) error {
	cpAudioProducerIdKey := fmt.Sprintf("%s%v", cacheCpAudioProducerIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, cpAudioProducerIdKey)
	return err
}

func (m *defaultProducerModel) FindOne(ctx context.Context, id int64) (*Producer, error) {
	cpAudioProducerIdKey := fmt.Sprintf("%s%v", cacheCpAudioProducerIdPrefix, id)
	var resp Producer
	err := m.QueryRowCtx(ctx, &resp, cpAudioProducerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", producerRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProducerModel) Insert(ctx context.Context, data *Producer) (sql.Result, error) {
	cpAudioProducerIdKey := fmt.Sprintf("%s%v", cacheCpAudioProducerIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, producerRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Uid, data.Status, data.Mark,data.CreatedAt)
	}, cpAudioProducerIdKey)
	return ret, err
}

func (m *defaultProducerModel) Update(ctx context.Context, data *Producer) error {
	cpAudioProducerIdKey := fmt.Sprintf("%s%v", cacheCpAudioProducerIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, producerRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Uid, data.Status, data.Mark, data.Id)
	}, cpAudioProducerIdKey)
	return err
}

func (m *defaultProducerModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCpAudioProducerIdPrefix, primary)
}

func (m *defaultProducerModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", producerRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProducerModel) tableName() string {
	return m.table
}
