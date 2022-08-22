package sensors

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxT = 70000 // 70° celsius

func getThermalZoneFilePaths() (files []string, err error) {
	baseDir := "/sys/class/thermal/"

	allFiles, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return files, err
	}

	for _, f := range allFiles {
		if strings.Contains(f.Name(), "thermal_zone") {
			files = append(files, baseDir+f.Name())
		}
	}

	return files, err
}

func getTemperaturesFromPaths(paths []string) (t []int, err error) {
	for _, p := range paths {
		dat, err := os.ReadFile(p + "/temp")
		if err != nil {
			return t, err
		}

		temp, err := strconv.Atoi(strings.TrimSpace(string(dat)))
		if err != nil {
			return t, err
		}

		t = append(t, temp)
	}

	return t, err

}

func isTemperaturesGood(t []int) (ok bool) {
	for _, temp := range t {
		if temp > maxT { // TODO: use user max temperature
			fmt.Printf("Warn: %v° is to hot !\n", temp/1000)
			ok = false
		}
	}

	return ok
}

func Temperature() {
	fmt.Println("Start temperature sensor 🌡️ ...")

	p, err := getThermalZoneFilePaths()
	if err != nil {
		log.Fatal(err)
	}

	t, err := getTemperaturesFromPaths(p)
	if err != nil {
		log.Fatal(err)
	}

	isTemperaturesGood(t)

}
