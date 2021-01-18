package cmd

import (
	"fmt"
	"os/exec"
)

// ReactTypescript function
func ReactTypescript()  {
	out, err := exec.Command("/bin/sh", "-c", "npx create-react-app my-app --template typescript && npm i bootstrap && cd $1/src && mkdir styles/ components/ && rm logo.svg App.css reportWebVitals.ts setupTests.ts && mv ./index.css ./styles/ && cd styles/ && touch index.scss && cd .. && sed -i '/reportWebVitals/d' index.tsx && printf '%s\n' 3 i 'import 'bootstrap/dist/css/bootstrap.css';' . w q | ed index.tsx").Output()
	
	if err != nil {
		// colorize the output string
		fmt.Println(string("\033[31m"), "Operation failed %v\n", err, string("\033[0m"))
		return
	}
	
	output := string(out[:])
    fmt.Println(output)

}
