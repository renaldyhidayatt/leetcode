package models

import "behaviouralpatterns/command/interfaces"

type Button struct {
	Command interfaces.Command
}

func (b *Button) Press() {
	b.Command.Execute()
}
