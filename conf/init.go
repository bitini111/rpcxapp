package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type listener struct {
}

func (*listener) Set(key []byte, val []byte) {
	//fmt.Println("Set", string(key), string(val))
	upDateConf(key, val)
}
func (*listener) Create(key []byte, val []byte) {
	//fmt.Println("Create", string(key), string(val))
}
func (*listener) Modify(key []byte, val []byte) {
	fmt.Println("Modify", string(key), string(val))
	upDateConf(key, val)
}
func (*listener) Delete(key []byte) {
	//fmt.Println("Delete", string(key))
}

func upDateConf(k []byte, v []byte) error {
	key := fmt.Sprintf("%s", k)
	if vv, ok := etcdKey[key]; ok {
		switch vv {
		case "RedisConf":
			{
				var ob RedisConf
				err := json.Unmarshal(v, &ob)
				if err != nil {
					fmt.Println("redis pool conf json 格式不正确，err:", err)
					return err
				}
				etdConfLock.Lock()
				redisPoolConf[key] = &ob
				etdConfLock.Unlock()
			}
		case "DBConf":
			{
				var ob DBConf
				err := json.Unmarshal(v, &ob)
				if err != nil {
					fmt.Println("db conf json 格式不正确，err:", err)
					return err
				}
				etdConfLock.Lock()
				mysqlConf[key] = &ob
				etdConfLock.Unlock()
			}
		case "string":
			{
				etdConfLock.Lock()
				otherConf[key] = string(v)
				etdConfLock.Unlock()
			}
		default:
			{
				return fmt.Errorf("json 没有相关配置 key:%s", key)
			}
		}
	}
	return nil
}

//初始化配置文件
func InitConf() error {
	parseCmd()
	err := ValidConf()
	if err != nil {
		fmt.Println("加载启动参数失败，err:", err)
		return err
	}

	f, err := os.OpenFile(CmdConf.Cfg, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开配置文件 '%s' 失败:%s", CmdConf.Cfg, err.Error())
	}
	defer f.Close()

	cfg := CustomCfg{}
	dec := json.NewDecoder(f)
	if err = dec.Decode(&cfg); err != nil {
		return fmt.Errorf("配置文件 '%s' 解析失败:%s", CmdConf.Cfg, err.Error())
	}
	for k, v := range cfg.Custom {
		etcdKey[k] = v
	}

	etcdWatch, err := NewEtcdWatcher(GetEtcdAddr())
	if err != nil {
		etcdWatch.ClearWatch()
		etcdWatch.Close(true)
		fmt.Println("加载启动参数失败，err:", err)
		return err
	}

	for k, _ := range etcdKey {
		//log.Println(k)
		etcdWatch.AddWatch(k, false, &listener{})
	}
	time.Sleep(1 * time.Second)

	return nil
}

func CloseEtcdWatch() {
	if etcdWatch != nil {
		etcdWatch.ClearWatch()
		etcdWatch.Close(true)
	}
}

func parseCmd() {
	AcPort := flag.Int("ac", 0, "access Listen Port")
	Ver := flag.String("v", "0.0.0.1", "Version")
	servName := flag.String("servName", "LoginService", "servName")
	ServID := flag.Int("servId", 1, "servId")
	IpAddress := flag.String("h", "127.0.0.1:8972", "IpAddress")
	EtcdAddress := flag.String("et", "127.0.0.1:12379", "IpAddress")
	Network := flag.String("net", "ws", "network")
	Level := flag.Int("l", 0, "Level")
	Type := flag.String("t", "H", "game类型 M-比赛 H-大厅") //默认为大厅使用的game
	plogConf := flag.String("plogCfg", "", "plog config file(ini)")
	Env := flag.Int("env", 0, "环境标识")
	Closed := flag.Bool("closed", false, "is service closed")
	ipCode := flag.String("ipcode", "", "ipcode file path")
	cfg := flag.String("cfg", "", "cfg file path")
	flag.Parse()
	args := new(CmdConfig)
	args.Version = *Ver
	args.ServerID = int32(*ServID)
	args.ServerName = *servName
	args.Level = int32(*Level)
	args.IpAddress = *IpAddress
	args.Type = *Type
	args.AcPort = *AcPort
	args.PLogConf = *plogConf
	args.ENV = int32(*Env)
	args.Network = *Network
	args.Closed = *Closed
	args.IpCode = *ipCode
	args.Cfg = *cfg
	args.EtcdAddress = *EtcdAddress
	CmdConf = args
	log.Println(args)
}

