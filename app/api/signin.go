package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/demostanis/hypertube/pages"
	ghttp "maragu.dev/gomponents/http"
)

type SigninParams struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	PasswordCheck string `json:"passwordCheck"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
}

func jsonInterface(params SigninParams) map[string]interface{} {
	return map[string]interface{}{
		"username":  params.Username,
		"email":     params.Email,
		"enabled":   true,
		"firstName": params.FirstName,
		"lastName":  params.LastName,
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     params.Password,
				"temporary": false,
			},
		},
	}
}

func createUser(params SigninParams, token string) error {
	jsonData, err := json.Marshal(jsonInterface(params))
	if err != nil {
		return fmt.Errorf("Error serializing json: %s", err.Error())
	}

	req, err := http.NewRequest("POST",
		"http://keycloak:8080/admin/realms/default/users",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Error creating request: %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request: %s\n", err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %s", err.Error())
	}

	if string(body) == "" {
		return nil
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("Error parsing response %s", err.Error())
	}

	return fmt.Errorf("Error registering: %s", data["errorMessage"])
}

func APISigninHandler(w http.ResponseWriter, r *http.Request) {
	var params SigninParams

	paramsInto(&params, w, r)

	token, err := adminAuthorization()
	if err != nil {
		fmt.Printf("Failed to get authorization: %s\n", err.Error())
		return
	}

	if params.Password != params.PasswordCheck {
		fmt.Printf("Error: Password doesn't match\n")
		ghttp.Adapt(pages.SigninHandler)(w, r)
		return
	}

	err = createUser(params, token)
	if err != nil {
		fmt.Printf("Failed to create user: %s\n", err.Error())
		ghttp.Adapt(pages.SigninHandler)(w, r)
		return
	}

	fmt.Printf("User %s registered\n", params.Username)
	ghttp.Adapt(pages.SigninHandler)(w, r)
}
