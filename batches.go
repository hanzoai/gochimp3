package gochimp3

import (
	"fmt"
	"net/url"
)

const (
	batches_path      = "/batches"
	single_batch_path = batches_path + "/%s"
)

func (api *API) GetBatchOperations(params *ListQueryParams) (*ListOfBatchOperations, error) {
	response := new(ListOfBatchOperations)

	err := api.Request("GET", batches_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.BatchOperations {
		l.api = api
	}

	return response, nil
}

type ListOfBatchOperations struct {
	baseList
	BatchOperations []BatchOperationResponse `json:"batches"`
}

func (api *API) GetBatchOperation(id string, params *BasicQueryParams) (*BatchOperationResponse, error) {
	endpoint := fmt.Sprintf(single_batch_path, id)
	response := new(BatchOperationResponse)
	response.api = api

	return response, api.Request("GET", endpoint, params, nil, response)
}

func (api *API) CreateBatchOperation(body *BatchOperationCreationRequest) (*BatchOperationResponse, error) {
	response := new(BatchOperationResponse)
	response.api = api
	return response, api.Request("POST", batches_path, nil, body, response)
}

type BatchOperationCreationRequest struct {
	Operations []BatchOperation `json:"operations"`
}

type BatchOperationResponse struct {
	Links []Link `json:"_links,omitempty"`

	ID                 string `json:"id"`
	Status             string `json:"status"`
	TotalOperations    int    `json:"total_operations"`
	FinishedOperations int    `json:"finished_operations"`
	ErroredOperations  int    `json:"errored_operations"`
	SubmittedAt        string `json:"submitted_at,omitempty"`
	CompletedAt        string `json:"completed_at,omitempty"`
	ResponseBodyUrl    string `json:"response_body_url"`

	api *API
}

type BatchOperation struct {
	Method      string     `json:"method"`
	Path        string     `json:"path"`
	Params      url.Values `json:"params,omitempty"`
	Body        string     `json:"body"`
	OperationID string     `json:"operation_id,omitempty"`

	api *API
}
