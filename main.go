package main

import (
	"flag"
	"log"
	"time"

	osquery "github.com/kolide/osquery-go"
	"github.com/kolide/osquery-go/plugin/table"
)

func main() {
	var (
		flSocketPath = flag.String("socket", "", "")
		flTimeout    = flag.Int("timeout", 0, "")
		_            = flag.Int("interval", 0, "")
		_            = flag.Bool("verbose", false, "")
	)
	flag.Parse()

	// allow for osqueryd to create the socket path otherwise it will error
	time.Sleep(2 * time.Second)

	server, err := osquery.NewExtensionManagerServer(
		"puppet_state",
		*flSocketPath,
		osquery.ServerTimeout(time.Duration(*flTimeout)*time.Second),
	)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	// Create and register a new table plugin with the server.
	// Adding a new table? Add it to the list and the loop below will handle
	// the registration for you.
	plugins := []osquery.OsqueryPlugin{
		table.NewPlugin("puppet_info", PuppetInfoColumns(), PuppetInfoGenerate),
		table.NewPlugin("puppet_logs", PuppetLogsColumns(), PuppetLogsGenerate),
		table.NewPlugin("puppet_state", PuppetStateColumns(), PuppetStateGenerate),
	}

	for _, p := range plugins {
		server.RegisterPlugin(p)
	}

	// Start the server. It will run forever unless an error bubbles up.
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
