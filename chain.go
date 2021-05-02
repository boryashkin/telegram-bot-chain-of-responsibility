package telegram_bot_chain_of_responsibility

import "github.com/go-telegram-bot-api/telegram-bot-api"

type MessageWrapper struct {
	Update tgbotapi.Update
}

type MessageHandler interface {
	SetNext(MessageHandler)
	GetNext() MessageHandler
	Handle(wrapper MessageWrapper) Result
}

type BaseHandler struct {
	next MessageHandler
}

func NewBaseHandler() BaseHandler {
	return BaseHandler{}
}

func (h *BaseHandler) Handle(wrapper MessageWrapper) Result {

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *BaseHandler) SetNext(MessageHandler) {
	return
}

func (h *BaseHandler) GetNext() MessageHandler {
	return h.next
}
