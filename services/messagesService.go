package services

import (
	"strings"

	"github.com/maurodanieldev/quasar-oper-fire/interfaces"
	"github.com/maurodanieldev/quasar-oper-fire/util"
)

type messagesService struct{}

func NewMessagesService() interfaces.IMessagesService {
	return &messagesService{}
}

func (s *messagesService) GetMessage(messages ...[]string) (msg string) {
	validate := true
	if len(messages) > 1 {
		length := len(messages[0])
		for index, _ := range messages {
			validate = (len(messages[index]) == length) && validate
		}
	}
	if validate && len(messages) > 1 {
		tmp := make([]string, len(messages[0]))
		for i, _ := range messages {
			tmp = util.GetPart(tmp, messages[i])
		}
		return strings.Join(tmp, " ")
	}
	return ""
}
