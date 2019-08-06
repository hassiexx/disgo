package gateway

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io/ioutil"

	"golang.org/x/xerrors"
)

func decompress(data []byte) ([]byte, error) {
	// Create zlib reader.
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, xerrors.Errorf("failed to create zlib reader: %w", err)
	}

	// Close reader before returning.
	defer reader.Close()

	// Read from zlib reader.
	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, xerrors.Errorf("failed to decompress zlib: %w", err)
	}

	return decompressed, nil
}

func marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, xerrors.Errorf("failed to marshal json: %w", err)
	}
	return data, nil
}

func unmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return xerrors.Errorf("failed to unmarshal json: %w", err)
	}
	return nil
}
