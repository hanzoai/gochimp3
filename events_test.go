package gochimp3

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventsJson(t *testing.T) {
	unix := time.Unix(0, 0)
	eventRequest := &EventRequest{
		Name:       "huge_deal",
		Properties: map[string]string{
			"R2": "D2",
		},
		IsSyncing:  false,
		OccurredAt: &unix,
	}

	bytes, err := json.Marshal(eventRequest)
	assert.NoError(t, err)
	backFromJson := &EventRequest{}
	assert.NoError(t, json.Unmarshal(bytes, backFromJson))
	assert.Equal(t, eventRequest, backFromJson)

	missingFields := &EventRequest{
		Name:       "big_deal",
	}

	bytes, err = json.Marshal(missingFields)
	assert.NoError(t, err)
	assert.Equal(t, "{\"name\":\"big_deal\"}", string(bytes))
	backFromJson = &EventRequest{}
	assert.NoError(t, json.Unmarshal(bytes, backFromJson))
	assert.Equal(t, missingFields, backFromJson)
	assert.Nil(t, backFromJson.Properties)
	assert.Nil(t, backFromJson.OccurredAt)
}