package archivaClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *ArchivaClient) AllRoles(session *Session) ([]*Role, error) {

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + allRolesEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
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

	var roles []*Role
	err = json.Unmarshal(responseBody, &roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
