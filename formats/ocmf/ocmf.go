package ocmf

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type OCMFString struct {
	Payload Payload
	Sig     Signature
}

// LossCompensation structure
type LossCompensation struct {
	Name           string  `json:"LN,omitempty"`
	Identifier     float64 `json:"LI,omitempty"`
	Resistance     float64 `json:"LR"`
	ResistanceUnit string  `json:"LU"`
}

// Reading structure
type Reading struct {
	Time           string  `json:"TM"`
	Transaction    string  `json:"TX,omitempty"`
	Value          float64 `json:"RV"`
	Identification string  `json:"RI,omitempty"`
	Unit           string  `json:"RU,omitempty"`
	CurrentType    string  `json:"RT,omitempty"`
	CumulatedLoss  float64 `json:"CL,omitempty"`
	ErrorFlags     string  `json:"EF,omitempty"`
	Status         string  `json:"ST"`
}

// UserAssignment structure
type UserAssignment struct {
	Status     bool     `json:"IS"`
	Level      string   `json:"IL,omitempty"`
	Flags      []string `json:"IF,omitempty"`
	Type       string   `json:"IT,omitempty"`
	Data       string   `json:"ID,omitempty"`
	TariffText string   `json:"TT,omitempty"`
}

// Payload structure
type Payload struct {
	FormatVersion         string           `json:"FV"`
	GatewayIdentification string           `json:"GI,omitempty"`
	GatewaySerial         string           `json:"GS,omitempty"`
	GatewayVersion        string           `json:"GV,omitempty"`
	Pagination            string           `json:"PG,omitempty"`
	MeterVendor           string           `json:"MV,omitempty"`
	MeterModel            string           `json:"MM,omitempty"`
	MeterSerial           string           `json:"MS,omitempty"`
	MeterFirmware         string           `json:"MF,omitempty"`
	UserAssignment        UserAssignment   `json:"UserAssignment,omitempty"`
	LossCompensation      LossCompensation `json:"LC,omitempty"`
	Readings              []Reading        `json:"RD,omitempty"`
}

type Signature struct {
	Signature string `json:"SD,omitempty"`
}

// https://github.com/SAFE-eV/OCMF-Open-Charge-Metering-Format/blob/master/OCMF-en.md
func Parse(data string) (OCMFString, error) {
	sections := strings.Split(data, "|")
	if len(sections) != 3 || sections[0] != "OCMF" {
		return OCMFString{}, fmt.Errorf("invalid format")
	}

	var payload Payload

	err := json.Unmarshal([]byte(sections[1]), &payload)

	if err != nil {
		return OCMFString{}, err
	}

	beginning, err := payload.getBeginning()
	if err != nil {
		return OCMFString{}, err
	}

	_, err = time.Parse("2006-01-02T15:04:05,000-0700", beginning.Time[:len(beginning.Time)-2]) //cut off synchronization flag
	if err != nil {
		return OCMFString{}, err
	}

	end, err := payload.getEnd()
	if err != nil {
		return OCMFString{}, err
	}

	_, err = time.Parse("2006-01-02T15:04:05,000-0700", end.Time[:len(end.Time)-2])
	if err != nil {
		return OCMFString{}, err
	}

	var signature Signature
	err = json.Unmarshal([]byte(sections[2]), &signature)
	if err != nil {
		return OCMFString{}, err
	}

	return OCMFString{
		Payload: payload,
		Sig:     signature,
	}, nil
}

func (s Payload) getEnd() (Reading, error) {
	stopCodes := map[string]bool{"B": true, "E": true, "L": true, "R": true, "A": true, "P": true}
	if s.Pagination[0] != 'T' {
		return Reading{}, fmt.Errorf("No transaction context found in pagination")
	}
	for _, reading := range s.Readings {
		_, ok := stopCodes[reading.Transaction]
		if ok {
			return reading, nil
		}
	}
	return Reading{}, fmt.Errorf("No end found")
}

func (s Payload) getBeginning() (Reading, error) {
	if s.Pagination[0] != 'T' {
		return Reading{}, fmt.Errorf("No transaction context found in pagination")
	}
	for _, reading := range s.Readings {
		if reading.Transaction == "B" {
			return reading, nil
		}
	}
	return Reading{}, fmt.Errorf("No beginning found")
}
func (s OCMFString) MeterValue() float64 {
	beginning, _ := s.Payload.getBeginning() //no error checking needed as we already checked in Parse()
	end, _ := s.Payload.getEnd()

	return end.Value - beginning.Value
}

func (s OCMFString) MeterUnit() string {
	return s.Payload.Readings[0].Unit
}

func (s OCMFString) Signature() string {
	return s.Sig.Signature
}

func (s OCMFString) Start() time.Time {
	beginning, _ := s.Payload.getBeginning() //no error checking needed as we already checked in Parse()
	t, _ := time.Parse("2006-01-02T15:04:05,000-0700", beginning.Time[:len(beginning.Time)-2])
	return t

}

func (s OCMFString) Stop() time.Time {
	end, _ := s.Payload.getEnd() //no error checking needed as we already checked in Parse()
	t, _ := time.Parse("2006-01-02T15:04:05,000-0700", end.Time[:len(end.Time)-2])
	return t
}

func (s OCMFString) Name() string {
	return "OCMF"
}

func (s OCMFString) String() (string, error) {
	payloadMarshal, err := json.Marshal(s.Payload)
	if err != nil {
		return "", err
	}
	signature, err := json.Marshal(s.Sig)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("OCMF|%s|%s", payloadMarshal, signature), nil
}
