package go_meter_parsing

import (
	"encoding/xml"
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

func MeterValuekWh(m formats.MeterFormat) (float64, error) {
	if m.MeterUnit() == "kWh" {
		return m.MeterValue(), nil
	} else if m.MeterUnit() == "Wh" {
		return m.MeterValue() / 1000, nil
	} else {
		return 0, fmt.Errorf("Invalid unit")
	}
}

type SignedData struct {
	XMLName       xml.Name `xml:"signedData"`
	Format        string   `xml:"format,attr"`
	Encoding      string   `xml:"encoding,attr"`
	TransactionId string   `xml:"transactionId,attr"`
	Data          string   `xml:",chardata"`
}

// Define struct for publicKey
type PublicKey struct {
	XMLName xml.Name `xml:"publicKey"`
	Key     string   `xml:",chardata"`
}

// Define struct for value
type Value struct {
	XMLName       xml.Name   `xml:"value"`
	TransactionId string     `xml:"transactionId,attr"`
	Context       string     `xml:"context,attr"`
	SignedData    SignedData `xml:"signedData"`
	PublicKey     PublicKey  `xml:"publicKey,omitempty"`
}

// Define struct for values
type TransparenzsoftwareReadable struct {
	XMLName xml.Name `xml:"values"`
	Value   []Value  `xml:"value"`
}

func WrapTransparenzsoftware(m formats.MeterFormat, data string, publickey string) (string, error) {
	if m.Name() == "Mennekes" { //mennekes can be parsed directly by transparenzsoftware
		return data, nil
	}

	var pk PublicKey //sometimes the publickey is already included in the data
	if publickey != "" {
		pk = PublicKey{
			Key: publickey,
		}
	}

	signedData := SignedData{
		Format:        m.Name(),
		TransactionId: "1",
		Encoding:      "plain",
		Data:          data,
	}

	value := Value{
		TransactionId: "1",
		Context:       "Transaction.Begin",
		SignedData:    signedData,
		PublicKey:     pk,
	}

	values := TransparenzsoftwareReadable{
		Value: []Value{value},
	}

	x, err := xml.Marshal(values)
	if err != nil {
		return "", err
	}
	return string(x), nil
}
