package gochimp3

const (
	campaign_folders_path       = "/campaign-folders"
	// single folder endpoint not implemented
)

type CampaignFolderQueryParams struct {
	ExtendedQueryParams
}

type ListOfCampaignFolders struct {
	baseList
	Folders []CampaignFolder `json:"folders"`
}

type CampaignFolder struct {
	withLinks

	Name  string `json:"name"`
	ID    string `json:"id"`
	Count uint `json:"count"`

	api *API
}

func (api API) GetCampaignFolders(params *CampaignFolderQueryParams) (*ListOfCampaignFolders, error) {
	response := new(ListOfCampaignFolders)

	err := api.Request("GET", campaign_folders_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.Folders {
		l.api = &api
	}

	return response, nil
}
