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

func Run(cfg *conf.AppConfig, ctl interface{}, shutdown func(s *server.Server)) error {
	srv := server.NewServer()
	r := serverplugin.NewEtcdV3Plugin(cfg.Network+"@"+cfg.IP, cfg.EtcdAddress, cfg.RpcPath, cfg.Version, cfg.ServerID)
	err := r.Start()
	if err != nil {
		srv.Close()
		os.Exit(1)
		return err
	}
	srv.Plugins.Add(r)

	go WaitTerminationSignal(srv, shutdown)

	srv.RegisterName(cfg.ServerName, ctl, cfg.ServerName) //服务名，以及服务的接收方法
	er := srv.Serve(cfg.Network, cfg.IP)
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
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ss.Shutdown(ctx)
	return
}
