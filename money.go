package money

import (
	"strconv"
	"math"
	"fmt"
	"errors"
)

//TODO: add doc
// Хранит деньги как Int, возвращает деленные на 100
type Money int64

var (
	roundOn float64 = 0.5
	places int = 2
)

//todo: add json/decode/encode
func (m *Money) UnmarshalJSON(data []byte) error {
	val, err := strconv.ParseFloat(string(data),64)
	if err != nil {
		return errors.New("Money.UnmarshalJSON : " + err.Error())
	}
	*m = Money(round(val)*100)
	return nil
}

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m Money) String() string {
	return fmt.Sprintf("%0.2f", m.Float64())
}

func (m Money) Float64() float64 {
	return float64(m) / 100
}

func FromFloat64(v float64) Money {
	if v >= 0 {
		return Money(v * 100 + 0.5)
	} else {
		return Money(v * 100 - 0.5)
	}
}

func round(val float64) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
