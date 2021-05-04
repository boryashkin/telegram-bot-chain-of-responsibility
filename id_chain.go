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
	next *MessageHandler
}

type UpdateMessageHandler struct {
	next *MessageHandler
}

type InlineQueryMessageHandler struct {
	next *MessageHandler
}

type ChannelPostMessageHandler struct {
	next *MessageHandler
}

type ChosenInlineResultMessageHandler struct {
	next *MessageHandler
}

type EditedChannelPostMessageHandler struct {
	next *MessageHandler
}

type EditedMessageMessageHandler struct {
	next *MessageHandler
}

type PreCheckoutQueryMessageHandler struct {
	next *MessageHandler
}

type ShippingQueryMessageHandler struct {
	next *MessageHandler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *CallbackQueryMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *UpdateMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *InlineQueryMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *ChannelPostMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *ChosenInlineResultMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *EditedChannelPostMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
}

func (h *EditedMessageMessageHandler) Handle(wrapper MessageWrapper) Result {
	if wrapper.Update.EditedMessage != nil {
		var chatID *int64
		if wrapper.Update.EditedMessage.Chat != nil {
			chatID = &wrapper.Update.EditedMessage.Chat.ID
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *EditedMessageMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *PreCheckoutQueryMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
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

	if h.next != nil {
		return (*h.next).Handle(wrapper)
	}

	return BaseResult{handled: false}
}
func (h *ShippingQueryMessageHandler) SetNext(handler MessageHandler) {
	h.next = &handler
}

func NewCallbackQueryMessageHandler() CallbackQueryMessageHandler {
	return CallbackQueryMessageHandler{}
}

func NewUpdateMessageHandler() UpdateMessageHandler {
	return UpdateMessageHandler{}
}

func NewInlineQueryMessageHandler() InlineQueryMessageHandler {
	return InlineQueryMessageHandler{}
}

func NewChannelPostMessageHandler() ChannelPostMessageHandler {
	return ChannelPostMessageHandler{}
}

func NewChosenInlineResultMessageHandler() ChosenInlineResultMessageHandler {
	return ChosenInlineResultMessageHandler{}
}

func NewEditedChannelPostMessageHandler() EditedChannelPostMessageHandler {
	return EditedChannelPostMessageHandler{}
}

func NewEditedMessageMessageHandler() EditedMessageMessageHandler {
	return EditedMessageMessageHandler{}
}

func NewPreCheckoutQueryMessageHandler() PreCheckoutQueryMessageHandler {
	return PreCheckoutQueryMessageHandler{}
}

func NewShippingQueryMessageHandler() ShippingQueryMessageHandler {
	return ShippingQueryMessageHandler{}
}

func NewFullChainMessageHandler() MessageHandler {
	base := NewEmptyHandler()

	cqm := NewCallbackQueryMessageHandler()
	umm := NewUpdateMessageHandler()
	iqm := NewInlineQueryMessageHandler()
	cpm := NewChannelPostMessageHandler()
	cir := NewChosenInlineResultMessageHandler()
	ecp := NewEditedChannelPostMessageHandler()
	emm := NewEditedMessageMessageHandler()
	pcq := NewPreCheckoutQueryMessageHandler()
	sqm := NewShippingQueryMessageHandler()

	cqm.SetNext(&umm)
	umm.SetNext(&iqm)
	iqm.SetNext(&cpm)
	cpm.SetNext(&cir)
	cir.SetNext(&ecp)
	ecp.SetNext(&emm)
	pcq.SetNext(&sqm)
	sqm.SetNext(&base)

	return &cqm
}
