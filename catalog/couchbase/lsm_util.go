package couchbase

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type cacheManagerState struct {
	serverUuid         string
	isRunning          bool
	isWaitingForNotify bool
}

var cacheMgrState cacheManagerState

func (ie IndexError) Error() string {
	return fmt.Sprintf("Code: %v, Msg: %v", ie.Code, ie.Msg)
}

var client = NewRestClient("http://localhost:8094")

func (li *lsmIndex) DropLsmIndex() query.Error {

	clog.To(catalog.CHANNEL, "Requesting Drop Index %s", li)

	//FIXME Not using ServerUUID. This needs to be used to update the cache manager state
	serverUuid, err := client.Drop(li.Id())

	if err != nil {
		clog.To(catalog.CHANNEL, "Error for Drop Index %v", err)
		return query.NewError(nil, err.Error())
	}

	//Update the Cache Uuid only if its in waiting state. If notification processing
	//is already happening, it will get updated from the server.
	if cacheMgrState.isWaitingForNotify {
		clog.To(catalog.CHANNEL, "Drop Updated Cache Uuid to %s. Was %s", serverUuid, cacheMgrState.serverUuid)
		cacheMgrState.serverUuid = serverUuid
	}

	clog.To(catalog.CHANNEL, "Drop Index Returned")

	return nil
}

func loadLsmIndexesForBucket(b *bucket) ([]*catalog.Index, error) {

	clog.To(catalog.CHANNEL, "Requesting List Index")

	//This function gets called when bucket gets used for first time. So even though
	//client is in sync with the server, we need indexes for this bucket.
	//Sending ServerUuid=0 makes the server send all the data again.
	cacheServerUuid := 0
	serverUuid, indexInfos, err := client.List(string(cacheServerUuid))

	if err != nil {
		clog.To(catalog.CHANNEL, "Error Response for List Index %v", err)
		return nil, err
	}

	clog.To(catalog.CHANNEL, "List Index Response %v", indexInfos)

	indexes := make([]*catalog.Index, 0, len(indexInfos))
	for _, i := range indexInfos {

		//Ignore if index is not for this bucket. Server sends data for all buckets.
		clog.To(catalog.CHANNEL, "Processing Index %v", i)
		if b.Name() == i.Bucket {

			var index catalog.Index

			if i.IsPrimary {
				index = &primaryLsmIndex{
					lsmIndex{
						name:   i.Name,
						uuid:   i.Uuid,
						bucket: b,
						using:  catalog.LSM,
					},
				}
				indexes = append(indexes, &index)
			} else {

				on := make(catalog.IndexKey, len(i.OnExprList))

				for pos, str := range i.OnExprList {
					expr, err := ast.UnmarshalExpression([]byte(str))
					if err != nil {
						return nil, errors.New("Cannot unmarshal expression for index " + i.Name)
					}
					on[pos] = expr
				}

				index = &lsmIndex{
					name:   i.Name,
					uuid:   i.Uuid,
					bucket: b,
					using:  catalog.LSM,
					on:     on,
				}
				indexes = append(indexes, &index)
			}
			clog.To(catalog.CHANNEL, "Processed Index %v", index)
		}
	}

	if cacheMgrState.isRunning != true {
		cacheMgrState.isRunning = true
		cacheMgrState.serverUuid = serverUuid
		go manageLSMIndexCache(b.pool)
	}

	return indexes, nil

}

func loadLsmIndexesForPool(p *pool) (string, error) {

	clog.To(catalog.CHANNEL, "Updating Pool with Indexes Information")

	serverUuid, indexInfos, err := client.List("")

	if err != nil {
		clog.To(catalog.CHANNEL, "Error Response for List Index %v", err)
		return "", err
	}

	clog.To(catalog.CHANNEL, "List Index Response %v", indexInfos)

	//Clear all the previous LSM indexes

	for _, bc := range p.bucketCache {
		bi := bc.(*bucket)
		for i, info := range bi.indexes {
			if info.Type() == catalog.LSM {
				delete(bi.indexes, i)
			}
		}
	}

	var b *bucket

	for _, i := range indexInfos {

		//Ignore if index is not for this bucket. Server sends data for all buckets.
		clog.To(catalog.CHANNEL, "Processing Index %v", i)
		//FIXME check if this type assertion is idiomatic or need a separate method in bucket
		if b = p.bucketCache[i.Bucket].(*bucket); b != nil {

			clog.To(catalog.CHANNEL, "Bucket Info %v", b)
			var index catalog.Index

			if i.IsPrimary {
				index = &primaryLsmIndex{
					lsmIndex{
						name: i.Name,
						uuid: i.Uuid,
						//FIXME add a private method to pool to return this
						bucket: b,
						using:  catalog.LSM,
					},
				}
				b.indexes[i.Name] = index
			} else {

				on := make(catalog.IndexKey, len(i.OnExprList))

				for pos, str := range i.OnExprList {
					expr, err := ast.UnmarshalExpression([]byte(str))
					if err != nil {
						return "", errors.New("Cannot unmarshal expression for index " + i.Name)
					}
					on[pos] = expr
				}

				index = &lsmIndex{
					name: i.Name,
					uuid: i.Uuid,
					//FIXME add a private method to pool to return this
					bucket: b,
					using:  catalog.LSM,
					on:     on,
				}
				b.indexes[i.Name] = index
			}
			clog.To(catalog.CHANNEL, "Processed Index %v", index)
		}
	}

	return serverUuid, nil

}

