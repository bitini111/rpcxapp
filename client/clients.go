package client

import (
	"fmt"
	"github.com/bitini111/rpcx/client"
	"github.com/bitini111/rpcx/plugin/etcdv3/clientplugin"
	"github.com/bitini111/rpcxapp/comm"
	"sync"
)

//封装了rpc客户端map的结构体
type Clients struct {
	mp   map[string]client.XClient
	mpcb map[string]RouteCallback
	etcd *clientplugin.EtcdV3Discovery
	mu   sync.RWMutex
}

var once sync.Once

func (c *Clients) getClient(serviceName string, cb RouteCallback, withBytes bool) (client.XClient, error) {
	once.Do(
		func() {
			BasePath = comm.Cfg.Base.RpcPath
			etcAddr = comm.Cfg.Base.EtcdAddress
		})
	c.mu.RLock()

	cc, ok := c.mp[serviceName]
	if ok {
		if cb != nil {
			//如果cb存在，而 mp map和 mpcb map里没有，则需要添加进去
			cachecb := c.mpcb[serviceName]
			if cachecb == nil {
				c.mu.RUnlock()
				c.mu.Lock()
				serverList := make(map[string]string)
				mySelect := newMySelector(serverList)
				if cb != nil {
					mySelect.(*MySelector).Dofunc = cb
					c.mpcb[serviceName] = cb
				}
				c.mp[serviceName].SetSelector(mySelect)
				c.mu.Unlock()
			} else {
				c.mu.RUnlock()
			}
		} else {
			c.mu.RUnlock()
		}

		return cc, nil
	} else {
		c.mu.RUnlock()      //解读锁
		c.mu.Lock()         //加写锁
		defer c.mu.Unlock() //解写锁
		//还需要再判断一次
		if cc, ok := c.mp[serviceName]; ok {
			if cb != nil {
				cachecb := c.mpcb[serviceName]
				if cachecb == nil {
					//此处已经有锁，不必再加
					serverList := make(map[string]string)
					mySelect := newMySelector(serverList)
					if cb != nil {
						mySelect.(*MySelector).Dofunc = cb
						c.mpcb[serviceName] = cb
					}
					c.mp[serviceName].SetSelector(mySelect)
				}
			}
			return cc, nil
		}
		var d client.ServiceDiscovery
		var er error
		if c.etcd != nil {
			d = c.etcd.Copy(BasePath, serviceName)
			//尝试判断是否关闭，如果关闭则新建连接
			//if c.etcd.IsClosed(BasePath, serviceName) {
			//	fmt.Printf("old etcd may be disconnected try create new zk")
			//	d, er = etcd_client.NewEtcdV3Discovery(BasePath, serviceName, etcAddr, false,nil)
			//	if er != nil {
			//		return nil, er
			//	}
			//	c.etcd = d.(*etcd_client.EtcdV3Discovery)
			//}
			if d == nil {
				return nil, fmt.Errorf(fmt.Sprintf("not found service[%s] in etcd", serviceName))
			}
		} else {
			//d := client.NewZookeeperDiscovery(BasePath, serviceName, ZkAddr, nil)
			d, er = clientplugin.NewEtcdV3Discovery(BasePath, serviceName, etcAddr, false, nil)
			if er != nil {
				return nil, er
			}
			c.etcd = d.(*clientplugin.EtcdV3Discovery)
		}
		var xclient client.XClient
		var plugins = client.NewPluginContainer()
		if withBytes {
			//透传使用的客户端，需要加插件
			xclient = client.NewXClient(serviceName, FailMode, SelectMode, d, MyOptions)
		} else {
			xclient = client.NewXClient(serviceName, FailMode, SelectMode, d, NormalOption)
		}
		//链路追踪插件
		xclient.SetPlugins(plugins)

		//用自己的Selector
		serverList := make(map[string]string)
		mySelect := newMySelector(serverList)
		if cb != nil {
			mySelect.(*MySelector).Dofunc = cb
			c.mpcb[serviceName] = cb
		}
		xclient.SetSelector(mySelect)

		c.mp[serviceName] = xclient
		return xclient, nil
	}
	return nil, fmt.Errorf("no client found")
}

func (c *Clients) delClient(serviceName string) {
	c.mu.Lock()         //加写锁
	defer c.mu.Unlock() //解写锁
	delete(c.mp, serviceName)
	delete(c.mpcb, serviceName)
	fmt.Printf("close ser=%s rpcx client", serviceName)
}
