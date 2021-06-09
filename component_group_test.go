package statuspage

import (
	"testing"
	"os"
    "github.com/stretchr/testify/assert"
//     "fmt"
)

var pageID string
var tokenID string

func init() {
	pageID = os.Getenv("STATUSPAGE_PAGE")
	tokenID = os.Getenv("STATUSPAGE_TOKEN")
}


func TestEnvVariables (t *testing.T) {

    assert.True(t, tokenID != "")
    assert.True(t, pageID != "")

}


func TestCreateComponentGroup (t *testing.T) {

    testComponent := "r374kwg6gc1s"
    client := NewClient(tokenID)

    cg := ComponentGroup {Name: "Test Group with API", Description: "Terraform created", Components: []string {testComponent}}

    compGroup, _ := CreateComponentGroup(client, pageID, &cg)
    assert.Equal(t, pageID, compGroup.PageID)
    assert.True(t, compGroup.Position > 0)

    ucg := ComponentGroup {Name: "Updated Test Group with API", Description: "Updated Terraform created", Components: []string {testComponent}}
    updatedCompGroup, _ := UpdateComponentGroup(client, pageID, compGroup.ID, &ucg)
    assert.Equal(t, pageID, updatedCompGroup.PageID)
    assert.True(t, updatedCompGroup.Position > 0)
    assert.True(t, updatedCompGroup.Position == compGroup.Position)

    readCompGroup, _ := GetComponentGroup(client, pageID, updatedCompGroup.ID)
    assert.Equal(t, pageID, readCompGroup.PageID)

    err := DeleteComponentGroup(client, pageID, readCompGroup.ID)
    assert.Equal(t, nil, err)
}

