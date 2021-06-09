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
	Position  int32  `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var componentGroupsResType = "component-groups"

func CreateComponentGroup(pageID string, componentGroup *ComponentGroup) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull
	err := createResource(
		pageID,
		componentGroupsResType,
		struct {
			ComponentGroup *ComponentGroup `json:"component_group"`
		}{componentGroup},
		&cg,
	)

	return &cg, err
}

func GetComponentGroup(pageID string, componentGroupID string) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull
	err := getResource(
		pageID,
		componentGroupsResType,
		componentGroupID,
		&cg,
	)

	return &cg, err
}

func UpdateComponentGroup(pageID string, componentGroupID string, componentGroup *ComponentGroup) (*ComponentGroupFull, error) {
	var cg ComponentGroupFull

	err := updateResource(
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

func DeleteComponentGroup(pageID string, componentGroupID string) (err error) {
	return deleteResource(pageID, componentGroupsResType, componentGroupID)
}
