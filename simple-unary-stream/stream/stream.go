// package stream implements basic stream read/write functions for
// grpc
package stream

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

// DefaultMaximumMessageSize is the default of maximum message size
const DefaultMaximumMessageSize = 1024 * 1024

// Stream represents bidirectional grpc communication
type Stream struct {
	reader      io.Reader
	writer      io.Writer
	maxBody     uint32
	compression bool
}

// New creates a new stream. It is not safe to use across multiple goroutines.
func New(o interface{}, compression bool) *Stream {
	s := Stream{
		maxBody:     DefaultMaximumMessageSize,
		compression: compression,
	}
	r, ok := o.(io.Reader)
	if r != nil && ok {
		s.reader = r
	}

	w, ok := o.(io.Writer)
	if w != nil && ok {
		s.writer = w
	}

	return &s
}

func (s *Stream) SetMaximumMessageSize(size uint32) {
	s.maxBody = size
}

// Read a single grpc message from the stream
func (s *Stream) Read(m proto.Message) error {
	if s.reader == nil {
		return errors.New("stream does not support reads")
	}

	prefix := []byte{0, 0, 0, 0, 0}
	_, err := s.reader.Read(prefix)
	if err != nil {
		return err
	}

	length := binary.BigEndian.Uint32(prefix[1:])
	if length == 0 {
		return errors.New("0 length")
	}

	if length > s.maxBody {
		return fmt.Errorf("length of %d is greater than maximum of %d",
			length, s.maxBody)
	}

	body := make([]byte, length)
	_, err = s.reader.Read(body)
	if err != nil {
		return err
	}

	if s.compression && prefix[0] == 1 {
		r, err := gzip.NewReader(bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		defer r.Close()

		data, err := ioutil.ReadAll(r)
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

// Write a single grpc message to the stream
func (s *Stream) Write(m proto.Message) error {
	if s.writer == nil {
		return errors.New("stream does not support writes")
	}

	body, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	if uint32(len(body)) > s.maxBody {
		return fmt.Errorf("length of %d is greater than maximum of %d",
			len(body), s.maxBody)
	}

	prefix := []byte{0, 0, 0, 0, 0}

	if s.compression {
		prefix[0] = 1
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		_, err := w.Write(body)
		if err != nil {
			w.Close()
			return err
		}
		if err := w.Close(); err != nil {
			return err
		}

		body = b.Bytes()
	}

	binary.BigEndian.PutUint32(prefix[1:], uint32(len(body)))

	_, err = s.writer.Write(prefix)
	if err != nil {
		return err
	}

	_, err = s.writer.Write(body)
	if err != nil {
		return err
	}

	return nil
}
