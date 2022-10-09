package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/bitini111/rpcx/share"
	"github.com/bitini111/rpcxapp/comm"
	"strconv"
	"strings"
)

//rpc协议中,透传map中的自定义字段名称.
const (
	//rpc透传时,map中的field名称.(放到const中主要是为了避免滥用冲突)
	ReqMetaDataKey_FIELD_srcName   = "srcName"   //传递rpc请求源的名称
	ReqMetaDataKey_FIELD_srcSvid   = "srcSvid"   //传递rpc请求源的svrid
	ReqMetaDataKey_FIELD_clientSeq = "clientSeq" //传递客户端请求序列号,rpc服务在必要时可以快速响应该seq的请求.(比如 game服快速响应操作)
	ReqMetaDataKey_FIELD_clientUid = "clientUid" //rpc调用时,传递客户端的这个socket连接对应的uid,便于业务做安全性检测
	ReqMetaDataKey_FIELD_dstSvid   = "dstSvid"   //rpcClient中记录下目标svrid,以便于选择合适的服务端
	ReqMetaDataKey_FIELD_remoteIp  = "remoteIp"  //请求方的ip地址；access转入登录的时候会用到。
)

//封装好发起rpc调用的客户端！
func XCall(ctx context.Context, serviceName string, funcName string, args TheMarshaler, reply interface{}) error {
	ctx = addMeatDataToCTX(ctx, 0)
	xclient, e := MyClients.getClient(serviceName, nil, false)
	if e != nil {
		return e
	}
	err := xclient.Call(ctx, funcName, args, reply)
	//此处可能不是某个连接关闭，只能是整个rpcx全部连接不可用才会报错ErrXClientShutdown

	//if err != nil && err == client.ErrXClientShutdown {
	if err != nil {
		MyClients.delClient(serviceName)
	}
	return err
}

//封装好发起rpc调用的客户端！(args和reply都是字节切片的格式)，以插件的形式实现！
func XCallBytes(ctx context.Context, serviceName string, funcName string, args interface{}, reply interface{}, cb RouteCallback) error {
	ctx = addMeatDataToCTX(ctx, 0)
	xclient, e := MyClientsWithBytes.getClient(serviceName, cb, true)
	if e != nil {
		return e
	}
	err := xclient.Call(ctx, funcName, args, reply)
	return err
}

//封装好发起rpc调用的客户端,且指定目标svrid！(args和reply都是字节切片的格式)，以插件的形式实现！
func XCallBytesWithSvid(ctx context.Context, serviceName string, funcName string, args interface{}, reply interface{}, dstSvid int32) error {
	ctx = addMeatDataToCTX(ctx, dstSvid)
	xclient, e := MyClientsWithBytes.getClient(serviceName, routeSelectBySvid, true)
	if e != nil {
		return e
	}
	err := xclient.Call(ctx, funcName, args, reply)
	return err
}

func XCallWithRoutefunc(ctx context.Context, serviceName string, funcName string, args TheMarshaler, reply interface{}, cb RouteCallback) error {
	ctx = addMeatDataToCTX(ctx, 0)
	//plog.VInfo("xcall 发起的时间为：%v",time.Now())
	xclient, e := MyClients.getClient(serviceName, cb, false)
	if e != nil {
		return e
	}
	//plog.VInfo("获取到的xclient时间为：%v",time.Now())
	err := xclient.Call(ctx, funcName, args, reply)
	//plog.VInfo("xcall结束的时间为：%v",time.Now())
	return err
}

//根据指定的svid发起rpc调用
func XCallWithSvid(ctx context.Context, serviceName string, funcName string, args TheMarshaler, reply interface{}, dstSvid int32) error {
	//设置目标svid的值
	ctx = addMeatDataToCTX(ctx, dstSvid)
	xclient, e := MyClients.getClient(serviceName, routeSelectBySvid, false)
	if e != nil {
		return e
	}
	err := xclient.Call(ctx, funcName, args, reply)
	//此处可能不是某个连接关闭，只能是整个rpcx全部连接不可用才会报错ErrXClientShutdown
	//if err != nil && err == client.ErrXClientShutdown {
	if err != nil {
		MyClients.delClient(serviceName)
	}
	return err
}

//根据svid进行rpc服务的路由选择func
func routeSelectBySvid(ss map[string]string, args interface{}) []string {
	var svid int32
	var ok bool
	ts := make([]string, 0, len(ss))

	if args != nil {
		//根据svid查找对应的地址
		svid, ok = args.(int32)
		if ok && svid > 0 {
			var strSvid string
			for k, v := range ss {
				if pos := strings.LastIndex(v, "/"); pos > -1 {
					if pos < len(v)-1 {
						strSvid = fmt.Sprintf("%s", v[pos+1:])
					}
				} else {
					strSvid = v
				}
				if strSvid == fmt.Sprintf("%d", svid) {
					ts = append(ts, k)
					break
				}
			}
		}
	}
	//
	if len(ts) == 0 {
		//plog.VFatal("指定svrId来找服务时出现异常:根据svid=%d 找到的匹配目标地址为：%v;所有的备选项有:%#v", svid, ts, ss)
	}
	return ts
}

