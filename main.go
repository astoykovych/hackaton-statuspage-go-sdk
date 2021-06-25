package statuspage

import (
	"bytes"
	"encoding/json"
	"fmt"
// 	"io"
	"io/ioutil"
    "net/http"
// 	"time"
    "github.com/go-resty/resty/v2"
//     "os"
)

const apiRoot = "https://api.statuspage.io/v1"
const logging_on = true

type RClient struct {
	token           string
	restyClient     *resty.Client
}

func NewClient(token string) *RClient {

    rc := RClient{
          		token:       token,
          		restyClient: resty.New().SetRetryCount(10),
          	}
	return &rc
}

func log(debug_name string, debug_value interface{}) {
    if logging_on {
        fmt.Printf("DEBUG: [%s] => ", debug_name)
        fmt.Println(debug_value)
    }
}


func createResource(client *RClient, pageID string, resourceType string, body interface{}, result interface{}) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType
    log("CREATE", actualURL)

    resp, _ := client.restyClient.R().SetAuthToken(client.token).
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Post(actualURL)

	if resp.StatusCode() == http.StatusCreated {
        bodyBytes, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
        if err != nil {
            return err
        }
        log("RESPONSE", resp)
        return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("failed creating resource, request returned %d, full response: %+v", resp.StatusCode(), resp)
}

func getResource(client *RClient, pageID string, resourceType string, ID string, result interface{}) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType + "/" + ID
    log("GET", actualURL)

    resp, _ := client.restyClient.R().SetAuthToken(client.token).
        Get(actualURL)

	switch resp.StatusCode() {
	case http.StatusOK:
		bodyBytes, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
		if err != nil {
			return err
		}
        log("RESPONSE", resp)
		return json.Unmarshal(bodyBytes, &result)

	case http.StatusNotFound:
		return nil

	default:
		return fmt.Errorf("could not find %s with ID: %s, http status %d", resourceType, ID, resp.StatusCode())
	}
}

func updateResource(client *RClient, pageID string, resourceType string, ID string, body interface{}, result interface{}) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType + "/" + ID
    log("UPDATE", actualURL)

    resp, _ := client.restyClient.R().SetAuthToken(client.token).
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Put(actualURL)

	if resp.StatusCode() == http.StatusOK {
        bodyBytes, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
        if err != nil {
            return err
        }
        log("RESPONSE", resp)
        return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("failed updating %s, request returned %d", resourceType, resp.StatusCode())
}


func deleteResource(client *RClient, pageID string, resourceType string, ID string) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType + "/" + ID
    log("DELETE", actualURL)

    resp, err := client.restyClient.R().SetAuthToken(client.token).
        Delete(actualURL)
    log("RESPONSE", resp)

    if err != nil {
        return err
    }

    if resp.StatusCode() == http.StatusNoContent || resp.StatusCode() == http.StatusOK {
        return nil
    }

    return fmt.Errorf("failed deleting %s, request returned %d", resourceType, resp.StatusCode())
}

func listResources(client *RClient, pageID string, resourceType string, qparams *map[string]string, result interface{}) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType
    log("List All", actualURL)

    resp, _ := client.restyClient.R().SetAuthToken(client.token).
        SetQueryParams(*qparams).
        Get(actualURL)

    if resp.StatusCode() == http.StatusOK {
            bodyBytes, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
            if err != nil {
                return err
            }
            log("RESPONSE", resp)
            return json.Unmarshal(bodyBytes, &result)
    	}

    return fmt.Errorf("failed getting all resources %s, request returned %d", resourceType, resp.StatusCode())
}