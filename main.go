package main

import (
	"log"
	"login/core"
	"login/core/database"
	"login/core/master"
	"login/core/models"
	"os"
	"os/exec"
	"time"

	"github.com/janeczku/go-spinner"
)

func main() {
	// Start the configs
	s := spinner.StartNew("Initializing")
	core.Initialize()

	// Initialize the database
	if err := database.New(); err != nil {
		log.Println("failed to initialize database", err)
		return
	}
	s.Stop()
	time.Sleep(5 * time.Millisecond)
	clearScreen()

	// Start the webserver
	s = spinner.StartNew("| Webserver Started | ")
	log.Print(models.Config.Secure)
	master.NewV2()
	s.Stop()
}

//clear screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
