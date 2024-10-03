package main

type LightOffCommand struct {
	devices []Device
}

func (c *LightOffCommand) execute() {
	for _, device := range c.devices {
		device.off()
	}
}
