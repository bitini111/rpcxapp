package rpcxapp

import (
	"context"
	"github.com/bitini111/rpcxapp/conf"
	serverplugin "github.com/bitini111/rpcxapp/plugin/etcdv3/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
)

func Run(ctl interface{}, shutdown func(s *server.Server)) error {
	srv := server.NewServer()
	r := serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: conf.CmdConf.Network + "@" + conf.CmdConf.IpAddress, //服务监听的ip端口
		EtcdServers:    conf.GetEtcdAddr(),                                  //zookeeper地址
		BasePath:       conf.GetBasePath(),                                  //zk的目录
		Metrics:        metrics.NewRegistry(),
		Version:        conf.CmdConf.Version,  //rpc的版本
		ServerID:       conf.CmdConf.ServerID, //rpc的svrid
		//UpdateInterval: time.Minute,
	}
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
	ss.RegisterOnShutdown(shutdown)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ss.Shutdown(ctx)
	return
}