package main

import (
	"fmt"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
	"github.com/rsmaxwell/archiva/internal/config"
)

func removeUnwantedUsers(client *archivaClient.ArchivaClient, cfg *config.Config, session *archivaClient.Session) error {

	users, err := client.GetUsers(session)
	if err != nil {
		return err
	}

	for _, user := range users {
		found := false
		for _, user2 := range cfg.Users {
			if user.Username == user2.Username {
				found = true
				break
			}
		}

		if found {
			fmt.Printf("user {Username: \"%s\"} is required\n", user.Username)
		} else {
			fmt.Printf("user {Username: \"%s\"} NOT required --> Deleting User...", user.Username)

			ok, err := client.DeleteUser(session, user.Username)
			if err != nil {
				fmt.Println("")
				return err
			}

			if ok {
				fmt.Println("done")
			} else {
				fmt.Println("NOT deleted")
			}
		}
	}

	return nil
}
