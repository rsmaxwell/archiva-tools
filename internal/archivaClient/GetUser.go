package archivaClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *ArchivaClient) GetUser(session *Session, username string) (*User, error) {

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + getUserEndpoint + "/" + url.PathEscape(username)

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

	switch resp.StatusCode {
	case 200:
	case 204:
		return nil, nil
	default:
		return nil, fmt.Errorf("%s", resp.Status)
	}

	var user User
	err = json.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
