package statuspage

type IncidentTemplate struct {
	Name         string  `json:"name"`
	GroupID      string  `json:"group_id"`
	UpdateStatus string  `json:"update_status"`
	Title        string  `json:"title"`
	Body         string  `json:"body"`
	ComponentIDs []string `json:"component_ids"`
	ShouldTweet  bool    `json:"should_tweet"`
	ShouldSendNotifications  bool    `json:"should_send_notifications"`
}

type IncidentTemplateFull struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
    Name         string  `json:"name"`
    GroupID      string  `json:"group_id"`
    UpdateStatus string  `json:"update_status"`
    Title        string  `json:"title"`
    Body         string  `json:"body"`
    Components []ComponentFull `json:"components"`
    ShouldTweet  bool    `json:"should_tweet"`
    ShouldSendNotifications  bool    `json:"should_send_notifications"`
}

var componentTemplateResType = "incident_templates"

func CreateIncidentTemplate(client *RClient, pageID string, incidentTemplate *IncidentTemplate) (*IncidentTemplateFull, error) {
	var i IncidentTemplateFull
	err := createResource(
		client,
		pageID,
		componentTemplateResType,
		struct {
			IncidentTemplate *IncidentTemplate `json:"template"`
		}{incidentTemplate},
		&i,
	)

	return &i, err
}

// func GetIncidentTemplate(client *RClient, pageID, incidentTemplateID string) (*IncidentTemplateFull, error) {
// 	var i IncidentTemplateFull
// 	err := getResource(client, pageID, componentTemplateResType, incidentTemplateID, &i)
//
// 	return &i, err
// }
//
// func UpdateIncidentTemplate(client *RClient, pageID, incidentTemplateID string, incidentTemplate *IncidentTemplate) (*IncidentTemplateFull, error) {
// 	var i IncidentTemplateFull
//
// 	err := updateResource(
// 		client,
// 		pageID,
// 		componentTemplateResType,
// 		incidentTemplateID,
// 		struct {
// 			IncidentTemplate *IncidentTemplate `json:"template"`
// 		}{incidentTemplate},
// 		&i,
// 	)
//
// 	return &i, err
// }
//
// func DeleteIncidentTemplate(client *RClient, pageID, incidentTemplateID string) (err error) {
// 	return deleteResource(client, pageID, componentTemplateResType, incidentTemplateID)
// }

func ListAllTemplates(client *RClient, pageID string, page string, per_page string) (*[]IncidentTemplateFull, error) {
	var it []IncidentTemplateFull
    qp := map[string]string{
                "page": page,
                "per_page": per_page,
            }
	err := listResources(
        client,
		pageID,
		componentTemplateResType,
        &qp,
		&it,
	)

	return &it, err
}