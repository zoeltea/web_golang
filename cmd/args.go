package cmd

import (
	"flag"
	"fmt"
	"os"
)

func ParseArgs() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	// Define command line flags
	envFile := flag.String("env", ".env", "Path to .env file")
	port := flag.String("port", "", "Server port (overrides .env)")
	migrate := flag.Bool("migrate", false, "Run database migrations")

	flag.Parse()

	// Handle flags
	if *envFile != ".env" {
		os.Setenv("ENV_FILE", *envFile)
	}

	if *port != "" {
		os.Setenv("SERVER_PORT", *port)
	}

	if *migrate {
		// Add your migration logic here
		fmt.Println("Running migrations...")
		os.Exit(0)
	}
}
