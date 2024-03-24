package mail

type ServiceI interface {
	ProcessMail(sendable SendableI) SendableI
}

type RealService struct{}

func (service RealService) processMail(mail SendableI) SendableI {
	return mail
}
