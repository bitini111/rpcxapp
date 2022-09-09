package rpcxapp

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

import (
	"context"
	"github.com/bitini111/rpcxapp/conf"
	serverplugin "github.com/bitini111/rpcxapp/plugin/etcdv3/server"
	"github.com/smallnest/rpcx/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctl interface{}, shutdown func(s *server.Server)) error {
	srv := server.NewServer()
	serviceAddress, etcdAddress, basePath, version, serid := conf.CmdConf.Network+"@"+conf.CmdConf.IpAddress, conf.GetEtcdAddr(), conf.GetBasePath(), conf.CmdConf.Version, conf.CmdConf.ServerID
	r := serverplugin.NewEtcdV3Plugin(serviceAddress, etcdAddress, basePath, version, serid)
	err := r.Start()
	if err != nil {
		srv.Close()
		os.Exit(1)
		return err
	}

	//srv.RegisterOnShutdown(shutdown)

	go WaitTerminationSignal(srv, shutdown)

	srv.RegisterName(conf.CmdConf.ServerName, ctl, conf.CmdConf.ServerName) //服务名，以及服务的接收方法
	er := srv.Serve(conf.CmdConf.Network, conf.CmdConf.IpAddress)
	if er != nil {
		log.Fatalln(err)
	}

	return nil
}

func WaitTerminationSignal(ss *server.Server, shutdown func(s *server.Server)) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGUSR1, syscall.SIGTERM)
	defer func() {
		signal.Stop(ch)
		close(ch)
	}()
	<-ch
	conf.CloseEtcdWatch()
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ss.Shutdown(ctx)
	return
}