package config

import (
	"flag"
	"github.com/luyaops/fw/common/log"
	"os"
	"strings"
)

var (
	RpcBind     string
	GatewayBind string
	dbURL       string
	EtcdAddr    string
	LogLevel    string
	Suffix      string
	isHelp      bool
	Endpoints   []string
)

func init() {
	flag.StringVar(&RpcBind, "rbind", "0.0.0.0:50051", "Bind RPC Server Address")
	flag.StringVar(&GatewayBind, "gbind", "0.0.0.0:8080", "Bind API Gateway Address")
	flag.StringVar(&EtcdAddr, "etcdAddr", "localhost:2379", "Multiple Etcd Address: 127.0.0.1:2379,192.168.0.10:2379")
	flag.StringVar(&dbURL, "dbURL", "root:123456@tcp(127.0.0.1:3306)/github.com/luyaops", "DB URL")
	flag.StringVar(&LogLevel, "logLevel", "debug", "Log Level: debug info warn error fatal")
	flag.BoolVar(&isHelp, "help", false, "Print this help")
	flag.Parse()

	// 转换EtcdAddr为[]Endpoints
	Endpoints = strings.Split(EtcdAddr, ",")

	log.SetLoggerLevel(LogLevel)
	if isHelp {
		showHelp()
	}
}

func showHelp() {
	flag.PrintDefaults()
	os.Exit(1)
}
