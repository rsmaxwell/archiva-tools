package main

import (
	"fmt"
	"os"
	"path"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
	"github.com/rsmaxwell/archiva/internal/basic"
	"github.com/rsmaxwell/archiva/internal/cmdline"
	"github.com/rsmaxwell/archiva/internal/config"
)

func main() {

	args, err := cmdline.GetArguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if args.Version {
		basic.PrintVersionInfo()
		os.Exit(0)
	}

	configfile := path.Join(args.Configdir, config.DefaultConfigFile)
	cfg, err := config.Open(configfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	admin := cfg.User("admin")
	if admin == nil {
		fmt.Println("admin user not found")
		os.Exit(1)
	}

	client, err := archivaClient.NewClient(cfg.Scheme, cfg.Host, cfg.Port, cfg.Base)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ok, err := client.CreateAdminUser(admin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !ok {
		fmt.Printf("Did not create 'admin' user\n")
		os.Exit(1)
	}

	fmt.Println("Success")
}
