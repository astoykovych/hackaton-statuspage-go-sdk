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
    "os"
)

const apiRoot = "https://api.statuspage.io/v1"

var tokenID string

func init() {
	tokenID = os.Getenv("STATUSPAGE_TOKEN")
}


func createResource(pageID string, resourceType string, body interface{}, result interface{}) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType
    fmt.Println("=CREATE====URL=====>")
    fmt.Println(actualURL)
    client := resty.New()
    resp, _ := client.R().SetAuthToken(tokenID).
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Post(actualURL)

	if resp.StatusCode() == http.StatusCreated {
        bodyBytes, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
        if err != nil {
            return err
        }
        fmt.Println(resp)
        return json.Unmarshal(bodyBytes, &result)
	}

	return fmt.Errorf("failed creating resource, request returned %d, full response: %+v", resp.StatusCode(), resp)
}

func getResource(pageID string, resourceType string, ID string, result interface{}) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType + "/" + ID
    fmt.Println("=GET====URL=====>")
    fmt.Println(actualURL)
    client := resty.New()
    resp, _ := client.R().SetAuthToken(tokenID).
        Get(actualURL)

	switch resp.StatusCode() {
	case http.StatusOK:
		bodyBytes, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
		if err != nil {
			return err
		}
        fmt.Println(resp)
		return json.Unmarshal(bodyBytes, &result)

	case http.StatusNotFound:
		return nil

	default:
		return fmt.Errorf("could not find %s with ID: %s, http status %d", resourceType, ID, resp.StatusCode())
	}
}

func updateResource(pageID string, resourceType string, ID string, body interface{}, result interface{}) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType + "/" + ID
    fmt.Println("=UPDATE====URL=====>")
    fmt.Println(actualURL)
    client := resty.New()
    resp, _ := client.R().SetAuthToken(tokenID).
        SetHeader("Content-Type", "application/json").
        SetBody(body).
        Put(actualURL)

	if resp.StatusCode() == http.StatusOK {
        bodyBytes, err := ioutil.ReadAll(bytes.NewReader(resp.Body()))
        if err != nil {
            return err
        }
        fmt.Println(resp)
        return json.Unmarshal(bodyBytes, &result)
	}


	return fmt.Errorf("failed updating %s, request returned %d", resourceType, resp.StatusCode())
}


func deleteResource(pageID string, resourceType string, ID string) error {
    actualURL := apiRoot + "/pages/" + pageID + "/" + resourceType + "/" + ID
    fmt.Println("=DELETE====URL=====>")
    fmt.Println(actualURL)
    client := resty.New()
    resp, err := client.R().SetAuthToken(tokenID).
        Delete(actualURL)

    if err != nil {
        return err
    }

    if resp.StatusCode() == http.StatusNoContent || resp.StatusCode() == http.StatusOK {
        return nil
    }

    return fmt.Errorf("failed deleting %s, request returned %d", resourceType, resp.StatusCode())
}
