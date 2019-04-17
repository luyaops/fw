package register

import (
	"luyaops/fw/common/etcd"
	"luyaops/fw/common/log"
)

func Register(store *etcd.Store, server, apiJson string) {
	if err := store.Put(server, apiJson); err != nil {
		log.Errorf("register %v server error:%v", server, err)
	} else {
		log.Infof("register %v server successfully", server)
	}
}
