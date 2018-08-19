package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/jukeizu/chat/api/chat"
	"github.com/jukeizu/chat/services/chat"
	"github.com/shawntoffel/services-core/command"
	"github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/logging"
	cleverbot "github.com/ugjka/cleverbot-go"
	"google.golang.org/grpc"
)

type ServiceConfig struct {
	Port            int
	CleverbotApiKey string
}

var serviceArgs command.CommandArgs

func init() {
	serviceArgs = command.ParseArgs()
}

func main() {
	logger := logging.GetLogger("services.chat", os.Stdout)

	serviceConfig := ServiceConfig{}

	err := config.ReadConfig(serviceArgs.ConfigFile, &serviceConfig)

	if err != nil {
		panic(err)
	}

	cleverbotSession := cleverbot.New(serviceConfig.CleverbotApiKey)

	service := chat.NewService(cleverbotSession)
	service = chat.NewLoggingService(logger, service)

	errChannel := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errChannel <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		port := fmt.Sprintf(":%d", serviceConfig.Port)

		listener, err := net.Listen("tcp", port)

		if err != nil {
			errChannel <- err

		}

		s := grpc.NewServer()
		pb.RegisterChatServer(s, service)

		logger.Log("transport", "grpc", "address", port, "msg", "listening")

		errChannel <- s.Serve(listener)
	}()

	logger.Log("stopped", <-errChannel)
}
