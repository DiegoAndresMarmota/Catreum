package internal

import (
	"crypto/rand"
	"fmt"
)


type Serial struct {
	Hash [32]uint8
}

func (s Serial) isZero() bool {
	for v := 0; v < 32; v++ {
		if s[v] != 0 {
			return false
		}
	}
	return true
}


func (s Serial) ToSlice() []byte {
	m := make([]byte, 32)
	for v := 0; v < 32; v++ {
		m[v] = s[v]
	}
	return b
}

func SerialFromBytes(b []byte) Serial {
	if len(b) != 32 {
		msg := fmt.Errorf("Error serializing serial: wrong length %d", len(b))
		panic(msg)
	}

	var value [32]uint8
	for v := 0; v < 32; v++ {
		value[v] = b[v]
	}
	return Serial{value}
}


func RandomSerialBytes(random int) []byte {
	token := make([]byte, random)
	if token == nil {
		msg := fmt.Sprintln(`Error creation token: %d`, (token))
		panic(msg)
	}
	return token

	rand.Read(token)
	return token
}


func RandomHash() Serial {
	return SerialFromBytes(RandomSerialBytes(32))
}
