package main

import (
	"regexp"

	"github.com/jukeizu/contract"
	cleverbot "github.com/ugjka/cleverbot-go"
)

type Server interface {
	Chat(contract.Request) (*contract.Response, error)
}

type server struct {
	cleverbot *cleverbot.Session
}

func NewServer(cleverBot *cleverbot.Session) Server {
	return &server{cleverBot}
}

func (s *server) Chat(request contract.Request) (*contract.Response, error) {
	r := regexp.MustCompile(`<[^<>]*>`)
	content := r.ReplaceAllString(request.Content, "")

	answer, err := s.cleverbot.Ask(content)
	if err != nil {
		return nil, err
	}

	message := contract.Message{
		Content: answer,
	}

	return &contract.Response{Messages: []*contract.Message{&message}}, nil
}
