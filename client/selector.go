package client

import (
	"context"
	"github.com/bitini111/rpcx/client"
	"github.com/valyala/fastrand"
)

type MySelector struct {
	servers map[string]string
	Dofunc  RouteCallback
}

func newMySelector(servers map[string]string) client.Selector {
	return &MySelector{servers: servers, Dofunc: nil}
}

func (s *MySelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	//ss := s.servers
	var ss []string
	if s.Dofunc != nil {
		dstSvid, _ := GetDstSvidByCTX(ctx)
		if dstSvid > 0 {
			//选择器从context中拿点对点请求的目标svid
			ss = s.Dofunc(s.servers, dstSvid)
		} else {
			//解决报错信息：指定svrId来找服务时出现异常:根据svid=0 找到的匹配目标地址为：[];所有的备选项有:map[string]string{\"tcp@127.0.0.1:30210\":\"1\"}",
			//2020-08-31调整：快速赛中(base/matchMngr/logic/logic.go RpcCall_MatchStart方法)同一个xclient可能会随机选择svrid也可能会根据指定的svrid来发起rpc调用。
			// 所以：当目标svrId为0的时候，就不要使用“指定svrId”的逻辑，而是改用告警并随机发送的方式！
			//ss = s.Dofunc(s.servers, args)
			ss = make([]string, 0, len(s.servers))
			for k := range s.servers {
				ss = append(ss, k)
			}
		}
	} else {
		ss = make([]string, 0, len(s.servers))
		for k := range s.servers {
			ss = append(ss, k)
		}
	}
	if len(ss) == 0 {
		return ""
	}

	i := fastrand.Uint32n(uint32(len(ss)))
	return ss[i]
}

func (s *MySelector) UpdateServer(servers map[string]string) {
	s.servers = make(map[string]string)
	for k, v := range servers {
		s.servers[k] = v
	}
	//plog.VFatal("【测试】UpdateServer更新之后的s.servers=%#v",s.servers)
}
