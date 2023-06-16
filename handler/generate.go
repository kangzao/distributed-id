package handler

import (
	"context"
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
	/**
	result, _ := redisUtil.Get(request.Stub)
	if result == nil {
		rsp.Start = 0
		rsp.End = COUNT
		redisUtil.Set(request.Stub, redisUtil.Int64ToBytes(rsp.End))
		return &rsp, nil
	}
	intResult := redisUtil.BytesToInt64(result)
	rsp.End = intResult + COUNT
	redisUtil.Set(request.Stub, redisUtil.Int64ToBytes(rsp.End))
	rsp.Start = intResult + 1
	return &rsp, nil
	*/
	end, _ := redisUtil.IncrBy(request.Stub, COUNT)
	rsp.End = end
	rsp.Start = end - COUNT + 1
	return &rsp, nil

}
