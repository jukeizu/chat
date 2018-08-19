package main

import (
	"os"

	pb "github.com/jukeizu/chat/api/chat"
	"github.com/jukeizu/chat/listeners/commands"
	"github.com/jukeizu/handler"
	"github.com/shawntoffel/services-core/command"
	"github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/logging"
	"google.golang.org/grpc"
)

type CommandConfig struct {
	ChatServiceEndpoint string
	Handler             handler.HandlerConfig
}

var commandArgs command.CommandArgs

func init() {
	commandArgs = command.ParseArgs()
}

func main() {
	logger := logging.GetLogger("commands.chat", os.Stdout)

	commandConfig := CommandConfig{}
	err := config.ReadConfig(commandArgs.ConfigFile, &commandConfig)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(commandConfig.ChatServiceEndpoint, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := pb.NewChatClient(conn)
	command := chat.NewCommand(client)

	handler, err := handler.NewCommandHandler(logger, commandConfig.Handler)
	if err != nil {
		panic(err)
	}

	handler.Start(command)
}