func manageLSMIndexCache(pool *pool) {

	defer func() {
		cacheMgrState.isRunning = false
	}()

	for {

		clog.To(catalog.CHANNEL, "Index Cache Manager Started for Pool %v", pool.bucketCache)
		cacheMgrState.isWaitingForNotify = true
		status, serverUuid, err := client.Notify(cacheMgrState.serverUuid)

		if err == nil {
			if status == INVALID_CACHE {
				clog.To(catalog.CHANNEL, "Index Cache Manager Received Notification : Invalid Cache. ServerUuid %s CacheServerUuid %s ", serverUuid, cacheMgrState.serverUuid)
				if cacheMgrState.serverUuid == serverUuid {
					clog.To(catalog.CHANNEL, "Index Cache Manager ServerUuid matches Cache Uuid. Must be our DDL Request. Ignoring.")
					continue
				} else {
					cacheMgrState.isWaitingForNotify = false
					serverUuid, err = loadLsmIndexesForPool(pool)
					if err == nil {
						cacheMgrState.serverUuid = serverUuid
						clog.To(catalog.CHANNEL, "Index Cache Manager ServerUUID updated to %s", cacheMgrState.serverUuid)
					} else {
						clog.To(catalog.CHANNEL, "Error while updating Index Cache for ServerUUID %s %v", cacheMgrState.serverUuid, err)
					}
					cacheMgrState.isWaitingForNotify = true
				}
			} else {
				clog.To(catalog.CHANNEL, "Cache Manager Received Notification %v. Ignoring.", status)
			}
		} else {
			clog.To(catalog.CHANNEL, "Server error waiting for notification %v. Retry in 10 secs", err)
			//FIXME make this more understandable
			time.Sleep(10e9)
		}

		clog.To(catalog.CHANNEL, "Index Cache Manager Exited for Pool %v", pool.bucketCache)
	}

}

func newPrimaryIndex(bkt *bucket) (*primaryLsmIndex, error) {

	index := IndexInfo{
		Name:      PRIMARY_INDEX,
		Using:     catalog.LSM,
		Bucket:    bkt.Name(),
		IsPrimary: true,
	}

	clog.To(catalog.CHANNEL, "Create Index Request %v", index)

	//FIXME Not using ServerUUID. This needs to be used to update the cache manager state
	serverUuid, indexinfo, err := client.Create(index)

	if err != nil {
		return nil, err
	}

	//Update the Cache Uuid only if its in waiting state. If notification processing
	//is already happening, it will get updated from the server
	if cacheMgrState.isWaitingForNotify {
		clog.To(catalog.CHANNEL, "Create Updated Cache Uuid to %s. Was %s", serverUuid, cacheMgrState.serverUuid)
		cacheMgrState.serverUuid = serverUuid
	}

	clog.To(catalog.CHANNEL, "Create Index Response %v", indexinfo)

	inst := primaryLsmIndex{
		lsmIndex{
			name:   PRIMARY_INDEX,
			uuid:   indexinfo.Uuid,
			using:  catalog.LSM,
			bucket: bkt,
		},
	}

	return &inst, nil

}

func newLsmIndex(name string, on catalog.IndexKey, bkt *bucket) (*lsmIndex, error) {

	var exprList []string

	exprList = make([]string, len(on))

	for i, e := range on {
		m, _ := json.Marshal(e)
		exprList[i] = string(m)
	}

	clog.To(catalog.CHANNEL, "Marshalled Index Expr %s", exprList)

	index := IndexInfo{
		Name:       name,
		Using:      catalog.LSM,
		OnExprList: exprList,
		Bucket:     bkt.Name(),
		IsPrimary:  false,
	}

	clog.To(catalog.CHANNEL, "Create Index Request %v", index)

	//FIXME Not using ServerUUID. This needs to be used to update the cache manager state
	serverUuid, indexinfo, err := client.Create(index)

	if err != nil {
		return nil, err
	}

	//Update the Cache Uuid only if its in waiting state. If notification processing
	//is already happening, it will get updated from the server
	if cacheMgrState.isWaitingForNotify {
		clog.To(catalog.CHANNEL, "Create Updated Cache Uuid to %s. Was %s", serverUuid, cacheMgrState.serverUuid)
		cacheMgrState.serverUuid = serverUuid
	}

	clog.To(catalog.CHANNEL, "Create Index Response %v", indexinfo)

	inst := lsmIndex{
		name:   name,
		uuid:   indexinfo.Uuid,
		using:  catalog.LSM,
		on:     on,
		bucket: bkt,
	}

	return &inst, nil
}
