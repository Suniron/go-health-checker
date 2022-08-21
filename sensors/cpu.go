package sensors

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func Cpu() {
	u, err := cpu.Percent(100, false)
	if err != nil {
		panic(nil)
	}
	fmt.Println(u)
}
