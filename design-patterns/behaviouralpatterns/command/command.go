package command

import "behaviouralpatterns/command/models"

func Commant() {
	tv := &models.Tv{}
	onCommand := &models.OnCommand{
		Device: tv,
	}

	offCommand := &models.OffCommand{
		Device: tv,
	}

	onButton := &models.Button{
		Command: onCommand,
	}

	onButton.Press()

	offButton := &models.Button{
		Command: offCommand,
	}

	offButton.Press()
}
