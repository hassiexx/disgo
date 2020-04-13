package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/klauspost/compress/zlib"
	"golang.org/x/xerrors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// RawMessage is an alias for jsoniter.RawMessage
type RawMessage jsoniter.RawMessage

// Marshal encodes v as JSON into w.
func Marshal(w io.Writer, v interface{}) error {
	if err := NewEncoder(w).Encode(v); err != nil {
		return xerrors.Errorf("marshal: %w", err)
	}

	return nil
}

// NewDecoder creates a new JSON decoder.
func NewDecoder(r io.Reader) *jsoniter.Decoder {
	return json.NewDecoder(r)
}

// NewEncoder creates a new JSON encoder.
func NewEncoder(w io.Writer) *jsoniter.Encoder {
	return json.NewEncoder(w)
}

// Unmarshal decodes JSON from a reader into v.
func Unmarshal(r io.Reader, v interface{}) error {
	if err := NewDecoder(r).Decode(v); err != nil {
		return xerrors.Errorf("unmarshal: %w", err)
	}

	return nil
}

// UnmarshalRaw decodes a raw JSON message into v.
func UnmarshalRaw(raw RawMessage, v interface{}) error {
	if err := json.Unmarshal(raw, v); err != nil {
		return xerrors.Errorf("unmarshal raw json message: %w", err)
	}

	return nil
}

// UnmarshalZlib is a helper func to decompress a zlib stream
// and unmarshal JSON.
func UnmarshalZlib(r io.Reader, v interface{}) error {
	// Create zlib reader.
	zlibR, err := zlib.NewReader(r)
	if err != nil {
		return xerrors.Errorf("create zlib reader: %w", err)
	}

	defer zlibR.Close()

	// Unmarshal.
	if err = NewDecoder(r).Decode(v); err != nil {
		return xerrors.Errorf("unmarshal from zlib reader: %w", err)
	}

	return nil
}
