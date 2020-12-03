package gochimp3

const (
	search_members_path = "/search-members"
)

type SearchMembersQueryParams struct {
	BasicQueryParams

	Query  string
	listID string
}

func (q *SearchMembersQueryParams) Params() map[string]string {
	m := q.BasicQueryParams.Params()
	m["query"] = q.Query
	m["list_id"] = q.listID
	return m
}

type SearchMembersResponse struct {
	ExactMatches Matches `json:"exact_matches"`
	FullSearch   Matches `json:"full_search"`
	Links        []Link  `json:"_links"`
}

type Matches struct {
	Members    []Member `json:"members"`
	TotalItems int64    `json:"total_items"`
}

func (list *ListResponse) SearchMembers(params *SearchMembersQueryParams) (*SearchMembersResponse, error) {
	response := new(SearchMembersResponse)

	params.listID = list.ID

	err := list.api.Request("GET", search_members_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
