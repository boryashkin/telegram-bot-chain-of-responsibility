package telegram_bot_chain_of_responsibility

import "github.com/go-telegram-bot-api/telegram-bot-api"

type MessageWrapper struct {
	Update tgbotapi.Update
}

type Result interface {
	IsHandled() bool
}

type BaseResult struct {
	handled bool
}

func (r BaseResult) IsHandled() bool {
	return r.handled
}

type MessageHandler interface {
	SetNext(MessageHandler)
	Handle(wrapper MessageWrapper) Result
}

type EmptyHandler struct {
	next *MessageHandler
}

func NewEmptyHandler() EmptyHandler {
	return EmptyHandler{}
}

func (h *EmptyHandler) Handle(wrapper MessageWrapper) Result {
	return BaseResult{handled: false}
}

func (h *EmptyHandler) SetNext(handler MessageHandler) {
	h.next = &handler
}
