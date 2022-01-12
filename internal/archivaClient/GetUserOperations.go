package archivaClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *ArchivaClient) GetUserOperations(session *Session, user *User) ([]*Operation, error) {

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + getUserOperationsEndpoint + "/" + url.PathEscape(user.Username)

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

	var operations []*Operation
	err = json.Unmarshal(responseBody, &operations)
	if err != nil {
		return nil, err
	}

	return operations, nil
}
