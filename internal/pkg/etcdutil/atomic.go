package etcdutil

import (
	"context"
	"strconv"

	etcdClientv3 "go.etcd.io/etcd/client/v3"
)

func AddInt64(ctx context.Context, etcdCli *etcdClientv3.Client, key string, delta int64) (int64, error) {
	var new int64

	for {
		getResp, getErr := etcdCli.Get(ctx, key)
		if getErr != nil {
			return new, getErr
		}
		if len(getResp.Kvs) == 0 {
			return new, nil
		}
		oldValue := string(getResp.Kvs[0].Value)
		oldValueUint, strconvErr := strconv.ParseInt(oldValue, 10, 64)
		if strconvErr != nil {
			return new, strconvErr
		}
		new = oldValueUint + delta
		newValue := strconv.FormatInt(new, 10)

		swapped, _ := CompareAndSwap(ctx, etcdCli, key, oldValue, newValue)
		if swapped {
			break
		}
	}

	return new, nil
}

func CompareAndSwap(ctx context.Context, etcdCli *etcdClientv3.Client, key, old, new string) (bool, error) {
	cmp := etcdClientv3.Compare(etcdClientv3.Value(key), "=", old)
	put := etcdClientv3.OpPut(key, new)

	txnResp, txnErr := etcdCli.Txn(ctx).If(cmp).Then(put).Commit()
	if txnErr != nil {
		return txnResp.Succeeded, txnErr
	}

	return txnResp.Succeeded, nil
}
