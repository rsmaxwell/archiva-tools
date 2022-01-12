package main

import (
	"fmt"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
	"github.com/rsmaxwell/archiva/internal/config"
)

func createWantedUsers(archivaClient *archivaClient.ArchivaClient, cfg *config.Config, session *archivaClient.Session) error {

	for _, user := range cfg.Users {

		user2, err := archivaClient.GetUser(session, user.Username)
		if err != nil {
			return err
		}
		if user2 == nil {
			fmt.Printf("user {Username: \"%s\"} not found --> Creating new User...", user.Username)

			ok, err := archivaClient.CreateUser(session, user)
			if err != nil {
				return err
			}

			if ok {
				fmt.Println("done")
			} else {
				fmt.Println("NOT created")
			}
			continue
		}

		equals := user.Compare(user2)
		if equals {
			fmt.Printf("user {Username: \"%s\"} is up-to-date\n", user.Username)
		} else {
			fmt.Printf("user {Username: \"%s\"} needs updating --> Updating User...", user.Username)

			ok, err := archivaClient.UpdateUser(session, user)
			if err != nil {
				fmt.Println("")
				return err
			}
			if ok {
				fmt.Println("done")
			} else {
				fmt.Println("NOT updated")
			}
		}
	}

	return nil
}
