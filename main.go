package statuspage

import (
// 	"bytes"
	"encoding/json"
	"fmt"
// 	"io"
	"io/ioutil"
    "net/http"
// 	"time"
    "github.com/go-resty/resty/v2"
    "os"
)

const apiRoot = "https://api.statuspage.io/v1"

var tokenID string

func init() {
	tokenID = os.Getenv("STATUSPAGE_TOKEN")
}

// // HTTPClient is the http wrapper for the application
// type HTTPClient interface {
// 	Do(req *http.Request) (*http.Response, error)
// }
//
// type IClient interface {
// 	doHTTPRequest(method, endpoint string, item interface{}) (resp *http.Response, err error)
// }
//
// type Client struct {
// 	token      string
// 	httpClient HTTPClient
// }
//
// func NewClient(token string) *Client {
// 	return &Client{
// 		token:      token,
// 		httpClient: &http.Client{},
// 	}
// }

// Allows overriding the HTTP Client, leaving the choice
// of using a retry library to the user
// func (client *Client) UseHTTPClient(httpClient HTTPClient) {
// 	client.httpClient = httpClient
// }
//
// func (client *Client) old_doHTTPRequest(method, endpoint string, item interface{}) (resp *http.Response, err error) {
// 	componentURL := apiRoot + endpoint
//
// 	var body io.Reader
//
// 	if item != nil {
// 		data, err := json.Marshal(item)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		body = bytes.NewReader(data)
// 	}
// 	req, err := http.NewRequest(method, componentURL, body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Authorization", "OAuth "+client.token)
//
// 	maxRetries := 10
// 	retryInterval := 10 * time.Second
//
// 	// Basic Retry logic around rate limiting
// 	resp, err = client.httpClient.Do(req)
// 	retries := 0
// 	for retries = 1; resp != nil && resp.StatusCode == 420 && retries <= maxRetries; retries = retries + 1 {
// 		time.Sleep(retryInterval)
// 		resp, err = client.httpClient.Do(req)
// 	}
//
// 	return resp, err
// }
//
// func old_createResource(client IClient, pageID, resourceType string, resource, result interface{}) error {
// 	return createResourceCustomURL(client, "/pages/"+pageID+"/"+resourceType+"s", resource, result)
// }
//
// func old_createResourceCustomURL(client IClient, URL string, resource, result interface{}) error {
// 	resp, err := client.doHTTPRequest(
// 		"POST",
// 		URL,
// 		resource,
// 	)
// 	if err != nil {
// 		return err
// 	}
//
// 	if resp.StatusCode == http.StatusCreated {
// 		defer resp.Body.Close()
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return err
// 		}
// 		return json.Unmarshal(bodyBytes, &result)
// 	}
//
// 	return fmt.Errorf("failed creating resource, request returned %d, full response: %+v", resp.StatusCode, resp)
// }


func createResource(pageID string, resource string, body interface{}, result interface{}) error {
    client := resty.New()
    resp, _ := client.R().SetAuthToken(tokenID).
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Post(fmt.Sprintf("s%/pages/%s/%s", apiRoot, pageID, resource))

	if resp.StatusCode() == http.StatusCreated {
        bodyBytes, err := ioutil.ReadAll(resp.Body())
        if err != nil {
            return err
        }
        return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("failed creating resource, request returned %d, full response: %+v", resp.StatusCode(), resp)
}

//
// func old_readResource(client IClient, pageID, ID, resourceType string, target interface{}) error {
// 	resp, err := client.doHTTPRequest(
// 		"GET",
// 		"/pages/"+pageID+"/"+resourceType+"s/"+ID,
// 		nil,
// 	)
// 	if err != nil {
// 		return err
// 	}
//
// 	if resp.Body != nil {
// 		defer resp.Body.Close()
// 	}
//
// 	switch resp.StatusCode {
// 	case http.StatusOK:
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return err
// 		}
// 		return json.Unmarshal(bodyBytes, target)
//
// 	case http.StatusNotFound:
// 		return nil
//
// 	default:
// 		return fmt.Errorf("could not find %s with ID: %s, http status %d", resourceType, ID, resp.StatusCode)
// 	}
// }
//
// func old_updateResource(client IClient, pageID, resourceType, ID string, resource interface{}, result interface{}) error {
// 	resp, err := client.doHTTPRequest(
// 		"PATCH",
// 		"/pages/"+pageID+"/"+resourceType+"s/"+ID,
// 		resource,
// 	)
// 	if err != nil {
// 		return err
// 	}
//
// 	if resp.StatusCode == http.StatusOK {
// 		defer resp.Body.Close()
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return err
// 		}
//
// 		return json.Unmarshal(bodyBytes, &result)
// 	}
//
// 	return fmt.Errorf("failed updating %s, request returned %d", resourceType, resp.StatusCode)
// }
//
// func old_deleteResource(client IClient, pageID, resourceType, ID string) error {
// 	resp, err := client.doHTTPRequest(
// 		"DELETE",
// 		"/pages/"+pageID+"/"+resourceType+"s/"+ID,
// 		nil,
// 	)
// 	if err != nil {
// 		return err
// 	}
//
// 	// StatusGroup deletion returns StatusOK instead of StatusNoContent like other resources
// 	if resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusOK {
// 		return nil
// 	}
//
// 	return fmt.Errorf("failed deleting %s, request returned %d", resourceType, resp.StatusCode)
// }
