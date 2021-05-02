package telegram_bot_chain_of_responsibility

type ChatIDHolder struct {
	handled            bool
	ChatID             *int64
	MessageID          *int
	InlineMessageID    *string
	CallbackQueryID    *string
	PreCheckoutQueryID *string
	ShippingQueryID    *string
}

func (r ChatIDHolder) IsHandled() bool {
	return r.handled
}

type CallbackQueryMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type UpdateMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type InlineQueryMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type ChannelPostMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type ChosenInlineResultMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type EditedChannelPostMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type EditedMessageMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type PreCheckoutQueryMessageHandler struct {
	BaseHandler
	next MessageHandler
}

type ShippingQueryMessageHandler struct {
	BaseHandler
	next MessageHandler
}

func (h *CallbackQueryMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.CallbackQuery != nil {
		var chatID *int64
		var messageID *int
		if wrapper.Update.CallbackQuery.Message != nil {
			messageID = &wrapper.Update.CallbackQuery.Message.MessageID
			if wrapper.Update.CallbackQuery.Message.Chat != nil {
				chatID = &wrapper.Update.CallbackQuery.Message.Chat.ID
			}
		}
		return ChatIDHolder{
			true,
			chatID,
			messageID,
			nil,
			&wrapper.Update.CallbackQuery.ID,
			nil,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *UpdateMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.Message != nil {
		var chatID *int64
		if wrapper.Update.Message.Chat != nil {
			chatID = &wrapper.Update.Message.Chat.ID
		}
		return ChatIDHolder{
			true,
			chatID,
			&wrapper.Update.Message.MessageID,
			nil,
			nil,
			nil,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *InlineQueryMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.InlineQuery != nil {
		return ChatIDHolder{
			true,
			nil,
			nil,
			&wrapper.Update.InlineQuery.ID,
			nil,
			nil,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *ChannelPostMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.ChannelPost != nil {
		var chatID *int64
		if wrapper.Update.ChannelPost.Chat != nil {
			chatID = &wrapper.Update.ChannelPost.Chat.ID
		}
		return ChatIDHolder{
			true,
			chatID,
			&wrapper.Update.ChannelPost.MessageID,
			nil,
			nil,
			nil,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *ChosenInlineResultMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.ChosenInlineResult != nil {
		return ChatIDHolder{
			true,
			nil,
			nil,
			&wrapper.Update.ChosenInlineResult.InlineMessageID,
			nil,
			nil,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *EditedChannelPostMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.EditedChannelPost != nil {
		var chatID *int64
		if wrapper.Update.EditedChannelPost.Chat != nil {
			chatID = &wrapper.Update.EditedChannelPost.Chat.ID
		}
		return ChatIDHolder{
			true,
			chatID,
			&wrapper.Update.EditedChannelPost.MessageID,
			nil,
			nil,
			nil,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *EditedMessageMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.EditedMessage != nil {
		var chatID *int64
		if wrapper.Update.EditedMessage.Chat != nil {
			chatID = &wrapper.Update.EditedChannelPost.Chat.ID
		}
		return ChatIDHolder{
			true,
			chatID,
			&wrapper.Update.EditedMessage.MessageID,
			nil,
			nil,
			nil,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *PreCheckoutQueryMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.PreCheckoutQuery != nil {
		return ChatIDHolder{
			true,
			nil,
			nil,
			nil,
			nil,
			&wrapper.Update.PreCheckoutQuery.ID,
			nil,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func (h *ShippingQueryMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.ShippingQuery != nil {
		return ChatIDHolder{
			true,
			nil,
			nil,
			nil,
			nil,
			nil,
			&wrapper.Update.ShippingQuery.ID,
		}
	}

	if h.GetNext() != nil {
		return h.GetNext().Handle(wrapper)
	}

	return BaseResult{handled: false}
}

func NewCallbackQueryMessageHandler(handler BaseHandler, nextHandler MessageHandler) CallbackQueryMessageHandler {
	return CallbackQueryMessageHandler{handler, nextHandler}
}

func NewUpdateMessageHandler(handler BaseHandler, nextHandler MessageHandler) UpdateMessageHandler {
	return UpdateMessageHandler{handler, nextHandler}
}

func NewInlineQueryMessageHandler(handler BaseHandler, nextHandler MessageHandler) InlineQueryMessageHandler {
	return InlineQueryMessageHandler{handler, nextHandler}
}

func NewChannelPostMessageHandler(handler BaseHandler, nextHandler MessageHandler) ChannelPostMessageHandler {
	return ChannelPostMessageHandler{handler, nextHandler}
}

func NewChosenInlineResultMessageHandler(handler BaseHandler, nextHandler MessageHandler) ChosenInlineResultMessageHandler {
	return ChosenInlineResultMessageHandler{handler, nextHandler}
}

func NewEditedChannelPostMessageHandler(handler BaseHandler, nextHandler MessageHandler) EditedChannelPostMessageHandler {
	return EditedChannelPostMessageHandler{handler, nextHandler}
}

func NewEditedMessageMessageHandler(handler BaseHandler, nextHandler MessageHandler) EditedMessageMessageHandler {
	return EditedMessageMessageHandler{handler, nextHandler}
}

func NewPreCheckoutQueryMessageHandler(handler BaseHandler, nextHandler MessageHandler) PreCheckoutQueryMessageHandler {
	return PreCheckoutQueryMessageHandler{handler, nextHandler}
}

func NewShippingQueryMessageHandler(handler BaseHandler, nextHandler MessageHandler) ShippingQueryMessageHandler {
	return ShippingQueryMessageHandler{handler, nextHandler}
}
