package archivaClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (c *ArchivaClient) UpdateManagedRepository(session *Session, repository *ManagedRepository) (bool, error) {

	repositoryX := repository.ToX()
	requestBody, _ := json.MarshalIndent(repositoryX, "", "    ")
	request := strings.NewReader(string(requestBody))

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + updateManagedRepositoryEndpoint

	req, err := http.NewRequest("POST", url, request)
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Origin", baseUrl)
	req.Header.Add("X-XSRF-TOKEN", session.token)

	httpClient := c.NewHttpClient(session)
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, fmt.Errorf("%s", resp.Status)
	}

	ok, err := strconv.ParseBool(string(responseBody))
	if err != nil {
		return false, err
	}

	return ok, nil
}
