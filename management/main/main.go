package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/jsrccb/nachuan/management"
)

var (
	confFile string
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./management.json", "management.json")
	flag.Parse()
}
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	var (
		err error
	)

	initArgs()
	initEnv()

	if err = management.InitConfig(confFile); err != nil {
		goto ERR
	}

	if err = management.InitMongodbMgr(); err != nil {
		goto ERR
	}

	if err = management.InitUserMgr(); err != nil {
		goto ERR
	}

	if err = management.InitRoleMgr(); err != nil {
		goto ERR
	}

	if err = management.InitClueMgr(); err != nil {
		goto ERR
	}

	if err = management.InitClueTraceMgr(); err != nil {
		goto ERR
	}

	if err = management.InitSaleChannelMgr(); err != nil {
		goto ERR
	}

	if err = management.InitOrderMgr(); err != nil {
		goto ERR
	}

	if err = management.InitCustomMgr(); err != nil {
		goto ERR
	}

	if err = management.InitCollMgr(); err != nil {
		goto ERR
	}

	if err = management.InitCostMgr(); err != nil {
		goto ERR
	}

	if err = management.InitApiserver(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}
	return

ERR:
	fmt.Println(err)

}
