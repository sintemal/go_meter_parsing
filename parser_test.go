package go_meter_parsing

import (
	"testing"
)

func TestWrapTransparenzsoftware(t *testing.T) {
	payload := `OCMF|{"FV":"1.0","GI":"LEM DCBM","GS":"1211751603","GV":"v1","PG":"T144","MV":"LEM","MS":"1211751603","MF":"MU-0.1.4.0_SU-0.0.8.0","IS":true,"IL":"HEARSAY","IF":["RFID_RELATED","OCPP_AUTH_TLS","ISO15118_NONE","PLMN_NONE"],"IT":"ISO14443","ID":"5EEFE0C7F64B050E9FB95C","CT":"EVSEID","CI":"458877","RD":[{"TM":"2021-10-06T13:13:56,000+0200 R","TX":"B","RV":57.584,"RI":"1-0:1.8.0","RU":"kWh","RT":"DC","EF":"","ST":"G","UC":{"UN":"No_Comp","UI":2,"UR":0}},{"RV":4.405,"RI":"1-0:2.8.0","RU":"kWh","ST":"G"},{"TM":"2021-10-06T13:15:13,000+0200 R","TX":"E","RV":58.685,"RI":"1-0:1.8.0","RU":"kWh","ST":"G"},{"RV":4.405,"RI":"1-0:2.8.0","RU":"kWh","ST":"G"}]}|{"SA":"ECDSA-secp256r1-SHA256","SD":"3045022100E1C80E7115294A3A4D10CFF8B94E959F941415E19A0913A0538312E585E1A37702200A2BE1B235C152CFC7555064A96533F192A8A1D4E0484C68FD46574947A3D9BA"}`

	meter_format, err := Parse(payload)
	if err != nil {
		t.Errorf("Parse() error = %v", err)
	}

	out, err := WrapTransparenzsoftware(meter_format, payload, "pubkey")

	if err != nil {
		t.Errorf("WrapTransparenzsoftware() error = %v", err)
	}
	t.Logf("WrapTransparenzsoftware() output = %v", out)
}
