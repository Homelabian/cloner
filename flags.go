package main

import "flag"

type RunningOptions struct {
	SingleRun bool
	OneJobID  int
}

func initFlags() RunningOptions {
	// Register all our flags
	Single := flag.Bool("single", false, "When set to true, do not start Cron, Run as a single Sync only")
	Job := flag.Int("id", -1, "Only sync this job")

	// Parse all our flags
	flag.Parse()

	// Build and Return the Struct
	OptionsOut := RunningOptions{
		SingleRun: *Single,
		OneJobID:  *Job,
	}
	return OptionsOut
}
