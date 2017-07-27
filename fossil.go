// Fossil fossil Graphite to sensision transpiler.
//
// Usage
//
// 		fossil  [flags]
// Flags:
//       --config string   config file to use
//       --help            display help
//   -v, --verbose         verbose output
//   -l, --listen          listen addresse
//   -d  --directory       directory to write metrics files
package main

import (
	log "github.com/Sirupsen/logrus"
	"stash.ovh.net/metrics/fossil/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Panicf("%v", err)
	}
}
