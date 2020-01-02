package model

type PeerResponse struct {
	Jsonrpc string `json:"jsonrpc"`
    Id int `json:"id"`
	Result string `json:"result"`
}