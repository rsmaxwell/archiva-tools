package archivaClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *ArchivaClient) AddManagedRepository(session *Session, repository *ManagedRepository) (*ManagedRepository, error) {

	repositoryX := repository.ToX()
	requestBody, _ := json.MarshalIndent(repositoryX, "", "    ")
	request := strings.NewReader(string(requestBody))

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + addManagedRepositoryEndpoint

	req, err := http.NewRequest("POST", url, request)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Origin", baseUrl)
	req.Header.Add("X-XSRF-TOKEN", session.token)

	httpClient := c.NewHttpClient(session)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	var repository2 ManagedRepository
	err = json.Unmarshal(responseBody, &repository2)
	if err != nil {
		return nil, err
	}

	return &repository2, nil
}
