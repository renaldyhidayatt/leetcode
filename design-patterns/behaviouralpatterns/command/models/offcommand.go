package models

import "behaviouralpatterns/command/interfaces"

type OffCommand struct {
	Device interfaces.Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
}
