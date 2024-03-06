package data

import (
	"context"
	"fmt"
	"idgen/internal/conf"
	"idgen/internal/pkg/etcdutil"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	etcdClientv3 "go.etcd.io/etcd/client/v3"
)

const snowflakeNodeSequenceKey = "id-gen/snowflake-node-sequence"

const snowflakeNodeKeyPrefix = "id-gen/snowflake-node"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEtcdClient, NewSnowflakeNode, NewSequenceRepo, NewSnowflakeRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	etcdCli       *etcdClientv3.Client
	snowflakeNode *snowflake.Node
}

// NewData .
func NewData(etcdCli *etcdClientv3.Client, snowflakeNode *snowflake.Node, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		etcdCli.Close()
	}
	return &Data{
		etcdCli:       etcdCli,
		snowflakeNode: snowflakeNode,
	}, cleanup, nil
}

func NewEtcdClient(conf *conf.Data, logger log.Logger) *etcdClientv3.Client {
	log := log.NewHelper(log.With(logger, "module", "idgen-service/data/etcd"))

	client, err := etcdClientv3.New(etcdClientv3.Config{Endpoints: conf.Etcd.Endpoints, DialTimeout: conf.Etcd.Timeout.AsDuration()})
	if err != nil {
		log.Fatalf("failed opening connection to etcd: %v", err)
	}
	return client
}

func newSnowflakeNodeID(ctx context.Context, etcdCli *etcdClientv3.Client, log *log.Helper) (int64, error) {
	var retry int8
	var snowflakeNodeID int64
	var nodeMax int64 = -1 ^ (-1 << snowflake.NodeBits)

	for {
		lastValue, err := etcdutil.AddInt64(ctx, etcdCli, snowflakeNodeSequenceKey, 1)
		if err != nil {
			if retry < 2 {
				retry++
				continue
			} else {
				log.Fatalf("failed to etcdutil.AddUint64: %v", err)
			}
		}

		snowflakeNodeID = int64(lastValue) % (nodeMax + 1)

		grant, err := etcdCli.Grant(ctx, 10)
		if err != nil {
			if retry < 2 {
				retry++
				continue
			} else {
				log.Fatalf("failed to etcdCli.Grant: %v", err)
			}
		}
		lease := grant.ID

		key := fmt.Sprintf("%s/%d", snowflakeNodeKeyPrefix, snowflakeNodeID)
		cmp := etcdClientv3.Compare(etcdClientv3.LeaseValue(key), "=", etcdClientv3.NoLease)
		put := etcdClientv3.OpPut(key, "occupied", etcdClientv3.WithLease(lease))

		txnResp, txnErr := etcdCli.Txn(ctx).If(cmp).Then(put).Commit()
		if txnErr != nil {
			if retry < 2 {
				retry++
				continue
			} else {
				log.Fatalf("failed to etcdCli.Txn: %v", err)
			}
		}

		if txnResp.Succeeded {
			ch, _ := etcdCli.KeepAlive(context.Background(), lease)
			go func() {
				for {
					if _, ok := <-ch; !ok {
						log.Fatalf("failed to etcdCli.KeepAlive: %v", err)
					}
				}
			}()

			break
		}
	}

	return snowflakeNodeID, nil
}

func NewSnowflakeNode(etcdCli *etcdClientv3.Client, logger log.Logger) *snowflake.Node {
	log := log.NewHelper(log.With(logger, "module", "idgen-service/data/snowflake"))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	snowflakeNodeID, _ := newSnowflakeNodeID(ctx, etcdCli, log)

	snowflake.Epoch = 1704067200000
	node, err := snowflake.NewNode(snowflakeNodeID)
	if err != nil {
		log.Fatalf("failed create snowflakeNode: %v", err)
	}

	return node
}
