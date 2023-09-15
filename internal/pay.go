package internal

import "io"

type Pay struct{}

func (p *Pay) Serialize(se io.Reader) error {
	return nil
}

func (p *Pay) Deserialize(de io.Reader) error {
	return nil
}