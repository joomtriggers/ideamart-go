package sms

import (
	"github.com/joomtriggers/ideamart/config"
	"github.com/joomtriggers/ideamart/sms/request"
)

type SenderInterface interface {
	Send() *SendResponse
	SMS() (*Sender, *request.SendRequest)
}

type Sender struct {
	request.SendRequest
	SendResponse
	config.Configuration
}

func (sender *Sender) Send() SendResponse {
	sender.SetApplicationId(sender.GetApplication())
	return sendRequest(sender)
}


