package cmd

import (
	"fmt"

	"github.com/bitini111/rpcxapp/comm"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

//var app = kingpin.New("gokins", "A golang workflow application.")

//优先级 命令行 -->配置文件
func Run() {
	err := initConf()
	if err != nil {
		fmt.Println(err.Error())
		regs()
	}

	//kingpin.MustParse(app.Parse(os.Args))
	//app.Parse(os.Args[1:])
	//kingpin.MustParse(app.Parse(os.Args[1:]))
}

func initConf() error {
	v := viper.New()
	v.SetConfigFile("./app.yml")
	err := v.ReadInConfig()
	if err != nil {
		//log.Fatalf("ReadInConfig error: %s", err.Error())
		fmt.Printf("ReadInConfig error: %s", err.Error())
		return err
	}
	if err = v.Unmarshal(&comm.Cfg); err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Println(cfg)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&comm.Cfg); err != nil {
			fmt.Println(err)
		}
		//fmt.Println(cfg)
	})
	return nil
}

func regs() {
	kingpin.Flag("host", "rpcx host").Short('h').StringVar(&comm.Cfg.Base.Host)
	kingpin.Flag("net", "rpcx network").Short('n').Default("ws").StringVar(&comm.Cfg.Base.Network)
	kingpin.Flag("sn", "rpcx servername").Short('s').StringVar(&comm.Cfg.Base.ServerName)
	kingpin.Flag("si", "rpcx serverid").Short('i').Int32Var(&comm.Cfg.Base.ServerID)
	kingpin.Flag("path", "rpcx path").Short('p').StringVar(&comm.Cfg.Base.RpcPath)
	kingpin.Flag("version", "rpcx version").Short('v').Default("0.0.0.1").StringVar(&comm.Cfg.Base.Version)
	kingpin.Flag("ectds", "rpcx ectds").Short('e').StringsVar(&comm.Cfg.Base.EtcdAddress)
	kingpin.Parse()
	fmt.Println(comm.Cfg.Base)
}
