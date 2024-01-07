package ocmf

import (
	"testing"
)

// Test example for parsing apayloads. Values taken from Transparenzsoftware.
func TestParsePayloads(t *testing.T) {

	samplePayloads := []string{
		`OCMF|{"FV":"1.0","GI":"LEM DCBM","GS":"1211751603","GV":"v1","PG":"T144","MV":"LEM","MS":"1211751603","MF":"MU-0.1.4.0_SU-0.0.8.0","IS":true,"IL":"HEARSAY","IF":["RFID_RELATED","OCPP_AUTH_TLS","ISO15118_NONE","PLMN_NONE"],"IT":"ISO14443","ID":"5EEFE0C7F64B050E9FB95C","CT":"EVSEID","CI":"458877","RD":[{"TM":"2021-10-06T13:13:56,000+0200 R","TX":"B","RV":57.584,"RI":"1-0:1.8.0","RU":"kWh","RT":"DC","EF":"","ST":"G","UC":{"UN":"No_Comp","UI":2,"UR":0}},{"RV":4.405,"RI":"1-0:2.8.0","RU":"kWh","ST":"G"},{"TM":"2021-10-06T13:15:13,000+0200 R","TX":"E","RV":58.685,"RI":"1-0:1.8.0","RU":"kWh","ST":"G"},{"RV":4.405,"RI":"1-0:2.8.0","RU":"kWh","ST":"G"}]}|{"SA":"ECDSA-secp256r1-SHA256","SD":"3045022100E1C80E7115294A3A4D10CFF8B94E959F941415E19A0913A0538312E585E1A37702200A2BE1B235C152CFC7555064A96533F192A8A1D4E0484C68FD46574947A3D9BA"}`,
		`OCMF|{"FV":"1.0","GI":"HTB","GS":"HTBSerial","GV":"0.0.1","PG":"T12345","MV":"HTB","MM":"HTB","MS":"HTBGenerated1","MF":"1.0","IS":true,"IL":"VERIFIED","IF":["RFID_PLAIN","OCPP_RS_TLS"],"IT":"ISO14443","ID":"8D2E2D536DCFD0","CI":"HTB","CT":"CBIDC","RD":[{"TM":"2019-04-02T17:28:45,867+0200 S","TX":"B","RV":3.51824,"RI":"1-b:1.8.e","RU":"kWh","EI":1523,"ST":"G"},{"TM":"2019-04-02T17:28:45,867+0200 S","TX":"X","RV":7.01047,"RI":"1-b:1.8.e","RU":"kWh","EI":1523,"ST":"G"}]}|{"SA":"ECDSA-secp192k1-SHA256","SD":"303402185A5E62A58B45C10C37166B7E43D26296D79C450DF30C2FA4021822099CFAEB793D9246FAE130C93964B43FC4C5F2CF1732E1"}`,
	}

	for _, payload := range samplePayloads {
		_, err := Parse(payload)
		if err != nil {
			t.Errorf("Parse() error = %v", err)
		}
	}

}
