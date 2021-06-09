package statuspage

import (
	"testing"
	"os"
	"assert"
)


var pageID string

func init() {
	pageID = os.Getenv("STATUSPAGE_PAGE")
}

func TestCreateComponentGroup(t *testing.T) {

    cg := struct {
          				Name        string
          				Description string
          			}{
          				"group_name",
          				"group description",
          			}

    compGroup, err := CreateComponentGroup(pageID, cg)

    assert.Equal(t, pageID, compGroup.PageID)
    assert.True(t, compGroup.Position > 0)
}

