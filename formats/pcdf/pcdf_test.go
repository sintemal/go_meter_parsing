package pcdf

import (
	"testing"
)

// Test example for parsing apayloads. Values taken from Transparenzsoftware.
func TestParsePayloads(t *testing.T) {

	payload := "128.8.0(ST:210902105044)(CT:210902105118)(CD:000034)(TV:1)(BV:1)(CSC:18)(SP:1)(RV:0000.619*kWh)(SI:DVE911*4*a103507e-161e-414a-b8c5-e1efb37fb72d)(CS:5febd3dc)(HW:39202000189)(DT:0)(PK:04e4c95c1ca877c9b8237ccc9bed242f1ff6b87e1988bbe5dcd76de4ee6aa4eac3d7a120708d04857d63ae75eba8e2d5c99512d322ae409f8f77387835234e4c96)(SG:304602210092f1efe8e275a37fa21c64eaa9a3f8191c22f36f661cabc159bc8649b1ab36cc02210088bef59ed13eb633c994489736c033eba4f247361fba1ccabbaf0004d16d7b70)"

	_, err := Parse(payload)
	if err != nil {
		t.Errorf("Parse() error = %v", err)
	}

}
