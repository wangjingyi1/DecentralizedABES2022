package DecentralizedABES

import (
	"hash"

	"github.com/Nik-U/pbc"
)

type Cipher struct {
	C0         *pbc.Element   `field:"2"`
	C1s        []*pbc.Element `field:"2"`
	C2s        []*pbc.Element `field:"0"`
	C3s        []*pbc.Element `field:"0"`
	CipherText []byte
	Policy     string
}

type Signature struct {
	Sig0      *pbc.Element
	Sig1xs    []*pbc.Element
	Sig2xs    []*pbc.Element
	upTemps   []*pbc.Element
	downTemps []*pbc.Element
	Policy    string
	M         string
	Hash      hash.Hash
	Attrs     []string
	s         *pbc.Element
}
