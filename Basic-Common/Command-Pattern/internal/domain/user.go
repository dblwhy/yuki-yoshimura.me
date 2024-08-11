package domain

import "yuki-yoshimura.me/command-pattern/internal/service/command"

type userService struct {
	commandInvoker command.Invoker
}

func NewUserDomain(invoker command.Invoker) *userService {
	return &userService{
		commandInvoker: invoker,
	}
}

func (u *userService) Create() error {
	/*
		TODO: Implement user domain logic here
	*/

	emailInfo := command.EmailInfo{From: "", To: []string{""}, Subject: "", Body: ""}
	u.commandInvoker.Publish(command.NewEmailCommand(emailInfo))

	return nil
}
