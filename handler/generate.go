package handler

import (
	"context"
	"distributed-id/global"
	"distributed-id/proto"
	redisUtil "distributed-id/redis"
)

type GenerateIDServer struct {
	proto.UnimplementedGenerateIdServer
}

const COUNT = 10000

func (g *GenerateIDServer) GetIds(ctx context.Context, request *proto.GetIdRequest) (*proto.GetIdResponse, error) {
	//获取一批id
	var rsp proto.GetIdResponse
	//如果结果为空，则说明是第一次访问该接口
	result, _ := redisUtil.Get(request.Stub)

	//sales:10000
	//完成初始化
	if result == nil {
		rsp.Start = 1
		rsp.End = COUNT                                              //[1,10000]
		redisUtil.Set(request.Stub, redisUtil.Int64ToBytes(rsp.End)) //记录最新的值
		return &rsp, nil                                             //返回号段
	}
	/**
	//第二次请求  10000
	intResult := redisUtil.BytesToInt64(result)
	rsp.End = intResult + COUNT                                  //20000
	redisUtil.Set(request.Stub, redisUtil.Int64ToBytes(rsp.End)) //记录最新的值
	rsp.Start = intResult + 1                                    //设置区间
	return &rsp, nil

	1、get key from redis  2、update value to redis  3、return [start,end]
	远程调用 - 序列化 反序列化 网络连接 传输
	*/

	end, _ := redisUtil.IncrBy(request.Stub, global.ServerConfig.Count)
	rsp.End = end
	rsp.Start = end - global.ServerConfig.Count + 1
	return &rsp, nil

}
