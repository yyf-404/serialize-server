package service

import (
	"context"
	"fmt"
	serialize "github.com/yyf-404/serialize_pb/pb"
)

type SerializeService struct {
}

const Success = 0
const failure = 1

func (s *SerializeService) SendMsg(ctx context.Context, req *serialize.MsgRequest) (*serialize.MsgReply, error) {
	fmt.Printf("serializeService.SendMsg|req:%+v", req)
	rsp := &serialize.MsgReply{Code: Success, Message: "success"}
	if req.Name == "fail" {
		rsp.Code = failure
		rsp.Message = "failure"
	}
	return rsp, nil
}
