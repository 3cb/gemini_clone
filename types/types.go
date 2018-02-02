package types

// Message defines the structure of messages received from Gemini websocket
type Message struct {
	Type      string  `json:"type"`
	Product   string  `json:"product"`
	EventID   int     `json:"eventId"`
	Sequence  int     `json:"socket_sequence"`
	Events    []Event `json:"events"`
	Timestamp int     `json:"timestampms"`
}

// Event defines the structure of Events field contained in Message struct
type Event struct {
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Price     string `json:"price"`
	Delta     string `json:"delta"`
	Remaining string `json:"remaining"`
	Side      string `json:"side"`

	TID       int64  `json:"tid"`
	Amount    string `json:"amount"`
	MakerSide string `json:"makerSide"`
}
