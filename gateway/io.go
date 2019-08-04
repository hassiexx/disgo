package gateway

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"io/ioutil"

	"golang.org/x/xerrors"
)

func compress(data []byte) ([]byte, error) {
	// Create buffer.
	var buf bytes.Buffer

	// Create zlib writer.
	writer := zlib.NewWriter(&buf)

	// Flush and close writer before returning.
	defer writer.Flush()
	defer writer.Close()

	// Compress data.
	_, err := writer.Write(data)
	if err != nil {
		return nil, xerrors.Errorf("failed to compress data using zlib: %v", err)
	}

	return buf.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
	// Create zlib reader.
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, xerrors.Errorf("failed to create zlib reader: %v", err)
	}

	// Close reader before returning.
	defer reader.Close()

	// Read from zlib reader.
	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		// Ignore unexpected EOF errors.
		// Go zlib does not treat the zlib suffix 0 0 255 255 as end of file.
		if err != io.ErrUnexpectedEOF {
			return nil, xerrors.Errorf("failed to decompress zlib stream: %v", err)
		}
	}

	return decompressed, nil
}

func marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, xerrors.Errorf("failed to marshal json: %v", err)
	}
	return data, nil
}

func unmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return xerrors.Errorf("failed to unmarshal json: %v", err)
	}
	return nil
}
