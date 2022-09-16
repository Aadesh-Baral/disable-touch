package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func main() {
	drivers, err := ioutil.ReadFile("./driver_prefixs.txt")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	drivers_array := strings.Split(string(drivers), "\n")
	cmd := exec.Command("./scripts/list_input.sh")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		output_array := strings.Split(strings.ToLower(string(output)), "↳")
		driver_found := false
		for _, s := range output_array {
			for _, driver := range drivers_array {
				if strings.Contains(s, driver) {
					driver_found = true
					touch_id := strings.Split(s, "=")[1][:2]
					fmt.Println("Driver found: \n", strings.Split(s, "[")[0])
					command := exec.Command("./scripts/disable_touch.sh", touch_id)
					_, err = command.Output()
					if err != nil {
						fmt.Printf("error: %v", err)
					} else {
						fmt.Println("Successfully disabled driver", strings.Split(s, "[")[0])
					}
				}
			}
		}
		if !driver_found {
			fmt.Println("Driver not found! Please add it to driver_prefixs.txt")
		}
	}
}
