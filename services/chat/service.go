package chat

import (
	"context"
	"regexp"

	cleverbot "github.com/ugjka/cleverbot-go"

	pb "github.com/jukeizu/chat/api/chat"
)

type service struct {
	Cleverbot *cleverbot.Session
}

func NewService(cleverbotSession *cleverbot.Session) pb.ChatServer {
	return &service{cleverbotSession}
}

func (s service) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageReply, error) {
	r := regexp.MustCompile(`<[^<>]*>`)
	content := r.ReplaceAllString(req.Message, "")

	answer, err := s.Cleverbot.Ask(content)

	return &pb.SendMessageReply{Message: answer}, err
}
