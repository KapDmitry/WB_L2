package main

import "fmt"

// Command - интерфейс команды
type Command interface {
	Execute()
}

// Receiver - получатель команд
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver executing action")
}

// ConcreteCommand - конкретная реализация команды
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

// Invoker - вызывающий объект
type Invoker struct {
	command Command
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) SetCommand(cmd Command) {
	i.command = cmd
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &Receiver{}
	command := NewConcreteCommand(receiver)
	invoker := NewInvoker()
	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
