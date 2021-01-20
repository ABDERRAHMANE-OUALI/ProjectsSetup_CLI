package cmd

import (
	"fmt"
	"os/exec"
)

// FirebaseFunctions function
func FirebaseFunctions()  {
	out, err := exec.Command("firebase init functions && cd functions && npm i firebase-admin@latest firebase-functions@latest").Output()
	if err != nil {
		// colorize the output string
		fmt.Println(string("\033[31m"), "Operation failed %v\n", err, string("\033[0m"))
		return
	}
	
	output := string(out[:])
    fmt.Println(output)

}
