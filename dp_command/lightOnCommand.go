package main

type LightOnCommand struct {
	devices []Device
}

func (c *LightOnCommand) execute() {
	for _, device := range c.devices {
		device.on()
	}
}
