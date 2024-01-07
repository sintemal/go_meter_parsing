package pcdf

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Porsche Charging Data Format
type PCDF struct {
	StartTime        time.Time     // ST
	StopTime         time.Time     // CT
	ChargingDuration time.Duration // CD
	BillValid        bool          // BV
	StopValid        bool          // SV
	ChargingCount    int           // CSC
	ChargeSoftware   string        // CS
	Hardware         string        // HW
	DCMeter          string        // DT
	signature        string        // SG
	ReadingValue     float64       // RV
}

func extractAttribute(data string, identifier string) (string, error) {
	pos := strings.Index(data, fmt.Sprintf("(%s", identifier))
	if pos == -1 {
		return "", fmt.Errorf("identifier %s not found", identifier)
	}
	end := strings.Index(data[pos:], ")")
	substring := data[pos+1 : end+pos]

	split := strings.Split(substring, ":")
	if len(split) != 2 {
		return "", fmt.Errorf("invalid attribute format")
	}

	return split[1], nil
}

func Parse(data string) (PCDF, error) {
	if !strings.HasPrefix(data, "128.8.0") {
		return PCDF{}, fmt.Errorf("invalid PCDF data")
	}

	startTime, err := extractAttribute(data, "ST")
	if err != nil {
		return PCDF{}, err
	}
	parsedStartTime, err := time.Parse("20060102150405", fmt.Sprintf("20%s", startTime))
	if err != nil {
		return PCDF{}, err
	}

	stopTime, err := extractAttribute(data, "CT")
	if err != nil {
		return PCDF{}, err
	}
	parsedStopTime, err := time.Parse("20060102150405", fmt.Sprintf("20%s", stopTime))
	if err != nil {
		return PCDF{}, err
	}

	duration, err := extractAttribute(data, "CD")
	if err != nil {
		return PCDF{}, err
	}
	parsedDuration, err := time.ParseDuration(fmt.Sprintf("%ss", duration))

	billValid, err := extractAttribute(data, "BV")
	if err != nil {
		return PCDF{}, err
	}
	billValidBool := billValid == "1"

	stopValid, err := extractAttribute(data, "SP")
	if err != nil {
		return PCDF{}, err
	}
	stopValidBool := stopValid == "1"

	chargingCount, err := extractAttribute(data, "CSC")
	if err != nil {
		return PCDF{}, err
	}
	parsedChargingCount, err := strconv.Atoi(chargingCount)

	chargeSoftware, err := extractAttribute(data, "CS")
	if err != nil {
		return PCDF{}, err
	}

	hardware, err := extractAttribute(data, "HW")
	if err != nil {
		return PCDF{}, err
	}

	dcMeter, err := extractAttribute(data, "DT")
	if err != nil {
		return PCDF{}, err
	}

	signature, err := extractAttribute(data, "SG")
	if err != nil {
		return PCDF{}, err
	}

	readingValue, err := extractAttribute(data, "RV")
	if err != nil {
		return PCDF{}, err
	}

	readingParts := strings.Split(readingValue, "*")
	if len(readingParts) != 2 || readingParts[1] != "kWh" {
		return PCDF{}, fmt.Errorf("invalid reading value format")
	}

	parsedReadingValue, err := strconv.ParseFloat(readingParts[0], 64)
	if err != nil {
		return PCDF{}, err
	}

	pcdf := PCDF{
		StartTime:        parsedStartTime,
		StopTime:         parsedStopTime,
		ChargingDuration: parsedDuration,
		BillValid:        billValidBool,
		StopValid:        stopValidBool,
		ChargingCount:    parsedChargingCount,
		ChargeSoftware:   chargeSoftware,
		Hardware:         hardware,
		DCMeter:          dcMeter,
		signature:        signature,
		ReadingValue:     parsedReadingValue,
	}
	return pcdf, nil
}

func (p PCDF) MeterValue() float64 {
	return p.ReadingValue
}

func (p PCDF) MeterUnit() string {
	return "kWh"
}

func (p PCDF) Signature() string {
	return p.signature
}

func (p PCDF) Start() time.Time {
	return p.StartTime
}

func (p PCDF) Stop() time.Time {
	return p.StopTime
}
