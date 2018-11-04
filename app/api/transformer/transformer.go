package transformer

type Envelope struct {
	Data interface{} `json:"data"`
}

// envelope wraps data in a data JSON structure
func envelope(data interface{}) *Envelope {
	return &Envelope{data}
}
