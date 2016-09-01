package gochimp3

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestProductSerialization(t *testing.T) {
	product := new(Product)
	product.PublishedAtForeign = time.Now()

	buf, _ := json.Marshal(product)
	fmt.Println("%s", string(buf))
}
