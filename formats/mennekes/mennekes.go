package mennekes

import (
	"encoding/xml"
	"math"
	"strings"
	"time"
	"unicode"
)

type ChargingProcess struct {
	XMLName          xml.Name       `xml:"ChargingProcess"`
	ServerId         string         `xml:"ServerId"`
	PublicKey        string         `xml:"PublicKey"`
	MeteringPoint    string         `xml:"MeteringPoint"`
	SiteAddress      SiteAddress    `xml:"SiteAddress"`
	CustomerIdent    string         `xml:"CustomerIdent"`
	Timestamp        string         `xml:"TimestampCustomerIdent"`
	MeasurementStart Measurement    `xml:"MeasurementStart"`
	MeasurementEnd   Measurement    `xml:"MeasurementEnd"`
	EventLogItems    []EventLogItem `xml:"EventLogItems>EventLogItem"`
}

type SiteAddress struct {
	ZipCode string `xml:"ZipCode"`
	Street  string `xml:"Street"`
	Town    string `xml:"Town"`
}

type Measurement struct {
	Timestamp    string `xml:"Timestamp"`
	Signature    string `xml:"Signature"`
	EventCounter int    `xml:"EventCounter"`
	MeterStatus  int    `xml:"MeterStatus"`
	Value        int    `xml:"Value"`
	Scaler       int    `xml:"Scaler"`
	Pagination   int    `xml:"Pagination"`
	SecondIndex  int    `xml:"SecondIndex"`
}

type EventLogItem struct {
	TimeBefore   string `xml:"TimeBefore"`
	TimeAfter    string `xml:"TimeAfter"`
	MeterStatus  string `xml:"MeterStatus"`
	EventCounter string `xml:"EventCounter"`
	EventCode    string `xml:"EventCode"`
	Signature    string `xml:"Signature"`
	SecondIndex  string `xml:"SecondIndex"`
}

func Parse(data string) (ChargingProcess, error) {
	var cp ChargingProcess
	err := xml.Unmarshal([]byte(data), &cp)

	if err != nil {
		return ChargingProcess{}, err
	}
	//check correct timeformat
	_, err = time.Parse(time.RFC3339, strings.TrimFunc(cp.MeasurementStart.Timestamp, unicode.IsSpace))
	if err != nil {
		return ChargingProcess{}, err
	}

	_, err = time.Parse(time.RFC3339, strings.TrimFunc(cp.MeasurementEnd.Timestamp, unicode.IsSpace))
	if err != nil {
		return ChargingProcess{}, err
	}

	return cp, nil
}

func (cp ChargingProcess) MeterValue() float64 {
	return (float64(cp.MeasurementEnd.Value-cp.MeasurementStart.Value) / 1000.0) * math.Pow10(cp.MeasurementEnd.Scaler)
}

func (cp ChargingProcess) MeterUnit() string {
	return "kWh"
}

func (cp ChargingProcess) Signature() string {
	return cp.MeasurementEnd.Signature
}

func (cp ChargingProcess) Start() time.Time {
	t, _ := time.Parse(time.RFC3339, strings.TrimFunc(cp.MeasurementStart.Timestamp, unicode.IsSpace)) //no error checking needed as we already checked in Parse()
	return t
}

func (cp ChargingProcess) Stop() time.Time {
	t, _ := time.Parse(time.RFC3339, strings.TrimFunc(cp.MeasurementEnd.Timestamp, unicode.IsSpace)) //no error checking needed as we already checked in Parse()
	return t
}

func (s ChargingProcess) Name() string {
	return "Mennekes"
}
