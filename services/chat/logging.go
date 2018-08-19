package chat

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	pb "github.com/jukeizu/chat/api/chat"
)

type loggingService struct {
	logger  log.Logger
	Service pb.ChatServer
}

func NewLoggingService(logger log.Logger, s pb.ChatServer) pb.ChatServer {
	return &loggingService{logger, s}
}

func (s *loggingService) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (reply *pb.SendMessageReply, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SendMessage",
			"request", *req,
			"reply", *reply,
			"took", time.Since(begin),
		)

	}(time.Now())

	reply, err = s.Service.SendMessage(ctx, req)

	return
}
