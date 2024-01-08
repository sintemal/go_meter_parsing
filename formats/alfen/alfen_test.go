package alfen

import "testing"

func TestParsePayloads(t *testing.T) {

	samplePayloads := []string{
		"AP;0;3;AMVBBEIORR2RGJLJ6YRZUGACQAXSDFCL66EIP3N7;BJKGK43UIRSXMAAROYYDCMZNUYFACRC2I4ADGAABIAAAAAAAQ6ACMAD5CH4FWAIAAEEAB7Y6ACJD2AAAAAAAAABRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASAAAAAEIAAAAA====;IGRCBV3TL45XIGPJU7QGD3H4V6ICQ75GLPWEFNEKZX3RTTKJI2FBXHPCWUIWL5OENEHE3SQRVACHG===;",
		"AP;1;3;AICIVT423BX3TJGK6QCCVRHQ63LJQUEVZWWTYQUZ;BJKGK43UIRSXMAAROYYDCMZNUYFACRC2I4ADGAABIAAAAAIAQ6ACMAD5CH4FWAIAAEEAB7Y6ACKD2AAAAAAAAABRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASAAAAAEQAAAAA====;C3J3MLA5XLF7QYYHA4RAJV7QLBWU5OB3M3DUKCUTREEQ5QORE45DMUQALYBEI2YOLNX7DYFRWGLYU===;",
		"AP;1;3;AICIVT423BX3TJGK6QCCVRHQ63LJQUEVZWWTYQUZ;BJKGK43UIRSXMAAROYYDCMZNUYFACRC2I4ADGCABIAAAAAEAQ6ACMAD5CH4FWAIAAEEAB7Y6ACKD2AAAAAAAAABRAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASAAAAAEQAAAAA====;XVRQMWYJVQHI5GP53EPYUOYLHDAGBUTUIXJVRRZXKD7CIRXO6TYBFG7H43OT2STWBQW6MU7LMOBD2===;",
	}

	for _, payload := range samplePayloads {
		_, err := Parse(payload)
		if err != nil {
			t.Errorf("Parse() error = %v", err)
		}
	}
}
