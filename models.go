package main

// Resistance struct.
type Resistance struct {
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
}

// Transmitter struct.
type Satelite struct {
	Name string `json:"name"`
	Transmitter
}

// Transmitter struct.
type Transmitter struct {
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

// Top Secret struct.
type TopSecret struct {
	Satellites []Satelite `json:"satellites"`
}
