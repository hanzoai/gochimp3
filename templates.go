package gochimp3

import (
	"errors"
	"fmt"
)

const (
	templates_path       = "/templates"
	single_template_path = templates_path + "/%s"
	template_default_path = single_template_path + "/default-content"
)


type TemplateQueryParams struct {
	ExtendedQueryParams

	CreatedBy string
	SinceCreatedAt string
	BeforeCreatedAt string
	Type string
	FolderId string
}

func (q *TemplateQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()
	m["created_by"] = q.CreatedBy
	m["since_created_at"] = q.SinceCreatedAt
	m["before_created_at"] = q.BeforeCreatedAt
	m["type"] = q.Type
	m["folder_id"] = q.FolderId

	return m
}

type ListOfTemplates struct {
	baseList
	Templates []TemplateResponse `json:"templates"`
}

type TemplateResponse struct {
	withLinks

	ID          uint `json:"id"`
	Type        string `json:"type"`
	Name        string   `json:"name"`
	DragAndDrop bool `json:"drag_and_drop"`
	Responsive  bool `json:"responsive"`
	Category    string `json:"category"`
	DateCreated string `json:"date_created"`
	CreatedBy   string `json:"created_by"`
	Active      bool `json:"activer"`
	FolderId    string `json:"folder_id"`
	Thumbnail   string `json:"thumbnail"`
	ShareUrl    string `json:"share_url"`

	api *API
}

type TemplateCreationRequest struct {
	Name     string `json:"name"`
	Html     string `json:"html"`
	FolderId string `json:"folder_id"`

}

type TemplateDefaultContentResponse struct {
	withLinks

	Sections map[string]string `json:"sections"`

	api *API
}

func (template *TemplateResponse) CanMakeRequest() error {
	if template.ID == 0 {
		return errors.New("No ID provided on template")
	}

	return nil
}

func (api *API) GetTemplates(params *TemplateQueryParams) (*ListOfTemplates, error) {
	response := new(ListOfTemplates)

	err := api.Request("GET", templates_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.Templates {
		l.api = api
	}

	return response, nil
}

func (api *API) GetTemplate(id string, params *BasicQueryParams) (*TemplateResponse, error) {
	endpoint := fmt.Sprintf(single_template_path, id)

	response := new(TemplateResponse)
	response.api = api

	return response, api.Request("GET", endpoint, params, nil, response)
}

func (api *API) CreateTemplate(body *TemplateCreationRequest) (*TemplateResponse, error) {
	response := new(TemplateResponse)
	response.api = api
	return response, api.Request("POST", templates_path, nil, body, response)
}

func (api *API) UpdateTemplate(id string, body *TemplateCreationRequest) (*TemplateResponse, error) {
	endpoint := fmt.Sprintf(single_template_path, id)

	response := new(TemplateResponse)
	response.api = api

	return response, api.Request("PATCH", endpoint, nil, body, response)
}

func (api *API) DeleteTemplate(id string) (bool, error) {
	endpoint := fmt.Sprintf(single_template_path, id)
	return api.RequestOk("DELETE", endpoint)
}

func (api *API) GetTemplateDefaultContent(id string, params *BasicQueryParams) (*TemplateDefaultContentResponse, error) {
	endpoint := fmt.Sprintf(template_default_path, id)
	response := new(TemplateDefaultContentResponse)
	response.api = api
	return response, api.Request("GET", endpoint, params, nil, response)
}