func ValidConf() error {
	//合法性校验
	if len(strings.Split(CmdConf.Version, ".")) != 4 {
		return fmt.Errorf("missing argument '-v' (version)")
	}
	if len(strings.Split(CmdConf.IpAddress, ".")) != 4 || len(strings.Split(CmdConf.IpAddress, ":")) != 2 {
		return fmt.Errorf("missing argument '-h' (listen)")
	}
	if CmdConf.ServerName == "" {
		return fmt.Errorf("missing argument '-name' (server name)")
	}
	if CmdConf.ServerID == 0 {
		return fmt.Errorf("missing argument '-servId' (server id)")
	}
	return nil
}

//func SetEtcdConf(r *serverplugin.EtcdV3RegisterPlugin)error{
//	if r==nil{
//		return fmt.Errorf("EtcdV3RegisterPlugin is empty")
//	}
//    etcdConf= r
//    return nil
//}

func GetDBConf(name string) *DBConf {

	ob, ok := mysqlConf[name]
	if ok {
		return ob
	}
	fmt.Println("the GetDBConf not exists!", name)
	//cli, err := clientv3.New(clientv3.Config{
	//	Endpoints: GetEtcdAddr(),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer cli.Close()
	//kvc := clientv3.NewKV(cli)
	//kv,err:=kvc.Get(context.TODO(), fmt.Sprintf("%s",name))
	//if err!=nil {
	//	return nil
	//}
	//if kv==nil{
	//	return nil
	//}
	//kvLen:=len(kv.Kvs)
	//if kvLen != 1{
	//	return nil
	//}
	//
	//var jsonValue bytes.Buffer
	//jsonValue.Write(kv.Kvs[0].Value)
	//
	//var conf DBConf
	//err = json.Unmarshal(jsonValue.Bytes(), &conf)
	//if err != nil {
	//	//log.Errorf("json.Unmarshal zk path %s: %v", string(jsonValue.Bytes()), err)
	//	fmt.Println(err)
	//	return nil
	//}
	//mysqlConf[name]=&conf
	return nil
}

func GetRedisPoolConf(name string) *RedisConf {
	//apolloConfLock.RLock()

	ob, ok := redisPoolConf[name]

	//apolloConfLock.RUnlock()
	if ok {
		return ob
	}
	fmt.Println("the GetRedisPoolConf not exists!", name)
	//cli, err := clientv3.New(clientv3.Config{
	//	Endpoints: GetEtcdAddr(),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer cli.Close()
	//kvc := clientv3.NewKV(cli)
	//kv,err:=kvc.Get(context.TODO(), fmt.Sprintf("%s",name))
	//if err!=nil {
	//	return nil
	//}
	//if kv==nil{
	//	return nil
	//}
	//kvLen:=len(kv.Kvs)
	//if kvLen != 1{
	//	return nil
	//}
	//
	//var jsonValue bytes.Buffer
	//jsonValue.Write(kv.Kvs[0].Value)
	//
	//var conf RedisConf
	//err = json.Unmarshal(jsonValue.Bytes(), &conf)
	//if err != nil {
	//	//log.Errorf("json.Unmarshal zk path %s: %v", string(jsonValue.Bytes()), err)
	//	fmt.Println(err)
	//	return nil
	//}
	//redisPoolConf[name]=&conf
	return nil
}

func GetBasePath() string {
	return GetStringConf("etcd.pro.base")
}

func GetEtcdAddr() (etcdAddr []string) {
	//zk的ip端口以逗号分隔
	theSlice := strings.Split(CmdConf.EtcdAddress, ",")
	if len(theSlice) <= 0 {
		return []string{""}
	}
	return theSlice
}

func GetNsqSrc() string {
	return GetStringConf("nsq.nsqd.n1")
}

func GetNSQLookupd() string {
	return GetStringConf("nsq.nslookupd.lk1")
}

func GetStringConf(name string) string {
	ob, ok := otherConf[name]

	//apolloConfLock.RUnlock()
	if ok {
		return ob
	}
	fmt.Println("the otherConf not exists!", name)
	return name
}
