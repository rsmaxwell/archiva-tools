package archivaClient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *ArchivaClient) Logout(session *Session) (bool, error) {

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + logoutEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
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

	if ok {
		session.jar = nil
		session.token = ""
	}

	return ok, nil
}
