package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
)

type Job struct {
	ID     int
	Cron   string
	Repo   Repo
	Output string
}

type CRON struct {
	Original string
}

type Repo struct {
	UseAuth      bool
	Username     string
	Password     string
	URL          string
	CloneOptions *git.CloneOptions
}

func getEnv() map[string]string {
	// Get all environment variables
	envs := os.Environ()
	newEnv := make(map[string]string)

	// Convert them to map of names and values
	for _, env := range envs {
		parts := strings.Split(env, "=")
		name := parts[0]
		value := parts[1]
		newEnv[name] = value
	}
	return newEnv
}

func loadJobs() {
	log.Info("Loading Jobs")
	// Get the formatted array of ENV Variables
	envs := getEnv()

	// Environment Variable Format:
	//
	// CLONER_1_CRON
	//    |   |   |
	//    |   |   |- Setting Specific
	//    |   |- Job ID
	//    |- Cloner Identifier

	// Recognised Environment Variables
	// CRON
	// REPO
	// REPOAUTH
	// REPOUSER
	// REPOPASS
	// OUTPUT

	// Split Env Variables into components
	confs := make(map[int]Job)
	for name, value := range envs {
		log.Trace("Evaluating ENV: " + name)
		name_parts := strings.Split(name, "_")

		// Is this variable part of our config?
		if name_parts[0] != "CLONER" {
			continue
		}

		// Set up some working variables
		id, err := strconv.Atoi(name_parts[1])
		if err != nil {
			log.Error(err)
		}
		setting := name_parts[2]

		// Blank job or pull existing if there is one
		job := Job{ID: id}
		if val, ok := confs[id]; ok {
			job = val
		}

		// Figure out what setting it is we are dealing with and do it
		switch setting {
		case "CRON":
			job.Cron = value
		case "REPO":
			job.Repo.URL = value
		case "REPOAUTH":
			asBool, err := strconv.ParseBool(value)
			if err != nil {
				log.Error(err)
			}
			job.Repo.UseAuth = asBool
		case "REPOUSER":
			job.Repo.Username = value
		case "REPOPASS":
			job.Repo.Password = value
		case "OUTPUT":
			job.Output = value
		default:
			log.Error("Unrecognised ENV Setting - " + setting + ", with value " + value)
		}

		// Push the Job back to the config
		confs[id] = job
	}

	// Push back into master jobs list
	AllJobs = confs
	count := strconv.Itoa(len(confs))

	log.Info("Loaded " + count + " Jobs")
}

func buildCloneOptions() {
	// Loop through all registered jobs and build the clone Options Object
	for i, j := range AllJobs {
		co := git.CloneOptions{}

		if j.Repo.UseAuth {
			co.URL = j.Repo.URL
			auth := http.BasicAuth{}
			auth.Username = j.Repo.Username
			auth.Password = j.Repo.Password
			co.Auth = &auth
		} else {
			co.URL = j.Repo.URL
		}

		// Push back to main Config
		j.Repo.CloneOptions = &co
		AllJobs[i] = j
	}
}

func validateJobs() {
	log.Info("Validating Jobs")
	for i, j := range AllJobs {
		changed := false

		// Ensure that the output directory is set
		if j.Output == "" {
			dir := "/" + strconv.Itoa(i) + "-output"
			log.Warning("No Output Directory set on Job " + strconv.Itoa(i) + ", using " + dir)
			j.Output = dir
			changed = true
		}

		// Ensure that the CRON is set to something
		if j.Cron == "" {
			log.Warning("No Cron set on Job" + strconv.Itoa(i) + ", using '0 0 0 * * *' (Daily @Midnight)")
			j.Cron = "0 0 0 * * *"
			changed = true
		}

		// Push back to all Jobs if Changed
		if changed {
			AllJobs[i] = j
		}
	}
}
