package main

import (
	"github.com/prometheus/common/log"
	"github.com/boxproject/lib-bitcore/serpcclient"
)

func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &serpcclient.ConnConfig{
		Host:         "YourRPCHost:YourRPCPort",
		User:         "YourRPCUserName",
		Pass:         "YourRPCPassword",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := serpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()
	rt,err := client.RescanBlockChain()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(rt)



}
