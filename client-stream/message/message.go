// package message implements reading and writing of gRPC messages
package message

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// MaximumMessageSize is the maximum message size
const MaximumMessageSize = 1024 * 1024

// Read a single grpc message. This supports gzip compression.
func Read(r io.Reader, m proto.Message) error {
	// gRPC prefix is:
	// - one byte flag to denote compression
	// - length of the messaage as big endian unsigned 32 bit integer in 4 bytes
	prefix := []byte{0, 0, 0, 0, 0}
	_, err := r.Read(prefix)
	if err != nil {
		return err
	}

	length := binary.BigEndian.Uint32(prefix[1:])
	if length == 0 {
		return errors.New("0 length message")
	}

	if length > MaximumMessageSize {
		return fmt.Errorf("length of %d is greater than maximum of %d",
			length, MaximumMessageSize)
	}

	body := make([]byte, length)
	_, err = r.Read(body)
	if err != nil {
		return err
	}

	if prefix[0] == 1 {
		g, err := gzip.NewReader(bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		defer g.Close()

		data, err := ioutil.ReadAll(g)
		if err != nil {
			return err
		}

		body = data
	}

	if err := proto.Unmarshal(body, m); err != nil {
		return err
	}
	return nil
}

// Write a single grpc message.  Supports gzip compression
func Write(w io.Writer, m proto.Message, compress bool) error {
	body, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	// gRPC prefix is:
	// - one byte flag to denote compression
	// - length of the messaage as big endian unsigned 32 bit integer in 4 bytes
	//   this length is the compressed length if compressed
	prefix := []byte{0, 0, 0, 0, 0}

	if compress {
		prefix[0] = 1
		var b bytes.Buffer
		g := gzip.NewWriter(&b)
		_, err := g.Write(body)
		if err != nil {
			g.Close()
			return err
		}
		if err := g.Close(); err != nil {
			return err
		}

		body = b.Bytes()
	}

	if uint32(len(body)) > MaximumMessageSize {
		return fmt.Errorf("length of %d is greater than maximum of %d",
			len(body), MaximumMessageSize)
	}

	binary.BigEndian.PutUint32(prefix[1:], uint32(len(body)))

	_, err = w.Write(prefix)
	if err != nil {
		return err
	}

	_, err = w.Write(body)
	if err != nil {
		return err
	}

	return nil
}
