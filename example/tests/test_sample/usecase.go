package usecases

type Approver interface {
	IsApproved() bool
}

type SendApprover struct {
	hasCompanyApproved bool
	hasTeamApproved    bool
	hasBossApproved    bool
}

type Sender interface {
	CanSend() bool
	SetMessage(message string)
	Send()
}

type MailSender struct {
	approver Approver
	email    string
	message  string
}

func (a *SendApprover) IsApproved() bool {
	return a.hasCompanyApproved && a.hasTeamApproved && a.hasBossApproved
}

func NewMailSender(approver Approver, email string) Sender {
	return &MailSender{approver: approver, email: email}
}

func (s *MailSender) CanSend() bool {
	return s.approver.IsApproved() && s.email != "" && s.message != ""
}

func (s *MailSender) SetMessage(message string) {
	s.message = message
}

func (s *MailSender) Send() {
	//
}
