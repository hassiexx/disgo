package gateway

import (
	"compress/zlib"
	"encoding/json"
	"io"

	"golang.org/x/xerrors"
)

func decompressUnmarshal(r io.Reader, v interface{}) error {
	// Create zlib reader.
	reader, err := zlib.NewReader(r)
	if err != nil {
		return xerrors.Errorf("failed to create zlib reader: %w", err)
	}

	// Close reader before returning.
	defer reader.Close()

	// Unmarshal.
	if err = json.NewDecoder(reader).Decode(v); err != nil {
		return xerrors.Errorf("failed to unmarshal decompressed zlib: %w", err)
	}

	return nil
}

func marshal(w io.Writer, v interface{}) error {
	// Marshal.
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return xerrors.Errorf("failed to marshal json: %w", err)
	}

	return nil
}

func unmarshal(r io.Reader, v interface{}) error {
	// Unmarshal.
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return xerrors.Errorf("failed to unmarshal json: %w", err)
	}

	return nil
}

func unmarshalRaw(raw json.RawMessage, v interface{}) error {
	// Unmarshal.
	if err := json.Unmarshal(raw, v); err != nil {
		return xerrors.Errorf("failed to unmarshal json: %w", err)
	}

	return nil
}

/*
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
*/
