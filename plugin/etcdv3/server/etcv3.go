package serverplugin

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rpcxio/libkv"
	"github.com/rpcxio/libkv/store"
	estore "github.com/rpcxio/rpcx-etcd/store"
	etcd "github.com/rpcxio/rpcx-etcd/store/etcdv3"
	"github.com/smallnest/rpcx/log"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	metrics "github.com/rcrowley/go-metrics"
)

func init() {
	etcd.Register()
}

//是否开启清除老节点
var isOpenCleanOldNodeInfo = false

type etcdServIDVers struct {
	ID   int64        `json:"ID"`
	Vers []etcdVersIP `json:"Vers"`
}

type etcdVersIP struct {
	Ver    string `json:"Ver"`
	IP     string `json:"IP"`
	Unique string `json:"Unique"`
}

// EtcdV3RegisterPlugin implements etcd registry.
type EtcdV3RegisterPlugin struct {
	// service address, for example, tcp@127.0.0.1:8972, quic@127.0.0.1:1234
	ServiceAddress string
	// etcd addresses
	EtcdServers []string
	// base path for rpcx server, for example com/example/rpcx
	BasePath string
	Metrics  metrics.Registry
	// Registered services
	// Registered services
	ServiceName string

	Services       []string
	metasLock      sync.RWMutex
	metas          map[string]string
	UpdateInterval time.Duration

	Options *store.Config
	kv      store.Store

	dying chan struct{}
	done  chan struct{}

	//the version of SERVER
	Version string
	//the ID of SERVER
	ServerID int32

	//Server Name path
	serverNamePath string

	//unique id mark regitser
	uniqueId string
}

// Start starts to connect etcd cluster
func (p *EtcdV3RegisterPlugin) Start() error {

	if p.Version == "" {
		log.Errorf("The Version of Server is NULL")
		return fmt.Errorf("The Version of Server is NULL")
	} else {
		ss := strings.Split(p.Version, ".")
		if len(ss) != 4 {
			log.Errorf("The Version Format is NOT CORRECT, Use: x.x.x.x")
			return fmt.Errorf("The Version Format is NOT CORRECT, Use: x.x.x.x")
		}
	}

	if p.serverNamePath == "" {
		p.serverNamePath = "L"
	}

	if p.ServerID == 0 {
		log.Errorf("The ID of Server is NULL")
		return errors.New("The ID of Server is NULL")
	}

	if p.done == nil {
		p.done = make(chan struct{})
	}
	if p.dying == nil {
		p.dying = make(chan struct{})
	}

	if p.kv == nil {
		kv, err := libkv.NewStore(estore.ETCDV3, p.EtcdServers, p.Options)
		if err != nil {
			log.Errorf("cannot create etcd registry: %v", err)
			return err
		}
		p.kv = kv
	}

	if p.BasePath[0] == '/' {
		p.BasePath = p.BasePath[1:]
	}

	err := p.kv.Put(p.BasePath, []byte("rpcx_path"), &store.WriteOptions{IsDir: true})
	if err != nil && !strings.Contains(err.Error(), "Not a file") {
		log.Errorf("cannot create etcd path %s: %v", p.BasePath, err)
		return err
	}

	return nil
}

// Stop unregister all services.
func (p *EtcdV3RegisterPlugin) Stop() error {
	close(p.dying)
	<-p.done
	if p.kv == nil {
		kv, err := libkv.NewStore(estore.ETCDV3, p.EtcdServers, p.Options)
		if err != nil {
			log.Errorf("cannot create etcd registry: %v", err)
			return err
		}
		p.kv = kv
	}

	if p.BasePath[0] == '/' {
		p.BasePath = p.BasePath[1:]
	}

	for _, name := range p.Services {
		strs := strings.Split(name, "#")
		if len(strs) != 2 {
			continue
		}

		err := p.delVerIPAddrJson(strs[0], strs[1])
		if err != nil {
			log.Errorf("cannot delete delVerIPAddrJson zk path %s:%s, %v", strs[0], strs[1], err)
			continue
		}
	}

	close(p.dying)
	<-p.done
	return nil
}

func (p *EtcdV3RegisterPlugin) EnableOpenCleanOldNode() {
	isOpenCleanOldNodeInfo = true
}

