package gochimp3

import (
	"encoding/json"
	"time"
)

func (order *Order) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Order
		ProcessedAtForeign string `json:"processed_at_foreign"`
		CancelledAtForeign string `json:"cancelled_at_foreign"`
		UpdatedAtForeign   string `json:"updated_at_foreign"`
	}{
		Order:              *order,
		ProcessedAtForeign: order.ProcessedAtForeign.Format(time.RFC3339),
		CancelledAtForeign: order.CancelledAtForeign.Format(time.RFC3339),
		UpdatedAtForeign:   order.UpdatedAtForeign.Format(time.RFC3339),
	}

	return json.Marshal(tmp)
}

func (product *Product) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Product
		PublishedAtForeign string `json:"published_at_foreign"`
	}{
		Product:            *product,
		PublishedAtForeign: product.PublishedAtForeign.Format(time.RFC3339),
	}

	return json.Marshal(tmp)
}
