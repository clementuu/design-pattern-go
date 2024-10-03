package main

type OnCommand struct {
	devices []Device
}

func (c *OnCommand) execute() {
	for _, device := range c.devices {
		device.on()
	}
}
