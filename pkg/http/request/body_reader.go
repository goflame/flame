package request

import (
	"encoding/json"
	"github.com/goflame/flame/pkg/http/response"
	"io"
)

type BodyReader struct {
	body *io.ReadCloser
}

func NewBodyReader(b *io.ReadCloser) *BodyReader {
	return &BodyReader{
		body: b,
	}
}

func (b *BodyReader) Byte() ([]byte, *response.Err) {
	bt, err := io.ReadAll(*b.body)
	if err != nil {
		return nil, response.NewError(err.Error()).Status(400)
	}
	return bt, nil
}

func (b *BodyReader) JSON(s *any) *response.Err {
	err := json.NewDecoder(*b.body).Decode(s)
	if err != nil {
		return response.NewError(err.Error()).Status(400)
	}
	return nil
}
