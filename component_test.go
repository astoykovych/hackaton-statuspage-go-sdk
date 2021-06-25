package statuspage

import (
	"testing"
	"strings"
    "github.com/stretchr/testify/assert"
    "fmt"
)

func TestCreateComponent (t *testing.T) {

    client := NewClient(tokenID)

    c := Component {Name: "Test Component with API",
        Description: "Terraform created",
        Showcase: true,
        Status: "operational",
        OnlyShowIfDegraded: false}

    comp, _ := CreateComponent(client, pageID, &c)
    fmt.Println(comp)
    assert.Equal(t, pageID, comp.PageID)
    assert.True(t, comp.Showcase)

    // Create group for testing
    cg := ComponentGroup {Name: "Test Component Group (for components) with API", Description: "Terraform created", Components: []string {comp.ID}}
    compGroup, _ := CreateComponentGroup(client, pageID, &cg)

    uc := Component {Name: "Test Updated Component with API",
        Description: "Updated Terraform created",
        GroupID: compGroup.ID,
        Showcase: false,
        Status: "operational",
        OnlyShowIfDegraded: false}
    updatedComp, _ := UpdateComponent(client, pageID, comp.ID, &uc)
    assert.Equal(t, compGroup.ID, updatedComp.GroupID)
    assert.True(t, strings.Contains(updatedComp.Description, "Updated"))

    getComp, _ := GetComponent(client, pageID, updatedComp.ID)
    assert.Equal(t, pageID, getComp.PageID)

    cs, _ := ListAllComponents(client, pageID, "1", "1")
    assert.Equal(t, (*cs)[0].PageID, pageID)

//     Cleanup
    gerr := DeleteComponentGroup(client, pageID, compGroup.ID)
    assert.Equal(t, nil, gerr)

    err := DeleteComponent(client, pageID, getComp.ID)
    assert.Equal(t, nil, err)
}

