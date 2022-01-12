package archivaClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

func NewSession(response map[string]interface{}, jar *cookiejar.Jar) (*Session, error) {

	token, ok := response["validationToken"].(string)
	if !ok {
		return nil, fmt.Errorf("validationToken not found")
	}

	return &Session{token: token, jar: jar}, nil
}

func (c *ArchivaClient) Login(u *User) (*Session, error) {
	response, jar, err := c.LoginX(u)
	if err != nil {
		return nil, err
	}

	return NewSession(response, jar)
}

func (c *ArchivaClient) LoginX(u *User) (map[string]interface{}, *cookiejar.Jar, error) {

	m := make(map[string]string)
	m["username"] = u.Username
	m["password"] = u.Password

	requestBody, _ := json.MarshalIndent(m, "", "    ")
	request := strings.NewReader(string(requestBody))

	baseUrl := c.baseUrl()
	url := baseUrl + "/" + loginEndpoint

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, nil, err
	}
	httpClient := c.NewHttpClientBasic()
	httpClient.Jar = jar

	req, err := http.NewRequest("POST", url, request)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Origin", baseUrl)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Response Body:\n%s\n", responseBody)
		return nil, nil, fmt.Errorf("could not logon: status: %s", resp.Status)
	}

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(responseBody, &responseMap)
	if err != nil {
		return nil, nil, err
	}

	return responseMap, jar, nil
}
