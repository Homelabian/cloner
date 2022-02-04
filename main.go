package main

import (
	"os"
	"os/signal"
	"strconv"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

var AllJobs map[int]Job
var CurrentFlags RunningOptions

func init() {
	// Build the flag options
	// We should start with this to catch help requests
	CurrentFlags = initFlags()

	// Build the jobs index
	AllJobs = make(map[int]Job)
	loadJobs()
	buildCloneOptions()
	validateJobs()
}

func main() {
	// Start the Cron
	c := cron.New()

	// Is this for a single job?
	if CurrentFlags.OneJobID != -1 {
		// This is a single job run
		if job, ok := AllJobs[CurrentFlags.OneJobID]; ok {
			log.Info("Cloning Job with ID " + strconv.Itoa(CurrentFlags.OneJobID))
			clone(job)
			scheduleJob(c, job)
		} else {
			log.Fatal("Single Job ID is not found")
		}
	} else {
		// This is a run of the entire jobs list
		for _, job := range AllJobs {
			log.Info("Cloning All Jobs")
			clone(job)
			scheduleJob(c, job)
		}
	}

	// Is this a single run or Cron run
	if CurrentFlags.SingleRun {
		log.Info("All clones completed")
		os.Exit(0)
	} else {
		log.Info("Starting Cron")
		c.Start()
	}

	// Register the signal to kill
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
