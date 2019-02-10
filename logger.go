package main

import (
	"time"

	"github.com/jukeizu/contract"
	"github.com/rs/zerolog"
)

type serverLogger struct {
	server Server
	logger zerolog.Logger
}

func NewServerLogger(server Server, logger zerolog.Logger) Server {
	return &serverLogger{server, logger}
}

func (l *serverLogger) Chat(request contract.Request) (response *contract.Response, err error) {
	defer func(begin time.Time) {
		l := l.logger.With().
			Str("intent", "chat").
			Str("took", time.Since(begin).String()).
			Logger()

		if err != nil {
			l.Error().Err(err).Msg("")
			return
		}

		l.Info().Msg("called")
	}(time.Now())

	response, err = l.server.Chat(request)

	return
}
