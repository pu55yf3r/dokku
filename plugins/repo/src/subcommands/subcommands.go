package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
	"github.com/dokku/dokku/plugins/repo"

	flag "github.com/spf13/pflag"
)

// main entrypoint to all subcommands
func main() {
	parts := strings.Split(os.Args[0], "/")
	subcommand := parts[len(parts)-1]

	var err error
	switch subcommand {
	case "gc":
		args := flag.NewFlagSet("repo:gc", flag.ExitOnError)
		args.Parse(os.Args[2:])
		appName := args.Arg(0)
		err = repo.CommandGc(appName)
	case "purge-cache":
		args := flag.NewFlagSet("repo:purge-cache", flag.ExitOnError)
		args.Parse(os.Args[2:])
		appName := args.Arg(0)
		err = repo.CommandPurgeCache(appName)
	default:
		common.LogFail(fmt.Sprintf("Invalid plugin subcommand call: %s", subcommand))
	}

	if err != nil {
		common.LogFail(err.Error())
	}
}
