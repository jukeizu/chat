package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jukeizu/contract"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	cleverbot "github.com/ugjka/cleverbot-go"
)

var Version = ""

var (
	flagPort = "10000"
)

func parseCli() {
	flag.StringVar(&flagPort, "p", flagPort, "port")
	flag.Parse()
}

func main() {
	parseCli()
	port := ":" + flagPort

	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("instance", xid.New().String()).
		Str("component", "intent.endpoint.chat").
		Str("version", Version).
		Logger()

	filename := os.Getenv("CLEVERBOT_TOKEN_FILE")
	tokenBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error().Err(err).
			Str("filename", filename).
			Msg("could not read secrets file")
		return
	}

	cleverBot := cleverbot.New(string(tokenBytes))
	s := NewServer(cleverBot)
	s = NewServerLogger(s, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/chat", contract.MakeHttpHandlerFunc(s.Chat))

	logger.Info().Str("address", port).Msg("listening")
	http.ListenAndServe(port, mux)
}
