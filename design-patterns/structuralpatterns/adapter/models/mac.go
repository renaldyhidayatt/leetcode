package models

import "fmt"

type Mac struct {
}

func (m *Mac) InsertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}
