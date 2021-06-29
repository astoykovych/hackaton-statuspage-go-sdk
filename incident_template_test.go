package statuspage

import (
	"testing"
	"strings"
    "github.com/stretchr/testify/assert"
    "fmt"
)

func TestIncidentTemplate (t *testing.T) {

    client := NewClient(tokenID)

    i := IncidentTemplate {Name: "Name - Test Template with API",
        Title: "Title - Test Template with API",
        UpdateStatus: "investigating",
        Body: "Test - Terraform created",
        }

    inc, _ := CreateIncidentTemplate(client, pageID, &i)
    fmt.Println(inc)
    assert.True(t, strings.Contains(inc.Body, "Terraform"))
    assert.False(t, inc.ShouldSendNotifications)

//     ui := IncidentTemplate {Name: "Updated Name - Test Template with API",
//         Title: "Title - Test Template with API",
//         UpdateStatus: "investigating",
//         Body: "Test - Terraform created",
//         }
//     updatedInc, _ := UpdateIncidentTemplate(client, pageID, inc.ID, &ui)
//     assert.True(t, strings.Contains(updatedInc.Name, "Updated"))

//     getInc, _ := GetIncidentTemplate(client, pageID, inc.ID)
//     assert.Equal(t, inc.ID, getInc.ID)

    ts, _ := ListAllTemplates(client, pageID, "1", "1")
    assert.True(t, strings.Contains((*ts)[0].Components[0].PageID, pageID))
//
//     Cleanup
//     gerr := DeleteIncidentTemplate(client, pageID, inc.ID)
//     assert.Equal(t, nil, gerr)

}

