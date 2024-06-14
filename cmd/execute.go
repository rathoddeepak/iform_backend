package cmd;

import (
	"flag"
	"fmt"

	"iform/pkg/cacheconf"
	"iform/boot"
)

const (
	iform         = "iform"
	migrateDatabase = "migrate"
	startServer   = "start"
	showConfig    = "config"
)

// getDefaultMessage - Returns Default message
func getDefaultMessage() string {
return `
commands:
-iform start          -> start the Server
-iform migrate        -> migrates database
-iform config         -> show the current environment 
`;
}


// Identifies command and executes accordingly
func Execute() {

	start := flag.String(iform, "", getDefaultMessage())
	flag.Parse();

	switch *start {

		case startServer:
			boot.InitServer()

		case migrateDatabase:
			boot.MigrateDatabase()

		case showConfig:
			fmt.Println(cacheconf.GetCurrentConfig())

		default:
			flag.PrintDefaults()
	}

}