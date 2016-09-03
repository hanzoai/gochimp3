package gochimp3

import (
	"encoding/json"
	"strings"
)

const (
	timeFormat = "2006-01-02T15:04:05-07:00"
)

func (address *Address) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Address
		CountryCode string `json:"country_code"`
	}{
		Address:     *address,
		CountryCode: strings.ToUpper(address.CountryCode),
	}
	return json.Marshal(tmp)
}

func (loc *MemberLocation) MarshalJSON() ([]byte, error) {
	tmp := struct {
		MemberLocation
		CountryCode string `json:"country_code"`
	}{
		MemberLocation: *loc,
		CountryCode:    strings.ToUpper(loc.CountryCode),
	}
	return json.Marshal(tmp)
}

func (store *Store) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Store
		CurrencyCode string `json:"currency_code"`
	}{
		Store:        *store,
		CurrencyCode: strings.ToUpper(store.CurrencyCode),
	}
	return json.Marshal(tmp)
}

func (cart *Cart) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Cart
		CurrencyCode string `json:"currency_code"`
	}{
		Cart:         *cart,
		CurrencyCode: strings.ToUpper(cart.CurrencyCode),
	}
	return json.Marshal(tmp)
}

func (order *Order) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Order
		CurrencyCode       string `json:"currency_code"`
		ProcessedAtForeign string `json:"processed_at_foreign"`
		CancelledAtForeign string `json:"cancelled_at_foreign"`
		UpdatedAtForeign   string `json:"updated_at_foreign"`
	}{
		Order:              *order,
		CurrencyCode:       strings.ToUpper(order.CurrencyCode),
		ProcessedAtForeign: order.ProcessedAtForeign.Format(timeFormat),
		CancelledAtForeign: order.CancelledAtForeign.Format(timeFormat),
		UpdatedAtForeign:   order.UpdatedAtForeign.Format(timeFormat),
	}
	return json.Marshal(tmp)
}

func (product *Product) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Product
		PublishedAtForeign string `json:"published_at_foreign"`
	}{
		Product:            *product,
		PublishedAtForeign: product.PublishedAtForeign.Format(timeFormat),
	}
	return json.Marshal(tmp)
}
