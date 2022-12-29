package models

type WindowsAdapter struct {
	WindowMachine *Windows
}

func (w *WindowsAdapter) InsertInSquarePort() {
	w.WindowMachine.InsertInCirclePort()
}
