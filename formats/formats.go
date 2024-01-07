package formats

import (
	"fmt"
	"time"
)

type MeterFormat interface {
	MeterValue() float64
	MeterUnit() string // one of kWh, Wh, mOhm, uOhm
	Signature() string
	Start() time.Time
	Stop() time.Time
}

func MeterValuekWh(m MeterFormat) (float64, error) {
	if m.MeterUnit() == "kWh" {
		return m.MeterValue(), nil
	} else if m.MeterUnit() == "Wh" {
		return m.MeterValue() / 1000, nil
	} else {
		return 0, fmt.Errorf("Invalid unit")
	}
}
