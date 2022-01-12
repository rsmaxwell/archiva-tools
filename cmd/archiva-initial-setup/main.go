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

	fmt.Println(os.Getwd())

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

	client, err := archivaClient.NewClient(cfg.Scheme, cfg.Host, cfg.Port, cfg.Base)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client.AcceptInsecureCerts(args.AcceptInsecureCerts)

	admin := cfg.User("admin")
	if admin == nil {
		fmt.Println("admin user not found")
		os.Exit(1)
	}

	session, err := client.Login(admin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = perform(client, cfg, session)

	_, err2 := client.Logout(session)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Success")
}

func perform(client *archivaClient.ArchivaClient, cfg *config.Config, session *archivaClient.Session) error {

	err := removeUnwantedRepositories(client, cfg, session)
	if err != nil {
		return err
	}

	err = createWantedRepositories(client, cfg, session)
	if err != nil {
		return err
	}

	err = removeUnwantedUsers(client, cfg, session)
	if err != nil {
		return err
	}

	err = createWantedUsers(client, cfg, session)
	if err != nil {
		return err
	}

	for _, user := range cfg.Users {

		fmt.Println("")
		fmt.Printf("---[ user.username: %s ]------------------------------------------------------------------\n", user.Username)

		err = removeUnwantedRoles(client, cfg, session, user)
		if err != nil {
			return err
		}

		err = createWantedRoles(client, cfg, session, user)
		if err != nil {
			return err
		}
	}

	return nil
}
