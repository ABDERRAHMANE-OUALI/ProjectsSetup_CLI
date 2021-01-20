package cmd

import (
	"fmt"
	"os/exec"
)

// Nextjs function
func Nextjs()  {
	out, err := exec.Command("/bin/sh", "-c", "yarn create next-app . && touch tsconfig.json && yarn add --dev typescript @types/react @types/node").Output()
	
	if err != nil {
		// colorize the output string
		fmt.Println(string("\033[31m"), "Operation failed %v\n", err, string("\033[0m"))
		return
	}
	
	output := string(out[:])
    fmt.Println(output)

}
