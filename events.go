package gochimp3

import (
	"fmt"
	"time"
)

const (
	member_events_path = single_member_path + "/events"
)

type EventRequest struct {
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties,omitempty"`
	IsSyncing  bool              `json:"is_syncing,omitempty"`
	OccurredAt *time.Time        `json:"occurred_at,omitempty"`
}

func (m *Member) AddEvent(e *EventRequest) error {
	if err := m.CanMakeRequest(); err != nil {
		return err
	}

	endpoint := fmt.Sprintf(member_events_path, m.ListID, m.ID)

	return m.api.Request("POST", endpoint, nil, e, nil)
}

func (m *Member) AddSimpleEvent(name string) error {
	return m.AddEvent(&EventRequest{Name: name})
}
