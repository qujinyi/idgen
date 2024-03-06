package biz

import (
	"context"

	v1 "idgen/api/idgen/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrNotFound is not found.
	ErrNotFound = errors.NotFound(v1.ErrorReason_NOT_FOUND.String(), "not found")
)

// Sequence is a Sequence model.
type Sequence struct {
	ID int64
}

// SequenceRepo is a Sequence repo.
type SequenceRepo interface {
	Next(context.Context) (*Sequence, error)
}

// SequenceUsecase is a Sequence usecase.
type SequenceUsecase struct {
	repo SequenceRepo
	log  *log.Helper
}

// NewSequenceUsecase new a Sequence usecase.
func NewSequenceUsecase(repo SequenceRepo, logger log.Logger) *SequenceUsecase {
	return &SequenceUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Next returns the next Sequence.
func (uc *SequenceUsecase) Next(ctx context.Context) (*Sequence, error) {
	uc.log.WithContext(ctx).Info("Sequence.Next")
	return uc.repo.Next(ctx)
}
