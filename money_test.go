//todo: add test with json
package money

import (
	"testing"
	"encoding/json"
	"strings"
)

func Test_Encode_Decode(t *testing.T) {
	testPrice := &struct{
		Price Money `json:"price"`
	}{}

	tests := []struct {
		in []byte
		out []byte
		storedPrice int64
		badFloat bool
	}{
		{[]byte(`{"price":32}`), []byte(`{"price":32.00}`), 3200, false},
		{[]byte(`{"price":32.01}`), []byte(`{"price":32.01}`), 3201, false},
		{[]byte(`{"price":124.01645345}`), []byte(`{"price":124.02}`), 12402, false},
		{[]byte(`{"price":15.22314}`), []byte(`{"price":15.22}`), 1522, false},
		{[]byte(`{"price":99.999}`), []byte(`{"price":100.00}`), 10000, false},
		{[]byte(`{"price":99.99}`), []byte(`{"price":99.99}`), 9999, false},
		{[]byte(`{"price":255.654}`), []byte(`{"price":255.65}`), 25565, false},
		{[]byte(`{"price":"$$$"}`), []byte(`{"price":32.01}`), 3201, true},
		{[]byte(`{"price":"32,01"}`), []byte(`{"price":32.01}`), 3201, true},
	}

	for _, test := range tests {

		err := json.Unmarshal(test.in, testPrice)
		if err != nil && !test.badFloat {
			t.Error(err)
		}
		if int64(testPrice.Price) != test.storedPrice && !test.badFloat {
			t.Errorf("Extracted incorrect price")
		}


		jsonBytes, err := json.Marshal(testPrice)
		if err != nil && !test.badFloat {
			t.Error(err)
		}
		if !strings.EqualFold(string(jsonBytes), string(test.out)) && !test.badFloat {
			t.Errorf("JSON output does not match")
		}

	}

}
