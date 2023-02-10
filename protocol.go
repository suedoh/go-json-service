// protocol.go

package main

import (
	"encoding/json"
)

type Request struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

type Response struct {
	Result string `json:"result"`
	Data   string `json:"data"`
}

func EncodeRequest(req Request) ([]byte, error) {
	return json.Marshal(req)
}

func DecodeRequest(b []byte) (Request, error) {
	var req Request
	err := json.Unmarshal(b, &req)
	return req, err
}

func EncodeResponse(res Response) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeResponse(b []byte) (Response, error) {
	var res Response
	err := json.Unmarshal(b, &res)
	return res, err
}

