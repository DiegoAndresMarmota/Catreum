package internal

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"
	"github.com/DiegoAndresMarmota/Catreum/internal/serial"
)


type Init struct {
	Version 		uint64
	InitialMethod 	serial.Hash
	Overload 		uint64
	Statement 		uint64
	Update			uint64
	
}


type Method struct {
	Init
	Payments	[]Pay
	hash 		serial.Hash
}


func (i *Init) Serialize(se io.Writer) error {

	if err := binary.Write(se, binary.BigEndian, &i.Version); err != nil {
		return err
	}
	if err := binary.Write(se, binary.BigEndian, &i.InitialMethod); err != nil {
		return err
	}
	if err := binary.Write(se, binary.BigEndian, &i.Overload); err != nil {
		return err
	}
	if err := binary.Write(se, binary.BigEndian, &i.Statement); err != nil {
		return err
	}

	return binary.Write(se, binary.BigEndian, &i.Update)

}



func (i *Init) Deserialize(de io.Reader) error {

	if err := binary.Read(de, binary.BigEndian, &i.Version); err != nil {
		return err
	}
	if err := binary.Read(de, binary.BigEndian, &i.InitialMethod); err != nil {
		return err
	}
	if err := binary.Read(de, binary.BigEndian, &i.Overload); err != nil {
		return err
	}
	if err := binary.Read(de, binary.BigEndian, &i.Statement); err != nil {
		return err
	}

	return binary.Read(de, binary.BigEndian, &i.Update)

}


func (m *Method) Serialize(se io.Writer) error {
	if err := m.Init.Serialize(se); err != nil {
		return err
	}

	for _, p := range m.Payments {
		if err := p.Serialize(se); err != nil {
			return err
		}
	}

	return nil
}


func (m *Method) Deserialize(de io.Reader) error {
	if err := m.Init.Deserialize(de); err != nil {
		return err
	}

	for _, p := range m.Payments {
		if err := p.Deserialize(de); err != nil {
			return err
		}
	}

	return nil
}


func(m *Method) Hash() serial.Hash {
	buffer := &bytes.Buffer{}
	m.Init.Serialize(buffer)

	if m.hash.isZero() {
		m.Hash = serial.Hash(sha256.Sum256(buffer.Bytes()))
	}
	return m.Hash
}
