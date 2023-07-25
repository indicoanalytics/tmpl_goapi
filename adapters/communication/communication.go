package communication

import (
	"api.default.indicoinnovation.pt/clients/mailgun"
	"api.default.indicoinnovation.pt/config/constants"
	"api.default.indicoinnovation.pt/entity"
)

type Communication struct {
	mailgun *mailgun.Mailgun
}

func New() *Communication {
	return &Communication{
		mailgun: mailgun.New(),
	}
}

func (comm *Communication) parseProvider(channel string) func(string, *entity.MessageAttributes) {
	if channel == constants.ChannelEmail && constants.EmailProvider == "mailgun" {
		return func(x string, y *entity.MessageAttributes) {
			comm.mailgun.Send(x, y)
		}
	}

	return func(x string, y *entity.MessageAttributes) {}
}

func (comm *Communication) Send(to string, channel string, messageAttr *entity.MessageAttributes) {
	comm.parseProvider(channel)(to, messageAttr)
}
