package formats

import (
	"fmt"
	"strings"
	"time"

	"github.com/sintemal/go_meter_parser/formats/mennekes"
	"github.com/sintemal/go_meter_parser/formats/ocmf"
	"github.com/sintemal/go_meter_parser/formats/pcdf"
)

type MeterFormat interface {
	MeterValue() float64
	MeterUnit() string // one of kWh, Wh, mOhm, uOhm
	Signature() string
	Start() time.Time
	Stop() time.Time
}

func Parse(data string) (MeterFormat, error) {
	// Check, which format we have
	if strings.HasPrefix(data, "128.8.0") {
		return pcdf.Parse(data)
	} else if strings.HasPrefix(data, "OCMF|") {
		return ocmf.Parse(data)
	} else if strings.HasPrefix(data, "<ChargingProcess") {
		return mennekes.Parse(data)
	} else {
		return nil, fmt.Errorf("Unknown format")
	}
}
