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

    client := NewClient(tokenID)

    // Component for testing
    c := Component {Name: "Test Component (for Groups testing) with API",
            Description: "Terraform created",
            Showcase: true,
            Status: "operational",
            OnlyShowIfDegraded: false}
    comp, _ := CreateComponent(client, pageID, &c)

    cg := ComponentGroup {Name: "Test Group with API", Description: "Terraform created", Components: []string {comp.ID}}

    compGroup, _ := CreateComponentGroup(client, pageID, &cg)
    assert.Equal(t, pageID, compGroup.PageID)
    assert.True(t, compGroup.Position > 0)

    ucg := ComponentGroup {Name: "Updated Test Group with API", Description: "Updated Terraform created", Components: []string {comp.ID}}
    updatedCompGroup, _ := UpdateComponentGroup(client, pageID, compGroup.ID, &ucg)
    assert.Equal(t, pageID, updatedCompGroup.PageID)
    assert.True(t, updatedCompGroup.Position > 0)
    assert.True(t, updatedCompGroup.Position == compGroup.Position)

    readCompGroup, _ := GetComponentGroup(client, pageID, updatedCompGroup.ID)
    assert.Equal(t, pageID, readCompGroup.PageID)

    cgs, _ := ListAllComponentGroups(client, pageID)
    assert.Equal(t, (*cgs)[0].PageID, pageID)

    // Cleanup
    err := DeleteComponentGroup(client, pageID, readCompGroup.ID)
    assert.Equal(t, nil, err)

    cerr := DeleteComponent(client, pageID, comp.ID)
    assert.Equal(t, nil, cerr)
}

