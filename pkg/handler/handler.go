package handler

import (
	"github.com/goflame/flame/pkg/http"
	"github.com/goflame/flame/pkg/http/response"
)

type Handler func(*http.Response, *http.Request) *response.Err
