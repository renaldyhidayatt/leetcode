package models

import "behaviouralpatterns/command/interfaces"

type OnCommand struct {
	Device interfaces.Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
}
