package main

import (
	"fmt"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
	"github.com/rsmaxwell/archiva/internal/config"
)

func createWantedRepositories(archivaClient *archivaClient.ArchivaClient, cfg *config.Config, session *archivaClient.Session) error {

	for _, repository := range cfg.Repositories {

		repository2, err := archivaClient.GetManagedRepository(session, repository.Id)
		if err != nil {
			return err
		}
		if repository2 == nil {
			fmt.Printf("repository {Id: \"%s\"} not found --> Creating new Repository...", repository.Id)

			_, err := archivaClient.AddManagedRepository(session, repository)
			if err != nil {
				return err
			}

			fmt.Println("done")
			continue
		}

		equals := repository.Compare(repository2)
		if equals {
			fmt.Printf("repository {Id: \"%s\"} is up-to-date\n", repository.Id)
		} else {
			fmt.Printf("repository {Id: \"%s\"} needs updating --> Updating Repository...", repository.Id)

			ok, err := archivaClient.UpdateManagedRepository(session, repository)
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
