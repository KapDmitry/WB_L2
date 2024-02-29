package main

import (
	"fmt"
)

type MessageHandler interface {
	SetNext(MessageHandler) MessageHandler
	HandleMessage(string) string
}

type TextMessageHandler struct {
	next MessageHandler
}

func (c *TextMessageHandler) SetNext(next MessageHandler) MessageHandler {
	c.next = next
	return next
}

func (c *TextMessageHandler) HandleRequest(request string) string {
	if len(request) < 10 {
		return fmt.Sprintf("Request %s handled by TextMessageHandler", request)
	}
	if c.next != nil {
		return c.next.HandleMessage(request)
	}
	return "End of chain, request not handled"
}

type AudioMessageHandler struct {
	next MessageHandler
}

func (c *AudioMessageHandler) SetNext(next MessageHandler) MessageHandler {
	c.next = next
	return next
}

func (c *AudioMessageHandler) HandleMessage(request string) string {
	if len(request) >= 10 && len(request) < 20 {
		return fmt.Sprintf("Request %s handled by ConcreteHandlerB", request)
	}
	if c.next != nil {
		return c.next.HandleMessage(request)
	}
	return "End of chain, request not handled"
}

type Message struct {
	bytes string
}

func main() {
	handlerA := &TextMessageHandler{}
	handlerB := &AudioMessageHandler{}
	msg := Message{bytes: "asdpasdas"}

	handlerA.SetNext(handlerB)
	handlerA.next.HandleMessage(msg.bytes)

}
