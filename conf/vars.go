package conf

import (
	"sync"
)

type CmdConfig struct {
	IpAddress   string //rpc服务监听的ip端口
	Version     string //rpc版本
	ServerID    int32  //rpc的svrid
	ServerName  string //rpc服务名称
	EtcdAddress string
	Network     string //协议类型
	Level       int32  //gamesvr使用的场次信息（仅供gamesvr使用，其他服可无视）
	Type        string //标识此game的类型 “M”标识比赛用game “H”标识大厅用game
	AcPort      int    //ac服务对外放开的端口
	PLogConf    string //plog的ini配置文件路径
	Closed      bool   //
	IpCode      string //
	Cfg         string //
	ENV         int32  //环境：1、线上环境；2、开发环境；3、本地环境
}

type CustomCfg struct {
	Custom map[string]string `json:"custom"`
}

type DBConf struct {
	DBCount    int32    `json:"DBCount"`
	DSN        []string `json:"DSN"`
	TableCount int32    `json:"TableCount"`
	Charset    string   `json:"Charset"`
	Database   string   `json:"Database"`
}

type RedisConf struct {
	Host           string `json:"Host"`           //IP地址:端口号
	Pass           string `json:"Pass"`           //密码
	MaxIdle        int32  `json:"MaxIdle"`        //最大空闲连接数
	MaxActive      int32  `json:"MaxActive"`      //最大连接数(包含了空闲的)
	IdleTimeout    int32  `json:"IdleTimeout"`    //池中的连接空闲多久之后超时（单位：秒）
	ConnectTimeout int32  `json:"ConnectTimeout"` //连接超时时间（单位:毫秒）
	ReadTimeout    int32  `json:"ReadTimeout"`    //读超时时间（单位:毫秒）
	WriteTimeout   int32  `json:"WriteTimeout"`   //写超时时间（单位:毫秒）
	Database       int32  `json:"Database"`       //数据库编号
}

var (
	CmdConf       *CmdConfig
	mysqlConf     = make(map[string]*DBConf)    //mysql 库表对应配置
	redisPoolConf = make(map[string]*RedisConf) //redis全局连接池配置
	otherConf     = make(map[string]string)     //redis全局连接池配置
	etdConfLock   sync.RWMutex
	etcdKey       = make(map[string]string)

	etcdWatch *EtcdWatcher
	//etcdConf      * serverplugin.EtcdV3RegisterPlugin
)
