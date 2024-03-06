package data

import (
	"context"

	"idgen/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type snowflakeRepo struct {
	data *Data
	log  *log.Helper
}

// NewSnowflakeRepo.
func NewSnowflakeRepo(data *Data, logger log.Logger) biz.SnowflakeRepo {
	return &snowflakeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *snowflakeRepo) Generate(ctx context.Context) (*biz.Snowflake, error) {
	return &biz.Snowflake{
		ID: r.data.snowflakeNode.Generate().Int64(),
	}, nil
}
