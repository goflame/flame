package request

import (
	"github.com/goflame/flame/pkg/http/response"
)

type BodyReader interface {
	JSON(*interface{}) response.Err
	FormData(*interface{}) response.Err
	Byte() ([]byte, response.Err)
}
