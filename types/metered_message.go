package types

import (
	"math/big"

	cbor "gx/ipfs/QmRoARq3nkUb13HSKZGepCZSWe5GrVPwx7xURJGZ7KWv9V/go-ipld-cbor"
)

// GasUnits represents number of units of gas consumed
type GasUnits = Uint64

// BlockGasLimit is the maximum amount of gas that can be used to execute messages in a single block
var BlockGasLimit = NewGasUnits(10000000)

func init() {
	cbor.RegisterCborType(MeteredMessage{})
}

// MeteredMessage contains a message and its associated gas price and gas limit
type MeteredMessage struct {
	Message  `json:"message"`
	GasPrice AttoFIL  `json:"gasPrice"`
	GasLimit GasUnits `json:"gasLimit"`
}

// Unmarshal a message from the given bytes.
func (msg *MeteredMessage) Unmarshal(b []byte) error {
	return cbor.DecodeInto(b, msg)
}

// Marshal the message into bytes.
func (msg *MeteredMessage) Marshal() ([]byte, error) {
	return cbor.DumpObject(msg)
}

// NewMeteredMessage accepts a message `msg`, a gas price `gasPrice` and a `gasLimit`.
// It returns a network message with the message, gas price and gas limit included.
func NewMeteredMessage(msg Message, gasPrice AttoFIL, gasLimit GasUnits) *MeteredMessage {
	return &MeteredMessage{
		Message:  msg,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}
}

// NewGasPrice constructs a gas price (in AttoFIL) from the given number.
func NewGasPrice(price int64) AttoFIL {
	return *NewAttoFIL(big.NewInt(price))
}

// NewGasUnits constructs a new GasUnits from the given number.
func NewGasUnits(cost uint64) GasUnits {
	return Uint64(cost)
}
