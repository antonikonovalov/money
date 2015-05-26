//todo: add test with json
package money

import (
	"testing"
	"encoding/json"
	"strings"
	"fmt"
)

func Test_Encode_Decode(t *testing.T) {
	testPrice := &struct{
		Price Money `json:"price"`
	}{}

	tests := []struct {
		in []byte
		out []byte
		storedPrice int64
	}{
		{[]byte(`{"price":32.01}`), []byte(`{"price":32.01}`), 3201},
		{[]byte(`{"price":124.01645345}`), []byte(`{"price":124.02}`), 12402},
		{[]byte(`{"price":15.22314}`), []byte(`{"price":15.22}`), 1522},
		{[]byte(`{"price":99.999}`),[]byte(`{"price":100.00}`), 10000},
		{[]byte(`{"price":99.99}`),[]byte(`{"price":99.99}`), 9999},
		{[]byte(`{"price":255.654}`),[]byte(`{"price":255.65}`), 25565},
	}

	for _, test := range tests {

		err1 := json.Unmarshal(test.in, testPrice)
		if err1 != nil {
			t.Errorf(fmt.Sprint(err1))
		}
		if int64(testPrice.Price) != test.storedPrice {
			t.Errorf("Extracted incorrect price")
		}


		jsonBytes, err2 := json.Marshal(testPrice)
		if err2 != nil {
			t.Errorf(fmt.Sprint(err2))
		}
		if !strings.EqualFold(string(jsonBytes), string(test.out)) {
			t.Errorf("JSON output does not match")
		}

	}

}
