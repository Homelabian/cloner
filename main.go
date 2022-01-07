package main

import (
	"os"
	"os/signal"
	"strconv"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

var AllJobs map[int]Job

func init() {
	AllJobs = make(map[int]Job)
	loadJobs()
	buildCloneOptions()
	validateJobs()
}

func main() {
	log.Info("Performing Initial Clone of all Jobs")
	// Perform initial Run of all jobs
	for _, j := range AllJobs {
		clone(j)
	}

	// Create the CRON
	c := cron.New()

	// Schedule all Jobs in CRON
	for i, j := range AllJobs {
		log.Info("Scheduling Job: " + strconv.Itoa(i))
		c.AddFunc(j.Cron, func() { clone(j) })
	}

	// Start the Cron
	log.Info("Starting Cron")
	c.Start()

	// Register signal to kill
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
