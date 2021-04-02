package model

import (
	"encoding/json"
	"unsafe"
)

// Content the is various types of Content conversion
type Content interface {
	GetString() (string, bool)
	GetBytes() ([]byte, bool)
	Raw() interface{}
}

func NewContent(val interface{}) Content {
	switch v := val.(type) {
	case Content:
		return v
	case string:
		return stringContent(v)
	case []byte:
		return bytesContent(v)
	default:
		return &rawContent{
			raw: v,
		}
	}
}

type stringContent string

func (s stringContent) GetString() (string, bool) {
	return string(s), true
}

func (s stringContent) GetBytes() ([]byte, bool) {
	conv := (*[]byte)(unsafe.Pointer(&s))
	return *conv, true
}

func (s stringContent) Raw() interface{} {
	return string(s)
}

type bytesContent []byte

func (b bytesContent) GetString() (string, bool) {
	conv := (*string)(unsafe.Pointer(&b))
	return *conv, true
}

func (b bytesContent) GetBytes() ([]byte, bool) {
	return []byte(b), true
}

func (b bytesContent) Raw() interface{} {
	return []byte(b)
}

type rawContent struct {
	raw interface{}
}

func (r rawContent) GetString() (string, bool) {
	return "", false
}

func (r rawContent) GetBytes() ([]byte, bool) {
	return nil, false
}

func (r rawContent) Raw() interface{} {
	return r.raw
}

func (r rawContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.raw)
}
