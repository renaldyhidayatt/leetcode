package models

import "structuralpatterns/adapter/interfaces"

type Client struct {
}

func (c *Client) InsertSquareUsbInComputer(com interfaces.Computer) {
	com.InsertInSquarePort()
}