func (p *EtcdV3RegisterPlugin) addVerIPAddrJson(path string, serverId int32, version string, IPAddr string) error {
	var exist bool
	var err error
	exist, err = p.kv.Exists(path)
	if err != nil {
		log.Errorf("Exist err zk path %s: %v", path, err)
		return err
	}
	var jsonIDVerIpAddr []byte

	var needNotify bool
	if exist {
		//Find the ids list of json
		//ps, err1 := p.kv.List(nodePath)
		ps, err := p.kv.Get(path)
		if err != nil {
			log.Errorf("p.kv.List zk path %s: %v", path, err)
			return err
		}
		//log.Errorf("ps.len %d", len(ps))
		var jsonValue bytes.Buffer
		jsonValue.Write(ps.Value)

		//Parse json to Struct
		var etcdServID etcdServIDVers
		err = json.Unmarshal(jsonValue.Bytes(), &etcdServID)
		if err != nil {
			log.Errorf("json.Unmarshal zk path %s: %v", string(jsonValue.Bytes()), err)
			return err
		}

		//Replace the exist IPAddr
		var Vers []etcdVersIP
		for _, pv := range etcdServID.Vers {
			if pv.Ver != version {
				Vers = append(Vers, pv)
			} else {
				if !isOpenCleanOldNodeInfo && pv.IP == IPAddr {
					//do NOTHING zk has exist nodePath
					goto doNothing
				}
				if pv.IP == IPAddr {
					arr := strings.Split(IPAddr, "@")
					if len(arr) != 2 {
						//do NOTHING zk has exist nodePath
						goto doNothing
					}
					//检测老的服务连接是否存活
					if c, err := net.DialTimeout(arr[0], arr[1], time.Second*3); err != nil {
						//连接超时，此时处理方式需要待确定，有可能网络抖动导致3秒内连接不上，此处发送的概率较低
						//目前的方式记录下错误信息，同时替换节点信息
						if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
							log.Errorf("connect timeout %s\n", err)
						}
						log.Errorf("old ip=%s return connect error %s\n", IPAddr, err)
						continue //服务不存在,删除节点
					} else {
						//老服务存在，不需要添加节点
						c.Close()
						goto doNothing
					}
				}
				log.Errorf("nothing change add new ip address")
				//IP地址不相同则继续添加新的节点
			}
			/* --- 保证修改唯一值为自己所拥有,不再跳出
			else {
				if pv.IP == IPAddr {
					//do NOTHING zk has exist nodePath
					//goto doNothing
				}
			}
			*/
		}

		etcdServID.Vers = Vers
		//Add new IPAddr
		verNode := &etcdVersIP{Ver: version, IP: IPAddr, Unique: p.uniqueId}
		etcdServID.Vers = append(etcdServID.Vers, *verNode)

		//排序只取最高5个版本
		if num := len(etcdServID.Vers); num > 5 {
			sort.Slice(etcdServID.Vers, func(i, j int) bool {
				return p.ver2int(etcdServID.Vers[i].Ver) < p.ver2int(etcdServID.Vers[j].Ver)
			})
			etcdServID.Vers = etcdServID.Vers[num-5:]
		}

		jsonIDVerIpAddr, err = json.Marshal(etcdServID)
		if err != nil {
			log.Errorf("json.Marshal json: %s: %v", string(jsonIDVerIpAddr), err)
			return err
		}
		needNotify = true
	} else {

		zkVers := &etcdVersIP{Ver: version, IP: IPAddr, Unique: p.uniqueId}
		zkServID := &etcdServIDVers{}
		zkServID.ID = int64(serverId)
		zkServID.Vers = append(zkServID.Vers, *zkVers)
		jsonIDVerIpAddr, err = json.Marshal(zkServID)
		if err != nil {
			log.Errorf("json.Marshal json: %s: %v", string(jsonIDVerIpAddr), err)
			return err
		}
		//Add a new node to zk path will automatic notify, NO need manual do it
		needNotify = false
	}

	err = p.kv.Put(path, jsonIDVerIpAddr, &store.WriteOptions{TTL: p.UpdateInterval * 2})
	if err != nil {
		log.Errorf("cannot create zk path %s: %v", path, err)
		return err
	}

	if needNotify {
		err = p.notifyWatcher(p.BasePath, p.ServiceName)
		if err != nil {
			log.Errorf("cannot notifyWatcher zk path %s: %v", path, err)
			return err
		}
	}

doNothing:
	return nil
}

