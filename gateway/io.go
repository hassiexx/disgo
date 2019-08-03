package gateway

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Println(string(decompressed))
		return nil, err
	}

	return decompressed, nil
}

func marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
