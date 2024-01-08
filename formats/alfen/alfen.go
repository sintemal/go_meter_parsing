package alfen

import (
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	AdapterIDOffset               = 0
	AdapterIDLength               = 10
	AdapterFirmwareVersionOffset  = 10
	AdapterFirmwareVersionLength  = 4
	AdapterFirmwareChecksumOffset = 14
	AdapterFirmwareChecksumLength = 2
	MeterIDOffset                 = 16
	MeterIDLength                 = 10
	StatusOffset                  = 26
	StatusLength                  = 4
	SecondIndexOffset             = 30
	SecondIndexLength             = 4
	TimestampOffset               = 34
	TimestampLength               = 4
	ObisIDOffset                  = 38
	ObisIDLength                  = 6
	UnitOffset                    = 44
	ScalarOffset                  = 45
	ValueOffset                   = 46
	ValueLength                   = 8
	UIDOffset                     = 54
	UIDLength                     = 20
	SessionIDOffset               = 74
	SessionIDLength               = 4
	PagingOffset                  = 78
	PagingLength                  = 4

	DatasetLength = 82
)

// Format is like "AP;0;3;AMVBBEIORR2RGJLJ6YRZUGACQAXSDFCL66EIP3N7;BJKGK43UIRSXMAAROYYDCMZNUYFACRC2I4ADGAABIAAAAAAAQ6ACMAD5CH4FWAIAAEEAB7Y6ACJD2AAAAAAAAABRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASAAAAAEIAAAAA====;IGRCBV3TL45XIGPJU7QGD3H4V6ICQ75GLPWEFNEKZX3RTTKJI2FBXHPCWUIWL5OENEHE3SQRVACHG===;
type AlfenData struct {
	Identifier              string
	Type                    int // 0 = start, 1 = stop, 2 = intermediate
	Pubkey                  string
	AdapterId               []byte
	AdatperFirmwareVersion  string
	AdapterFirmwareChecksum []byte
	MeterId                 []byte
	Status                  []byte
	SecondIndex             uint32
	Timestamp               time.Time
	ObisId                  []byte
	Unit                    []byte
	Scalar                  int8
	Value                   int64
	UID                     string
	SessionId               int64
	Paging                  int64
	Dataset                 string
	Sig                     string
}

func Parse(data string) (AlfenData, error) {
	parts := strings.Split(data[:len(data)-1], ";")
	if len(parts) != 6 {
		return AlfenData{}, fmt.Errorf("invalid parts format")
	}

	identifier := parts[0]

	if identifier != "AP" {
		return AlfenData{}, fmt.Errorf("invalid AP format")
	}
	typ := parts[1]
	typInt, err := strconv.Atoi(typ)
	if err != nil {
		return AlfenData{}, err
	}
	pubkey := parts[3]
	dataset := parts[4]
	signature := parts[5]

	dataset_dec, err := base32.StdEncoding.DecodeString(dataset)
	if err != nil {
		return AlfenData{}, err
	}

	if len(dataset_dec) != DatasetLength {
		return AlfenData{}, fmt.Errorf("invalid dataset length")
	}

	adapterId := dataset_dec[AdapterIDOffset : AdapterIDOffset+AdapterIDLength]

	adapterFirmwareVersion := string(dataset_dec[AdapterFirmwareVersionOffset : AdapterFirmwareVersionOffset+AdapterFirmwareVersionLength])
	adapterFirmwareChecksum := dataset_dec[AdapterFirmwareChecksumOffset : AdapterFirmwareChecksumOffset+AdapterFirmwareChecksumLength]
	meterId := dataset_dec[MeterIDOffset : MeterIDOffset+MeterIDLength]
	status := dataset_dec[StatusOffset : StatusOffset+StatusLength]

	//secondIndex is little endian u32
	secondIndex := binary.LittleEndian.Uint32(dataset_dec[SecondIndexOffset : SecondIndexOffset+SecondIndexLength])

	//timestamp is little endian u32
	timestampInt := binary.LittleEndian.Uint32(dataset_dec[TimestampOffset : TimestampOffset+TimestampLength])
	timestamp := time.Unix(int64(timestampInt), 0)

	obisId := dataset_dec[ObisIDOffset : ObisIDOffset+ObisIDLength]
	unit := dataset_dec[UnitOffset : UnitOffset+1]
	//scalar is just the byte, not a string
	scalar := int8(dataset_dec[ScalarOffset])

	//value is little endian u64
	value := int64(binary.LittleEndian.Uint64(dataset_dec[ValueOffset : ValueOffset+ValueLength]))

	uid := string(dataset_dec[UIDOffset : UIDOffset+UIDLength])

	//sessionId is little endian u32
	sessionId := int64(binary.LittleEndian.Uint32(dataset_dec[SessionIDOffset : SessionIDOffset+SessionIDLength]))

	//paging is little endian u32
	paging := int64(binary.LittleEndian.Uint32(dataset_dec[PagingOffset : PagingOffset+PagingLength]))

	//encode signature as hex

	signatureHex := hex.EncodeToString([]byte(signature))

	return AlfenData{
		Type:                    typInt,
		Identifier:              identifier,
		AdapterId:               adapterId,
		AdatperFirmwareVersion:  adapterFirmwareVersion,
		AdapterFirmwareChecksum: adapterFirmwareChecksum,
		MeterId:                 meterId,
		Status:                  status,
		SecondIndex:             secondIndex,
		Timestamp:               timestamp,
		ObisId:                  obisId,
		Unit:                    unit,
		Scalar:                  scalar,
		Value:                   value,
		UID:                     uid,
		SessionId:               sessionId,
		Paging:                  paging,
		Dataset:                 dataset,
		Sig:                     signatureHex,
		Pubkey:                  pubkey,
	}, nil
}

func (a AlfenData) MeterValue() float64 {
	return float64(a.Value) * math.Pow10(int(a.Scalar)) / 1000.0 //convert to kWH
}

func (a AlfenData) MeterUnit() string {
	return "kWh"
}

func (a AlfenData) Signature() string {
	return a.Sig
}

func (a AlfenData) Start() time.Time {
	return a.Timestamp
}

func (a AlfenData) Stop() time.Time {
	return a.Timestamp
}

func (a AlfenData) Name() string {
	return "alfen"
}
