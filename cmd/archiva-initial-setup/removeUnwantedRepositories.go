package main

import (
	"fmt"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
	"github.com/rsmaxwell/archiva/internal/config"
)

func removeUnwantedRepositories(client *archivaClient.ArchivaClient, cfg *config.Config, session *archivaClient.Session) error {

	allRepos, err := client.GetManagedRepositories(session)
	if err != nil {
		return err
	}

	// Check each repository is in the configuration
	for _, repository := range allRepos {
		found := false
		for _, repository2 := range cfg.Repositories {
			if repository.Id == repository2.Id {
				found = true
				break
			}
		}

		if found {
			fmt.Printf("repository {Id: \"%s\"} is required\n", repository.Id)
		} else {
			fmt.Printf("repository {Id: \"%s\"} NOT required --> Deleting Repository...", repository.Id)
			deleteContent := true
			ok, err := client.DeleteManagedRepository(session, repository.Id, deleteContent)
			if err != nil {
				fmt.Println("")
				fmt.Println(err)
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
