package gateway

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"io/ioutil"
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
		return nil, err
	}

	return buf.Bytes(), nil
}

func decompress(data []byte) ([]byte, error) {
	// Create zlib reader.
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Close reader before returning.
	defer reader.Close()

	// Read from zlib reader.
	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		// Ignore unexpected EOF errors.
		// Go zlib does not treat the zlib suffix 0 0 255 255 as end of file.
		if err != io.ErrUnexpectedEOF {
			return nil, err
		}
	}

	return decompressed, nil
}

func marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
