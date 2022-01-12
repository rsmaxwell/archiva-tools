package main

import (
	"encoding/json"
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

	session, err := client.Login(admin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, item := range cfg.Cleanup {
		fmt.Println(item)
	}

	result, err := client.GetManagedRepositories(session)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, repository := range result {

		text, err := json.MarshalIndent(repository, "", "    ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", text)
	}

	_, err = client.Logout(session)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Success")
}
