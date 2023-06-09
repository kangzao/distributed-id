package api

import (
	"context"
	proto "distributed-id/proto"
	redisUtil "distributed-id/redis"
)

type GenerateIdServer struct {
	proto.UnimplementedGenerateIdServer
}

func (generateIdServer *GenerateIdServer) GetIds(ctx context.Context, request *proto.GetIdRequest) (*proto.GetIdResponse, error) {
	var rsp proto.GetIdResponse
	//如果结果为空，则说明是第一次访问该接口
	result, _ := redisUtil.Get(request.Stub)
	if result == nil {
		rsp.Start = 0
		rsp.End = int64(request.Count)
		redisUtil.Set(request.Stub, redisUtil.Int64ToBytes(rsp.End))
		return &rsp, nil
	}
	intResult := redisUtil.BytesToInt64(result)
	rsp.End = intResult + request.Count
	redisUtil.Set(request.Stub, redisUtil.Int64ToBytes(rsp.End))
	rsp.Start = intResult + 1
	return &rsp, nil
}
