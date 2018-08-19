package chat

import (
	"context"

	pb "github.com/jukeizu/chat/api/chat"
	"github.com/jukeizu/handler"
)

type Command interface {
	handler.Command
}

type command struct {
	Client pb.ChatClient
}

func NewCommand(client pb.ChatClient) Command {
	return &command{client}
}

func (c *command) IsCommand(request handler.Request) (bool, error) {
	for _, mention := range request.Mentions {
		if mention.Id == request.Bot.Id {
			return true, nil
		}
	}

	return false, nil
}

func (c *command) Handle(request handler.Request) (handler.Results, error) {
	chatRequest := &pb.SendMessageRequest{Message: request.Content}

	reply, err := c.Client.SendMessage(context.Background(), chatRequest)
	if err != nil {
		return handler.Results{}, err
	}

	result := handler.Result{
		Content: reply.Message,
	}

	return handler.Results{result}, nil
}
