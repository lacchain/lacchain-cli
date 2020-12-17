package model

//Event ...
type Event struct {
    NodeAdded bool
    EnodeHigh [32]byte
    EnodeLow [32]byte
    EnodeIp [16]byte
    EnodePort uint16
}