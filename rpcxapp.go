package rpcxapp

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

import (
	"context"
	"fmt"
	"github.com/bitini111/rpcx/plugin/etcdv3/serverplugin"
	"github.com/bitini111/rpcx/server"
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppConfig struct {
	Host        string   `json:"host" yaml:"host"`
	Network     string   `json:"network" yaml:"network"`
	ServerID    int32    `json:"serverID" yaml:"serverID"`
	ServerName  string   `json:"serverName" yaml:"serverName"`
	RpcPath     string   `json:"rpcPath" yaml:"rpcPath"`
	Version     string   `json:"version" yaml:"version"`
	EtcdAddress []string `json:"etcdAddress" yaml:"etcdAddress"`
}

func Run(cfg *AppConfig, ctl interface{}, shutdown func(s *server.Server)) error {
	srv := server.NewServer()
	r := serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: cfg.Network + "@" + cfg.Host,
		EtcdServers:    cfg.EtcdAddress,
		BasePath:       cfg.RpcPath,
		Metrics:        metrics.NewRegistry(),
		Version:        cfg.Version,
		ServiceName:    cfg.ServerName,
		ServerID:       int64(cfg.ServerID),
	}
	//r := serverplugin.NewEtcdV3Plugin(cfg.Network+"@"+cfg.Host, cfg.EtcdAddress, cfg.RpcPath, cfg.Version, cfg.ServerID)
	err := r.Start()
	if err != nil {
		srv.Close()
		os.Exit(1)
		return err
	}
	srv.Plugins.Add(r)

	go WaitTerminationSignal(srv, shutdown)
	name := fmt.Sprintf("%s/%s/%d#%s", r.BasePath, r.ServiceName, r.ServerID, r.Version)
	srv.RegisterName(name, ctl, "") //服务名，以及服务的接收方法
	er := srv.Serve(cfg.Network, cfg.Host)
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