func (p *EtcdV3RegisterPlugin) delVerIPAddrJson(path string, version string) error {
	var exist bool
	var err error
	exist, err = p.kv.Exists(path)
	if err != nil {
		log.Errorf("Exist err zk path %s: %v", path, err)
		return err
	}

	if exist {

		var jsonIDVerIpAddr []byte
		//Find the ids list of json
		//ps, err1 := p.kv.List(nodePath)
		ps, err := p.kv.Get(path)
		if err != nil {
			log.Errorf("p.kv.List zk path %s: %v", path, err)
			return err
		}

		var jsonValue bytes.Buffer
		jsonValue.Write(ps.Value)

		//Parse json to Struct
		var zkServID etcdServIDVers
		err = json.Unmarshal(jsonValue.Bytes(), &zkServID)
		if err != nil {
			log.Errorf("json.Unmarshal zk path %s: %v", string(jsonValue.Bytes()), err)
			return err
		}

		//Replace the exist IPAddr
		var Vers []etcdVersIP
		for _, pv := range zkServID.Vers {
			if pv.Ver != version {
				Vers = append(Vers, pv)
			}
		}

		if len(zkServID.Vers) == len(Vers) {
			goto doNothing
		}

		zkServID.Vers = Vers

		jsonIDVerIpAddr, err = json.Marshal(zkServID)
		if err != nil {
			log.Errorf("json.Marshal json: %s: %v", string(jsonIDVerIpAddr), err)
			return err
		}

		err = p.kv.Put(path, jsonIDVerIpAddr, &store.WriteOptions{TTL: p.UpdateInterval * 2})
		if err != nil {
			log.Errorf("cannot create zk path %s: %v", path, err)
			return err
		}
		err = p.notifyWatcher(p.BasePath, p.ServiceName)
		if err != nil {
			log.Errorf("cannot notifyWatcher zk path %s: %v", path, err)
			return err
		}
	}
doNothing:
	return nil
}

func (p *EtcdV3RegisterPlugin) notifyWatcher(base string, name string) error {

	nodePath := fmt.Sprintf("%s/%s/%s/%d", base, "L", name, -1)

	var exist bool
	var err error
	exist, err = p.kv.Exists(nodePath)
	_, err = p.kv.Exists(nodePath)
	if err != nil {
		log.Errorf("Exist err zk path %s: %v", nodePath, err)
		return err
	}

	if exist {
		err = p.kv.Delete(nodePath)
		if err != nil {
			log.Errorf("Delete err zk path %s: %v", nodePath, err)
			return err
		}
	} else {
		err = p.kv.Put(nodePath, []byte(""), &store.WriteOptions{IsDir: true})
		if err != nil {
			log.Errorf("Delete err zk path %s: %v", nodePath, err)
			return err
		}
	}
	//err = p.kv.Put(nodePath, []byte(""), &store.WriteOptions{IsDir: true})
	//if err != nil {
	//	log.Errorf("Delete err zk path %s: %v", nodePath, err)
	//	return err
	//}

	return nil
}

// HandleConnAccept handles connections from clients
func (p *EtcdV3RegisterPlugin) HandleConnAccept(conn net.Conn) (net.Conn, bool) {
	if p.Metrics != nil {
		metrics.GetOrRegisterMeter("connections", p.Metrics).Mark(1)
	}
	return conn, true
}

// PreCall handles rpc call from clients
func (p *EtcdV3RegisterPlugin) PreCall(_ context.Context, _, _ string, args interface{}) (interface{}, error) {
	if p.Metrics != nil {
		metrics.GetOrRegisterMeter("calls", p.Metrics).Mark(1)
	}
	return args, nil
}

// Register handles registering event.
// this service is registered at BASE/serviceName/thisIpAddress node
func (p *EtcdV3RegisterPlugin) Register(name string, rcvr interface{}, metadata string) (err error) {
	if strings.TrimSpace(name) == "" {
		err = errors.New("Register service `name` can't be empty")
		return
	}

	if p.kv == nil {
		etcd.Register()
		kv, err := libkv.NewStore(estore.ETCDV3, p.EtcdServers, nil)
		if err != nil {
			log.Errorf("cannot create etcd registry: %v", err)
			return err
		}
		p.kv = kv
	}

	if p.BasePath[0] == '/' {
		p.BasePath = p.BasePath[1:]
	}

	p.ServiceName = name
	text := fmt.Sprintf("%d[%s-%d-%s-%s]", time.Now().UnixNano(), p.Version, p.ServerID, p.ServiceName, p.ServiceAddress)
	p.uniqueId = fmt.Sprintf("%x", md5.Sum([]byte(text)))
	p.uniqueId = fmt.Sprintf("%s-%s", time.Now().Format("20060102 15:04:05"), p.uniqueId)

	err = p.kv.Put(p.BasePath, []byte("rpcx_path"), &store.WriteOptions{IsDir: true})
	if err != nil && !strings.Contains(err.Error(), "Not a file") {
		log.Errorf("cannot create etcd path %s: %v", p.BasePath, err)
		return err
	}
	//
	nodePath := fmt.Sprintf("%s/%s/%s", p.BasePath, p.serverNamePath, name)
	err = p.kv.Put(nodePath, []byte(name), &store.WriteOptions{IsDir: true})
	if err != nil && !strings.Contains(err.Error(), "Not a file") {
		log.Errorf("cannot create etcd path %s: %v", nodePath, err)
		return err
	}

	nodePath = fmt.Sprintf("%s/%s/%s/%d", p.BasePath, p.serverNamePath, name, p.ServerID)
	err = p.addVerIPAddrJson(nodePath, p.ServerID, p.Version, p.ServiceAddress)
	if err != nil {
		log.Errorf("addVerIPAddrJson err etcd path %s: %v", nodePath, err)
		return err
	}
	Fname := nodePath + "#" + p.Version
	p.Services = append(p.Services, Fname)

	p.metasLock.Lock()
	if p.metas == nil {
		p.metas = make(map[string]string)
	}
	p.metas[Fname] = metadata
	p.metasLock.Unlock()
	return
}

