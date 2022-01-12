package main

import (
	"fmt"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
	"github.com/rsmaxwell/archiva/internal/config"
)

func createWantedRoles(client *archivaClient.ArchivaClient, cfg *config.Config, session *archivaClient.Session, user *archivaClient.User) error {

	for _, roleName := range user.AssignedRoles {

		role, err := client.GetRole(session, roleName)
		if err != nil {
			return err
		}

		found := false
		for _, u := range role.Users {
			if u.Username == user.Username {
				found = true
				break
			}
		}

		if found {
			fmt.Printf("role {Name: \"%s\"} is already assigned\n", roleName)
		} else {

			fmt.Printf("role {Name: \"%s\"} needs to be assigned --> Assigning Role...", roleName)

			ok, err := client.AssignRoleByName(session, user.Username, roleName)
			if err != nil {
				return err
			}

			if ok {
				fmt.Println("done")
			} else {
				fmt.Println("NOT Assigned")
			}
		}
	}

	return nil
}
