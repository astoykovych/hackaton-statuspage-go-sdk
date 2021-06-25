package statuspage

type Component struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	GroupID            string `json:"group_id,omitempty"`
	Showcase           bool   `json:"showcase,omitempty"`
	Status             string `json:"status,omitempty"`
	OnlyShowIfDegraded bool   `json:"only_show_if_degraded,omitempty"`
}

type ComponentFull struct {
	Component
	ID              string `json:"id"`
	PageID          string `json:"page_id"`
	Position        int32  `json:"position"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	AutomationEmail string `json:"automation_email"`
}

var componentResType = "components"

func CreateComponent(client *RClient, pageID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull
	err := createResource(
		client,
		pageID,
		componentResType,
		struct {
			Component *Component `json:"component"`
		}{component},
		&c,
	)

	return &c, err
}

func GetComponent(client *RClient, pageID string, componentID string) (*ComponentFull, error) {
	var c ComponentFull
	err := getResource(client, pageID, componentResType, componentID, &c)

	return &c, err
}

func UpdateComponent(client *RClient, pageID, componentID string, component *Component) (*ComponentFull, error) {
	var c ComponentFull

	err := updateResource(
		client,
		pageID,
		componentResType,
		componentID,
		struct {
			Component *Component `json:"component"`
		}{component},
		&c,
	)

	return &c, err
}

func DeleteComponent(client *RClient, pageID, componentID string) (err error) {
	return deleteResource(client, pageID, componentResType, componentID)
}

func ListAllComponents(client *RClient, pageID string, page string, per_page string) (*[]ComponentFull, error) {
	var cs []ComponentFull
	qp := map[string]string{
           		"page": page,
           		"per_page": per_page,
           	}
	err := listResources(
        client,
		pageID,
		componentResType,
		&qp,
		&cs,
	)

	return &cs, err
}
