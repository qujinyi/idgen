package data

import (
	"context"
	"time"

	"idgen/internal/biz"
	"idgen/internal/pkg/etcdutil"

	"github.com/go-kratos/kratos/v2/log"
)

const sequenceKey = "id-gen/sequence"

type sequenceRepo struct {
	data        *Data
	log         *log.Helper
	sequenceKey string
}

// NewSequenceRepo.
func NewSequenceRepo(data *Data, logger log.Logger) biz.SequenceRepo {
	return &sequenceRepo{
		data:        data,
		log:         log.NewHelper(logger),
		sequenceKey: sequenceKey,
	}
}

func (r *sequenceRepo) Next(ctx context.Context) (*biz.Sequence, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	nextValue, err := etcdutil.AddInt64(ctx, r.data.etcdCli, r.sequenceKey, 1)

	return &biz.Sequence{
		ID: nextValue,
	}, err
}
