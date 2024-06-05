/***************************
 * iForm | Main Entry file
 *  -> Loads env
 *  -> Connect to database
 *  -> Start Server
****************************/

package main;

import (
	"iform/cmd"
	"iform/pkg/helpers/env"
)

func main() {
	env.LoadEnvFile(".env", true)
	cmd.Execute()
}