package formats

import (
	"time"
)

type MeterFormat interface {
	MeterValue() float64
	MeterUnit() string // one of kWh, Wh, mOhm, uOhm
	Signature() string
	Start() time.Time
	Stop() time.Time
}
