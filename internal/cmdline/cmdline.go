package cmdline

import (
	"flag"
)

// Config type
type CommandlineArguments struct {
	Configdir           string
	Version             bool
	AcceptInsecureCerts bool
}

func GetArguments() (CommandlineArguments, error) {

	configdir := flag.String("config", ".", "configuration directory")

	version := flag.Bool("version", false, "display the version")

	acceptInsecureCerts := flag.Bool("insecure", false, "accept Insecure Certs")

	flag.Parse()

	var args CommandlineArguments
	args.Configdir = *configdir
	args.Version = *version
	args.AcceptInsecureCerts = *acceptInsecureCerts

	return args, nil
}
