package statuspage

type ComponentGroup struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Components  []string `json:"components,omitempty"`
}

type ComponentGroupFull struct {
	ComponentGroup
	ID        string `json:"id"`
	PageID    string `json:"page_id"`
	Position  string `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var componentGroupsResType = "component-groups"

func CreateComponentGroup(client *RClient, pageID string, componentGroup *ComponentGroup) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull
	err := createResource(
        client,
		pageID,
		componentGroupsResType,
		struct {
			ComponentGroup *ComponentGroup `json:"component_group"`
		}{componentGroup},
		&cg,
	)

	return &cg, err
}

func GetComponentGroup(client *RClient, pageID string, componentGroupID string) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull
	err := getResource(
        client,
		pageID,
		componentGroupsResType,
		componentGroupID,
		&cg,
	)

	return &cg, err
}

func UpdateComponentGroup(client *RClient, pageID string, componentGroupID string, componentGroup *ComponentGroup) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull

	err := updateResource(
        client,
		pageID,
		componentGroupsResType,
		componentGroupID,
		struct {
			ComponentGroup *ComponentGroup `json:"component_group"`
		}{componentGroup},
		&cg,
	)

	return &cg, err
}

func DeleteComponentGroup(client *RClient, pageID string, componentGroupID string) (err error) {
	return deleteResource(client, pageID, componentGroupsResType, componentGroupID)
}

func ListAllComponentGroups(client *RClient, pageID string) (*[]ComponentGroupFull, error) {
	var cgs []ComponentGroupFull
	err := listResources(
        client,
		pageID,
		componentGroupsResType,
		&cgs,
	)

	return &cgs, err
}
