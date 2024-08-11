package lambda

import (
	"context"
	"fmt"

	"yuki-yoshimura.me/command-pattern/internal/service/command"
)

var commandUnmarshaler = map[command.CommandType]command.Unmarshaler{
	command.CommandTypeEmail: command.UnmarshalEmailCommand,
	// Add more command types here as needed
}

type lambdaCommandHandler struct{}

func NewLambdaCommandHandler() *lambdaCommandHandler {
	return &lambdaCommandHandler{}
}

func (l *lambdaCommandHandler) HandleRequest(ctx context.Context, event *command.CommandPayload) error {
	if event == nil {
		return fmt.Errorf("received nil event")
	}

	deserializer, exists := commandUnmarshaler[event.Type]
	if !exists {
		return fmt.Errorf("received unexpected command type: %s", string(event.Type))
	}
	command, err := deserializer(event.Body)
	if err != nil {
		return err
	}

	return command.Execute()
}
