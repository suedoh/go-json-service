package main

import (
	"encoding/json"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

func NewRequest(method string, params string) *Request {
	return &Request{
		Method: method,
		Params: params,
	}
}

func (r *Request) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Request) FromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

