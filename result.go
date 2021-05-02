package telegram_bot_chain_of_responsibility

type Result interface {
	IsHandled() bool
}

type BaseResult struct {
	handled bool
}

func (r BaseResult) IsHandled() bool {
	return r.handled
}
