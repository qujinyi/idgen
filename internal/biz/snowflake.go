package biz

import (
	"context"

	v1 "idgen/api/idgen/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrContentMissing is content missing.
	ErrContentMissing = errors.NotFound(v1.ErrorReason_CONTENT_MISSING.String(), "content missing")
)

// Snowflake is a Snowflake model.
type Snowflake struct {
	ID int64
}

// SnowflakeRepo is a Snowflake repo.
type SnowflakeRepo interface {
	Generate(context.Context) (*Snowflake, error)
}

// SnowflakeUsecase is a Snowflake usecase.
type SnowflakeUsecase struct {
	repo SnowflakeRepo
	log  *log.Helper
}

// NewSnowflakeUsecase new a Snowflake usecase.
func NewSnowflakeUsecase(repo SnowflakeRepo, logger log.Logger) *SnowflakeUsecase {
	return &SnowflakeUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Generate returns the generate Snowflake.
func (uc *SnowflakeUsecase) Generate(ctx context.Context) (*Snowflake, error) {
	uc.log.WithContext(ctx).Info("Snowflake.Generate")
	return uc.repo.Generate(ctx)
}
