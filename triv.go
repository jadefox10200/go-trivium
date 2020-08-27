//An illustration of trivium. This however, is NOT memory efficient as it stores single bits in uint8's
//It does however, output the correct bit sequence per the spec as given in the book:
//Understanding Cryptography by Christof Paar.

//go version: go

package trivium

import (
	"fmt"	
)

type Triv struct {
	A	LFSR
	B	LFSR
	C	LFSR
}

type LFSR struct {	
	FF 	[]uint8
}

func (t *Triv) LoadKey(k [10]uint8) error {

	var ks string

	//key is loaded into leftmost 80 bits of B
	ks += fmt.Sprintf("%08b", k[0])	
	ks += fmt.Sprintf("%08b", k[1])
	ks += fmt.Sprintf("%08b", k[2])
	ks += fmt.Sprintf("%08b", k[3])
	ks += fmt.Sprintf("%08b", k[4])
	ks += fmt.Sprintf("%08b", k[5])
	ks += fmt.Sprintf("%08b", k[6])
	ks += fmt.Sprintf("%08b", k[7])
	ks += fmt.Sprintf("%08b", k[8])
	ks += fmt.Sprintf("%08b", k[9])		

	kb := []byte(ks)

	if len(kb) < 80 {return fmt.Errorf("must load 10 uint8 values")}

	for i := 0; i < 80; i++ {
		t.B.FF[i] = uint8(kb[i] - 48)
	}	

	return nil
}

func (t *Triv) Loadiv(iv [10]uint8) error {

	var is string

	//iv is loaded into leftmost 80 bits of A
	is += fmt.Sprintf("%08b", iv[0])	
	is += fmt.Sprintf("%08b", iv[1])
	is += fmt.Sprintf("%08b", iv[2])
	is += fmt.Sprintf("%08b", iv[3])
	is += fmt.Sprintf("%08b", iv[4])
	is += fmt.Sprintf("%08b", iv[5])
	is += fmt.Sprintf("%08b", iv[6])
	is += fmt.Sprintf("%08b", iv[7])
	is += fmt.Sprintf("%08b", iv[8])
	is += fmt.Sprintf("%08b", iv[9])		

	ib := []byte(is)

	if len(ib) < 80 {return fmt.Errorf("must load 10 uint8 values")}

	for i := 0; i < 80; i++ {
		//subtract 48 from each string value as 0 == 48  and 1 == 49 in ASCII
		t.A.FF[i] = uint8(ib[i] - 48)
	}	

	return nil
}

func NewTriv() Triv {

	la:= LFSR{		
		FF: make([]uint8, 93),		
	}

	lb:= LFSR{		
		FF: make([]uint8, 84),		
	}

	lc:= LFSR{		
		FF: make([]uint8, 111),		
	}

	lc.FF[108] = 1
	lc.FF[109] = 1
	lc.FF[110] = 1

	var t = Triv{
		A: la,
		B: lb,
		C: lc,
	}	

	return t

}

func(t *Triv) Init() {
	for i := 0; i < 1152; i++{
		t.Clock()
	}
}

func (t *Triv) Clock() uint8 {

	//get output
	a1 := t.A.FF[66-1] ^ t.A.FF[93-1] ^ (t.A.FF[91-1] & t.A.FF[92-1])
	b1 := t.B.FF[69-1] ^ t.B.FF[84-1] ^ (t.B.FF[82-1] & t.B.FF[83-1])
	c1 := t.C.FF[66-1] ^ t.C.FF[111-1] ^ (t.C.FF[109-1] & t.C.FF[110-1])
	si := a1 ^ b1 ^ c1

	//get input values:
	a0 := t.A.FF[69-1] ^ c1
	b0 := t.B.FF[78-1] ^ a1
	c0 := t.C.FF[87-1] ^ b1

	//shift over each lfsr: 
	for i := len(t.A.FF) - 1; i > 0; i-- {		
		t.A.FF[i] = t.A.FF[i-1]		
	}

	for i := len(t.B.FF) - 1; i > 0; i-- {		
		t.B.FF[i] = t.B.FF[i-1]
	}

	for i := len(t.C.FF) - 1; i > 0; i-- {		
		t.C.FF[i] = t.C.FF[i-1]
	}

	//load new inputs
	t.A.FF[0] = a0
	t.B.FF[0] = b0
	t.C.FF[0] = c0
	
	return si
}