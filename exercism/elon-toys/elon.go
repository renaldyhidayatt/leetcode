package elontoys

import "fmt"

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Drive %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	return c.battery/c.batteryDrain*c.speed >= trackDistance
}
