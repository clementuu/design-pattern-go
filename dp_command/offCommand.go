package main

type OffCommand struct {
	devices []Device
}

func (c *OffCommand) execute() {
	for _, device := range c.devices {
		device.off()
	}
}
