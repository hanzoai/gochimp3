package gochimp3

import "fmt"

const (
	segments_path       = "/lists/%s/segments"
	single_segment_path = segments_path + "/%s"
)

type ListOfSegments struct {
	baseList

	Segments []Segment `json:"segments"`
	ListID   string    `json:"list_id"`
}

type SegmentRequest struct {
	Name          string          `json:"name"`
	StaticSegment []string        `json:"static_segment"`
	Options       *SegmentOptions `json:"options,omitempty"`
}

type Segment struct {
	SegmentRequest

	ID          int    `json:"id"`
	MemberCount int    `json:"member_count"`
	Type        string `json:"type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	ListID      string `json:"list_id"`

	withLinks
}

type SegmentOptions struct {
	Match      string               `json:"match"`
	Conditions []SegmentConditional `json:"conditions"`
}

// SegmentBatchRequest represents arguments for bach modifying a static
// segment. Note that both options must be provided, even if empty slice, and
// cannot be nil, or mailchimp will return a 400.
type SegmentBatchRequest struct {
	MembersToAdd    []string `json:"members_to_add"`
	MembersToRemove []string `json:"members_to_remove"`
}

// SegmentBatchResponse is the object returned by MailChimp from a request to
// batch modify a static segment
type SegmentBatchResponse struct {
	MembersAdded   []Member `json:"members_added"`
	MembersRemoved []Member `json:"members_removed"`

	Errors []SegmentBatchError `json:"errors"`

	TotalAdded   int `json:"total_added"`
	TotalRemoved int `json:"total_removed"`
	ErrorCount   int `json:"error_count"`

	withLinks
}

// SegmentBatchError contains errors returned from batch modifying a static
// segment
type SegmentBatchError struct {
	EmailAddresses []string `json:"email_addresses"`
	Error          string   `json:"error"`
}

// SegmentConditional represents parameters to filter by
type SegmentConditional struct {
	Field string      `json:"field"`
	OP    string      `json:"op"`
	Value interface{} `json:"value"`
}

type SegmentQueryParams struct {
	ExtendedQueryParams

	Type            string
	SinceCreatedAt  string
	BeforeCreatedAt string
	SinceUpdatedAt  string
	BeforeUpdatedAt string
}

func (q SegmentQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()

	m["type"] = q.Type
	m["since_created_at"] = q.SinceCreatedAt
	m["since_updated_at"] = q.SinceUpdatedAt
	m["before_created_at"] = q.BeforeCreatedAt
	m["before_updated_at"] = q.BeforeUpdatedAt

	return m
}

func (list ListResponse) GetSegments(params *SegmentQueryParams) (*ListOfSegments, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(segments_path, list.ID)
	response := new(ListOfSegments)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) GetSegment(id string, params *BasicQueryParams) (*Segment, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_segment_path, list.ID, id)
	response := new(Segment)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) CreateSegment(body *SegmentRequest) (*Segment, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(segments_path, list.ID)
	response := new(Segment)

	return response, list.api.Request("POST", endpoint, nil, &body, response)
}

func (list ListResponse) UpdateSegment(id string, body *SegmentRequest) (*Segment, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_segment_path, list.ID, id)
	response := new(Segment)

	return response, list.api.Request("PATCH", endpoint, nil, &body, response)
}

// BatchModifySegment adds and/or removes one or more emails from a static
// segment using POST /lists/{list_id}/segments/{segment_id}. NOTE: You MUST
// check SegmentBatchResponse for errors, as there may be multiple errors (i.e.
// multiple failures to add/remove), and err may still be nil.
func (list ListResponse) BatchModifySegment(id string, body *SegmentBatchRequest) (*SegmentBatchResponse, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_segment_path, list.ID, id)
	response := new(SegmentBatchResponse)

	return response, list.api.Request("POST", endpoint, nil, &body, response)
}

func (list ListResponse) DeleteSegment(id string) (bool, error) {
	if err := list.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(single_segment_path, list.ID, id)
	return list.api.RequestOk("DELETE", endpoint)
}
