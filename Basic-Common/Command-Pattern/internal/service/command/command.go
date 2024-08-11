package command

type Command interface {
	Serialize() (string, error)
	Execute() error
}

type CommandType string

type CommandPayload struct {
	Type CommandType `json:"type"`
	Body string      `json:"body"`
}

type Unmarshaler func(string) (Command, error)
