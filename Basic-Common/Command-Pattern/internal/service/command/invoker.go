package command

import (
	"fmt"
	"log"
)

type Invoker interface {
	Publish(command Command) error
}

type invoker struct {
}

func NewInvoker() *invoker {
	return &invoker{}
}

func (i *invoker) Publish(command Command) error {
	commandPayload, err := command.Serialize()
	if err != nil {
		log.Fatal(err)
	}

	/*
		TODO: Implement and DI SQS client in NewInvoker, then publish `commandPayload` to SQS
	*/
	fmt.Printf("Successfully published message!: %s\n", commandPayload)

	return nil
}
