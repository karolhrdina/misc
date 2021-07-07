package ports

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

var (
	ErrJsonObjectExpected = errors.New("json object expected")
	ErrWrongPortIDType    = errors.New("string port ID exptected")
)

// PortParser parses ports from any stream (io.Reader) in json format.
// The methods are inspired by `database/sql` Rows `Next()/Scan()` methods.
type PortParser struct {
	dec         *json.Decoder
	err         error
	openingRead bool
}

// NewPortParser returns new initialized PortParser from `stream`
func NewPortParser(stream io.Reader) *PortParser {
	return &PortParser{
		dec: json.NewDecoder(stream),
		err: nil,
	}
}

// Next returns true if there are more ports to process,
// false otherwise.
func (p *PortParser) Next() bool {
	return p.err == nil && p.dec.More()
}

// Scan scans the next Port into `port` if Next() had returned true previously,
// io.EOF otherwise.
func (p *PortParser) Scan(port *Port) error {
	if !p.openingRead {
		// read opening token
		t, err := p.dec.Token()
		if err != nil {
			p.err = err
			return err
		}
		// make sure it's a json delimiter
		bracket, ok := t.(json.Delim)
		if !ok {
			p.err = ErrJsonObjectExpected
			return p.err
		}
		if bracket != '{' {
			p.err = ErrJsonObjectExpected
			return p.err
		}
		p.openingRead = true
	}

	// if there are more ports to read ...
	if p.dec.More() {
		t, err := p.dec.Token()
		if err != nil {
			p.err = err
			return err
		}
		id, ok := t.(string)
		if !ok {
			// since we had previously tested that we're inside a json object
			// and only properties with string key can be encountered here,
			// this check should never fail... but leaving this here if the
			// format would change and we could forget about this check
			p.err = ErrWrongPortIDType
			return p.err
		}

		// decode port
		err = p.dec.Decode(port)
		if err != nil {
			p.err = err
			return err
		}
		port.Key = id
		return nil
	}

	return io.EOF
}
