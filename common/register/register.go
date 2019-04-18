package register

import (
	"github.com/luyaops/fw/common/constants"
	"github.com/luyaops/fw/common/etcd"
	"github.com/luyaops/fw/common/log"
)

func Register(store *etcd.Store, server, apiJson string) {
	if err := store.Put(constants.RegistryPrefix+server, apiJson); err != nil {
		log.Fatalf("Register %v server error:%v", server, err)
	} else {
		log.Infof("Register %v server successfully", server)
	}
}
