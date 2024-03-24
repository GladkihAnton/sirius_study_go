package mail

type SendableI interface {
	GetFrom() string
	GetTo() string
}

type abstractSendable struct {
	from string
	to   string
}

func (sendable *abstractSendable) GetFrom() string {
	return sendable.from
}

func (sendable *abstractSendable) GetTo() string {
	return sendable.to
}
