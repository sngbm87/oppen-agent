package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sngbm87/oppen-agent/internal/agent"
)

func main() {
	fmt.Println("Oppen Agent starting... YOLO Cosmic mode engaged.")
	// TODO: Initialize daemon, heartbeat, cron, model harness, etc.
	if err := agent.Run(); err != nil {
		log.Fatal(err)
	}
}