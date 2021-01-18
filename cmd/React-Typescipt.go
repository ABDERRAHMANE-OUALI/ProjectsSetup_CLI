package cmd

import (
	"fmt"
	"os/exec"
)

// ReactTypescript function
func ReactTypescript()  {
	out, err := exec.Command("npx create-react-app my-app --template typescript").Output()
	
	if err != nil {
		// colorize the output string
		fmt.Println(string("\033[31m"), "Operation failed %v\n", err, string("\033[0m"))
		return
	}
	
	output := string(out[:])
    fmt.Println(output)

}
