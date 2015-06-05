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

func TestString(t *testing.T) {
	tests := []struct {
		money Money
		string string
	}{
		{3200, "32.00"},
		{3201, "32.01"},
		{12401, "124.01"},
		{9999, "99.99"},
		{-567, "-5.67"},
		{0, "0.00"},
	}

	for _, test := range tests {
		converted := test.money.String()
		if converted != test.string {
			t.Errorf(`Error converting money to string: result "%s", expected "%s"`, converted, test.string)
		}
	}
}

func TestToFloat64Conversion(t *testing.T) {
	tests := []struct {
		money Money
		float64 float64
	}{
		{3200, 32.00},
		{3201, 32.01},
		{12401, 124.01},
		{9999, 99.99},
		{-567, -5.67},
		{0, 0.00},
	}

	for _, test := range tests {
		converted := test.money.Float64()
		if converted != test.float64 {
			t.Errorf(`Error converting money to float64: result "%f", expected "%f"`, converted, test.float64)
		}
	}
}

func TestFromFloat64Conversion(t *testing.T) {
	tests := []struct {
		float64 float64
		money Money
	}{
		{9.7788, 978},
	    {9.77, 977},
	    {9.771, 977},
		{9.7750001, 978},
		{9.7749999, 977},
		{0.0, 0},
		{-0.0, 0},
		{-10.40, -1040},
		{-155.62, -15562},
		{-9.771, -977},
		{-9.7749999, -977},
		{-9.7750001, -978},
	}

	for _, test := range tests {
		converted := FromFloat64(test.float64)
		if converted != test.money {
			t.Errorf(`Error converting float64 to money: result "%d", expected "%d"`, converted, test.money)
		}
	}
}
