package main

func main() {
	tv := &Tv{}
	radio := &Radio{}
	airConditionner := &AirConditionner{}
	light := &Light{}
	hallogen := &HallogenLight{}

	devices := make([]Device, 0)
	devices = append(devices, tv)
	devices = append(devices, radio)
	devices = append(devices, airConditionner)

	lights := make([]Device, 0)
	lights = append(lights, light)
	lights = append(lights, hallogen)

	onCommand := &OnCommand{
		devices: devices,
	}
	offCommand := &OffCommand{
		devices: devices,
	}

	lightOn := &LightOnCommand{
		devices: lights,
	}
	lightOff := &LightOffCommand{
		devices: lights,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()

	lightsOnButton := &Button{
		command: lightOn,
	}
	lightsOnButton.press()

	lightsOffButton := &Button{
		command: lightOff,
	}
	lightsOffButton.press()
}
