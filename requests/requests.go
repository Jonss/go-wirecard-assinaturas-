package requests

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/Jonss/go-wirecard-assinaturas/config"
)

type httpMethod string

const (
	//GET method
	GET httpMethod = "GET"
	//POST method
	POST httpMethod = "POST"
	PUT  httpMethod = "PUT"
)

// Do executes a request
func Do(method httpMethod, endpoint string, body []byte) (*http.Response, error) {
	wirecardConfig := config.WirecardConfig
	token := wirecardConfig.Token
	key := wirecardConfig.Key

	url := string(wirecardConfig.Env) + endpoint

	req, err := http.NewRequest(string(method), url, bytes.NewBuffer(body))
	fmt.Println(req)
	fmt.Println(req.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.SetBasicAuth(token, key)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
