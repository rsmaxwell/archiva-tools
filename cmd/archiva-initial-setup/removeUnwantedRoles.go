package main

import (
	"fmt"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
	"github.com/rsmaxwell/archiva/internal/config"
)

func removeUnwantedRoles(client *archivaClient.ArchivaClient, cfg *config.Config, session *archivaClient.Session, user *archivaClient.User) error {

	roles, err := client.GetEffectivelyAssignedRoles(session, user)
	if err != nil {
		return err
	}

	rolesAssignedToUser := []*archivaClient.Role{}
	for _, role := range roles {

		role2, err := client.GetRole(session, role.Name)
		if err != nil {
			return err
		}

		for _, u := range role2.Users {
			if u.Username == user.Username {
				rolesAssignedToUser = append(rolesAssignedToUser, role2)
			}
		}
	}

	// Check each role assigned to the given user, is in the configuration
	for _, role := range rolesAssignedToUser {

		found := false
		for _, roleName := range user.AssignedRoles {
			if roleName == role.Name {
				found = true
				break
			}
		}

		if found {
			fmt.Printf("role {Name: \"%s\"} is required\n", role.Name)
		} else {
			fmt.Printf("role {Name: \"%s\"} NOT required --> Unassigning Role...", role.Name)
			ok, err := client.UnassignRoleByName(session, user.Username, role.Name)
			if err != nil {
				fmt.Println("")
				fmt.Println(err)
				return err
			}
			if ok {
				fmt.Println("done")
			} else {
				fmt.Println("NOT Unassigned")
			}
		}
	}

	return nil
}
