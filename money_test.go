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
		{[]byte(`{"price":32}`), []byte(`{"price":32.0000}`), 320000, false},
		{[]byte(`{"price":32.01}`), []byte(`{"price":32.0100}`), 320100, false},
		{[]byte(`{"price":124.01645345}`), []byte(`{"price":124.0165}`), 1240165, false},
		{[]byte(`{"price":15.22314}`), []byte(`{"price":15.2231}`), 152231, false},
		{[]byte(`{"price":99.999}`), []byte(`{"price":99.9990}`), 999990, false},
		{[]byte(`{"price":99.99}`), []byte(`{"price":99.9900}`), 999900, false},
		{[]byte(`{"price":255.654}`), []byte(`{"price":255.6540}`), 2556540, false},
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
		{3200, "0.3200"},
		{3201, "0.3201"},
		{12401, "1.2401"},
		{9999, "0.9999"},
		{-567, "-0.0567"},
		{0, "0.0000"},
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
		{3200, 0.3200},
		{3201, 0.3201},
		{12401, 1.2401},
		{9999, 0.9999},
		{-567, -0.0567},
		{0, 0.0000},
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
		{9.7788, 97788},
		{9.77, 97700},
		{9.771, 97710},
		{9.7750001, 97750},
		{9.7749999, 97750},
		{0.0, 0},
		{-0.0, 0},
		{-10.40, -104000},
		{-155.62, -1556200},
		{-9.771, -97710},
		{-9.7749999, -97750},
		{-9.7750001, -97750},
	}

	for _, test := range tests {
		converted := FromFloat64(test.float64)
		if converted != test.money {
			t.Errorf(`Error converting float64 to money: result "%d", expected "%d"`, converted, test.money)
		}
	}
}
