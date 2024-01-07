package main

import (
	"fmt"
	"strings"

	"github.com/sintemal/go_meter_parsing/formats"
	"github.com/sintemal/go_meter_parsing/formats/mennekes"
	"github.com/sintemal/go_meter_parsing/formats/ocmf"
	"github.com/sintemal/go_meter_parsing/formats/pcdf"
)

func Parse(data string) (formats.MeterFormat, error) {
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
