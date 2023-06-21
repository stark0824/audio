package svc

import (
	"audio/apps/model/drama"
	"audio/apps/model/follow"
	"audio/apps/model/producer"
	"audio/apps/model/recommend"
	"audio/apps/rpc/drama/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	SqlCache       cache.CacheConf
	DramaModel     drama.DramaModel
	ProducerModel  producer.ProducerModel
	RecommendModel recommend.RecPositionModel
	FollowModel    follow.FollowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		DramaModel:     drama.NewDramaModel(sqlx.NewMysql(string(c.DB.DataSource)), c.SqlCache),
		ProducerModel:  producer.NewProducerModel(sqlx.NewMysql(string(c.DB.DataSource)), c.SqlCache),
		RecommendModel: recommend.NewRecPositionModel(sqlx.NewMysql(string(c.DB.DataSource)), c.SqlCache),
		FollowModel:    follow.NewFollowModel(sqlx.NewMysql(string(c.DB.DataSource)), c.SqlCache),
	}
}
