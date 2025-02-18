package main

import "fmt"

// Usage:
// GUI, Undo/Redo actions, Task Scheduling & Queueing, Workflow System

// Pros:
// Decouple sender and receiver
// Add new command without changing existing code
// Support undo/redo

// Cons:
// Each action requires a new command, increase complexity

func main() {
	light := &Light{} // Receiver

	onCommand := &LightOnCommand{light}
	offCommand := &LightOffCommand{light}

	invoker := &Invoker{} // Sender

	invoker.SetCommand(onCommand)
	invoker.Execute()

	invoker.SetCommand(offCommand)
	invoker.Execute()

	fmt.Println(light)

	invoker.Undo()

	fmt.Println(light)
}

type Command interface {
	Execute()
	Undo()
}

type Light struct {
	state bool
}

func (l *Light) On() {
	println("Light is on")
	l.state = true
}

func (l *Light) Off() {
	println("Light is off")
	l.state = false
}

type LightOnCommand struct {
	light *Light
}

func (l *LightOnCommand) Execute() {
	l.light.On()
}

func (l *LightOnCommand) Undo() {
	l.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func (l *LightOffCommand) Execute() {
	l.light.Off()
}

func (l *LightOffCommand) Undo() {
	l.light.On()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) Execute() {
	i.command.Execute()
}

func (i *Invoker) Undo() {
	i.command.Undo()
}
