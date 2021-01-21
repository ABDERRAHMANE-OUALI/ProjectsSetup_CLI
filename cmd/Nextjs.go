package cmd

import (
	"log"
	"os"
	"os/exec"
)

// Nextjs function
func Nextjs()  {
	cmd := exec.Command("/bin/sh", "-c", "yarn create next-app . && touch tsconfig.json && yarn add --dev typescript @types/react @types/node && yarn add sass")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return
	}
	
}
