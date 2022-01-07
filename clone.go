package main

import (
	"os"
	"strconv"

	git "github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
)

func clone(job Job) {
	// Log Started
	log.Info("Started Cloning Job: " + strconv.Itoa(job.ID))

	// Does the directory exist?
	if _, err := os.Stat(job.Output); os.IsNotExist(err) {
		log.Warning("Output Directory doesn't exist, may be incorrect settings. Continuing with clone.")
	} else {
		log.Debug("Removing Existing Output Directory")
		err := os.RemoveAll(job.Output)
		if err != nil {
			log.Error(err)
		}
	}

	// Do the clone
	_, err := git.PlainClone(job.Output, false, job.Repo.CloneOptions)
	if err != nil {
		log.Error(err)
	}

	// Log complete
	log.Info("Finnished Clonning")
}
