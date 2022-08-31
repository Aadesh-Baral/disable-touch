package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("./scripts/list_input.sh")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		output_array := strings.Split(strings.ToLower(string(output)), "↳")
		for _, s := range output_array {
			if strings.Contains(s, "touchscreen") {
				touch_id := strings.Split(s, "=")[1][:2]
				// fmt.Println(touch_id)
				command := exec.Command("./scripts/disable_touch.sh", touch_id)
				_, err = command.Output()
				if err != nil {
					fmt.Printf("error: %v", err)
				}
			}
		}

	}
}