//向context中追加数据(以达到rpc框架底层透传一些公共数据的目的:见rpcx issue209)
//透传的信息包括源服务名，源svid，目标seq。
func addMeatDataToCTX(ctx context.Context, dstSvid int32) context.Context {
	//发起rpc调用的时候，把rpc客户端的name和id带过去
	srcName := comm.Cfg.Base.ServerName
	srcSvid := comm.Cfg.Base.ServerID

	mp, ok := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	if !ok || mp == nil {
		mp = make(map[string]string)
	}
	//var mp  = make(map[string]string)
	if len(srcName) > 0 {
		mp[ReqMetaDataKey_FIELD_srcName] = srcName
	}
	if srcSvid > 0 {
		mp[ReqMetaDataKey_FIELD_srcSvid] = fmt.Sprintf("%d", srcSvid)
	}
	if dstSvid > 0 {
		mp[ReqMetaDataKey_FIELD_dstSvid] = fmt.Sprintf("%d", dstSvid)
	}

	nctx := context.WithValue(ctx, share.ReqMetaDataKey, mp)
	return nctx
}

//获取context透传下来的dstSvid
//发起指定了svrid的调用时，会用到该方法。
func GetDstSvidByCTX(ctx context.Context) (int32, error) {
	val, ok := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	if !ok {
		return 0, errors.New("no dst found!")
	}
	if val[ReqMetaDataKey_FIELD_dstSvid] == "0" || val[ReqMetaDataKey_FIELD_dstSvid] == "" {
		return 0, errors.New("no dst found!")
	}
	tSvid, err := strconv.Atoi(val[ReqMetaDataKey_FIELD_dstSvid])
	if err != nil {
		return 0, errors.New("illegal dst!")
	}
	return int32(tSvid), nil
}

//把seq也透传到下一个调用中去。
func NewContextWithSeq(ctx context.Context, seq int32) context.Context {
	if seq > 0 {
		return AddValueToCTX(ctx, ReqMetaDataKey_FIELD_clientSeq, fmt.Sprintf("%d", seq))
	}
	return ctx
}

//把uid也透传到下一个调用中去。
func NewContextWithUid(ctx context.Context, uid int64) context.Context {
	if uid > 0 {
		return AddValueToCTX(ctx, ReqMetaDataKey_FIELD_clientUid, fmt.Sprintf("%d", uid))
	}
	return ctx
}

//从context中获取本次rpc来源的uid
func GetClientUidByCTX(ctx context.Context) (int32, error) {
	val, ok := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	if !ok {
		return 0, errors.New("no uid found!")
	}
	//if val[gDefine.ReqMetaDataKey_FIELD_clientUid] == "0" || val[gDefine.ReqMetaDataKey_FIELD_clientUid] == ""{
	//	return 0,errors.New("no uid found!")
	//}
	tSvid, err := strconv.Atoi(val[ReqMetaDataKey_FIELD_clientUid])
	if err != nil {
		return 0, errors.New("illegal uid!")
	}
	return int32(tSvid), nil
}

//把ip地址透传到下一个调用中去。
func NewContextWithRemoteIp(ctx context.Context, remoteIp string) context.Context {
	if len(remoteIp) > 0 {
		return AddValueToCTX(ctx, ReqMetaDataKey_FIELD_remoteIp, remoteIp)
	}
	return ctx
}

//从context中获取本次rpc来源的ip地址
func GetRemoteIpByCTX(ctx context.Context) (string, error) {
	val, ok := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	if !ok {
		return "", errors.New("no remoteIp found!")
	}
	remoteIp, ok := val[ReqMetaDataKey_FIELD_remoteIp]
	return remoteIp, nil
}

func routeSelect(ss map[string]string, args interface{}) []string {
	ts := make([]string, 0, len(ss))
	for k := range ss {
		ts = append(ts, k)
	}
	if args != nil {
		//fmt.Println(args.(*OperationReq).Uid)
	}
	//fmt.Println(ts)
	return ts
}

func AddValueToCTX(ctx context.Context, key, data string) context.Context {
	val, ok := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	if !ok || val == nil {
		val = make(map[string]string)
	}
	val[key] = data

	return context.WithValue(ctx, share.ReqMetaDataKey, val)
}

func GetValueFromCTX(ctx context.Context, key string) (string, bool) {
	val, ok := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	if !ok {
		return "", false
	}
	data, ok := val[key]
	return data, ok
}
