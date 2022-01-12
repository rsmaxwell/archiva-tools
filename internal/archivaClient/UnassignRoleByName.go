package archivaClient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (c *ArchivaClient) UnassignRoleByName(session *Session, username string, roleName string) (bool, error) {

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + unassignRoleByNameEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Origin", baseUrl)
	req.Header.Add("X-XSRF-TOKEN", session.token)

	q := req.URL.Query()
	q.Add("principal", username)
	q.Add("roleName", roleName)
	req.URL.RawQuery = q.Encode()

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
		fmt.Printf("Response Body:\n%s\n", responseBody)
		return false, fmt.Errorf("%s", resp.Status)
	}

	ok, err := strconv.ParseBool(string(responseBody))
	if err != nil {
		return false, err
	}

	return ok, nil
}
