package vismanet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBranchGetAll(t *testing.T) {
	req := client.NewBranchGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
