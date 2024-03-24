package mail

type Message struct {
	message string

	abstractSendable
}

func (message *Message) GetMessage() string {
	return message.message
}
