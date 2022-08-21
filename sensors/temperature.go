package sensors

import (
	"fmt"

	"github.com/ssimunic/gosensors"
)

func Temperature() {
	fmt.Println("Start temperature sensor 🌡️ ...")
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		if err.Error() == "lm-sensors missing" {
			fmt.Println("Error: you have to install \"lm-sensors\" tool with \"sudo apt-get install lm-sensors\"")
		} else {
			fmt.Println("An error is occured with temperature sensor")
		}
		panic(err)
	}

	fmt.Println(sensors.JSON())
}
