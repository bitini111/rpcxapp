package rpcxapp

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"githup.bitini111.rpcxapp/conf"
	serverplugin "githup.bitini111.rpcxapp/plugin/etcdv3/server"
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

	srv.RegisterOnShutdown(shutdown)
	WaitTerminationSignal := func(ss *server.Server) {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		defer func() {
			signal.Stop(ch)
			close(ch)
		}()
		<-ch
		ss.Close()
	}

	WaitTerminationSignal(srv)

	srv.RegisterName(conf.CmdConf.ServerName, ctl, conf.CmdConf.ServerName) //服务名，以及服务的接收方法
	er := srv.Serve(conf.CmdConf.Network, conf.CmdConf.IpAddress)
	if er != nil {
		log.Fatalln(err)
	}

	return nil
}
