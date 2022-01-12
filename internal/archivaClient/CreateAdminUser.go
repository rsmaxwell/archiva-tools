package archivaClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (c *ArchivaClient) CreateAdminUser(admin *User) (bool, error) {

	m := make(map[string]string)
	m["username"] = admin.Username
	m["fullName"] = admin.FullName
	m["email"] = admin.Email
	m["password"] = admin.Password

	requestBody, _ := json.MarshalIndent(m, "", "    ")
	request := strings.NewReader(string(requestBody))

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + createAdminUserEndpoint

	req, err := http.NewRequest("POST", url, request)
	if err != nil {
		return false, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Origin", baseUrl)

	httpClient := c.NewHttpClientBasic()
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