func (p *EtcdV3RegisterPlugin) RegisterFunction(serviceName, fname string, fn interface{}, metadata string) error {
	return p.Register(serviceName, fn, metadata)
}

func (p *EtcdV3RegisterPlugin) Unregister(name string) (err error) {
	if len(p.Services) == 0 {
		return nil
	}

	if strings.TrimSpace(name) == "" {
		err = errors.New("Register service `name` can't be empty")
		return
	}

	if p.kv == nil {
		etcd.Register()
		kv, err := libkv.NewStore(estore.ETCDV3, p.EtcdServers, nil)
		if err != nil {
			log.Errorf("cannot create etcd registry: %v", err)
			return err
		}
		p.kv = kv
	}

	err = p.kv.Put(p.BasePath, []byte("rpcx_path"), &store.WriteOptions{IsDir: true})
	if err != nil && !strings.Contains(err.Error(), "Not a file") {
		log.Errorf("cannot create etcd path %s: %v", p.BasePath, err)
		return err
	}

	nodePath := fmt.Sprintf("%s/%s", p.BasePath, name)
	err = p.kv.Put(nodePath, []byte(name), &store.WriteOptions{IsDir: true})
	if err != nil && !strings.Contains(err.Error(), "Not a file") {
		log.Errorf("cannot create etcd path %s: %v", nodePath, err)
		return err
	}

	nodePath = fmt.Sprintf("%s/%s/%s", p.BasePath, name, p.ServiceAddress)

	err = p.kv.Delete(nodePath)
	if err != nil {
		log.Errorf("cannot create consul path %s: %v", nodePath, err)
		return err
	}

	if len(p.Services) > 0 {
		var services = make([]string, 0, len(p.Services)-1)
		for _, s := range p.Services {
			if s != name {
				services = append(services, s)
			}
		}
		p.Services = services
	}

	p.metasLock.Lock()
	if p.metas == nil {
		p.metas = make(map[string]string)
	}
	delete(p.metas, name)
	p.metasLock.Unlock()
	return
}

func (p *EtcdV3RegisterPlugin) GetkeyString(path string) (error, []byte) {
	var (
		exist     bool
		err       error
		jsonValue bytes.Buffer
	)

	if path[0] == '/' {
		path = path[1:]
	}

	exist, err = p.kv.Exists(path)
	if err != nil {
		log.Errorf("Exist err zk path %s: %v", path, err)
		return err, jsonValue.Bytes()
	}
	if !exist {
		return fmt.Errorf("%s is empty", path), jsonValue.Bytes()
	}

	//Find the ids list of json
	//ps, err1 := p.kv.List(nodePath)
	ps, err := p.kv.Get(path)
	if err != nil {
		log.Errorf("p.kv.List zk path %s: %v", path, err)
		return err, nil
	}

	jsonValue.Write(ps.Value)
	return nil, jsonValue.Bytes()

}

//版本号转成数字
func (p *EtcdV3RegisterPlugin) ver2int(ver string) int64 {
	ret, fields := int64(0), strings.SplitN(ver, ".", 4)
	if n, err := strconv.Atoi(fields[0]); err == nil && n > 0 && n < 65536 {
		ret |= int64(n) << 48
	}
	if n, err := strconv.Atoi(fields[1]); err == nil && n > 0 && n < 65536 {
		ret |= int64(n) << 32
	}
	if n, err := strconv.Atoi(fields[2]); err == nil && n > 0 && n < 65536 {
		ret |= int64(n) << 16
	}
	if n, err := strconv.Atoi(fields[3]); err == nil && n > 0 && n < 65536 {
		ret |= int64(n)
	}
	return ret
}
