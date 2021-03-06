package core

import "strings"

var resistances = []Resistance{
	{Name: "Kenobi", Location: []float64{-500, -200}},
	{Name: "Skywalker", Location: []float64{100, -100}},
	{Name: "Sato", Location: []float64{500, 100}},
}

// input: los satélites que se recibieron en el mensaje
// output: las distancias al emisor de cada satélite y los mensajes recibidos
func GetData(satellities []Satelite) (distances []float64, messages [][]string) {
	for _, satelite := range satellities {
		distances = append(distances, satelite.Distance)
		messages = append(messages, satelite.Message)
	}

	return
}

// input: distancia al emisor tal cual se recibe en cada satélite
// output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
func GetLocation(distances ...float64) (x, y float64) {
	var W, Z, y2 float64
	a := resistances[0].Location
	b := resistances[1].Location
	c := resistances[2].Location

	dA := distances[0]
	dB := distances[1]
	dC := distances[2]

	W = dA*dA - dB*dB - a[0]*a[0] - a[1]*a[1] + b[0]*b[0] + b[1]*b[1]
	Z = dB*dB - dC*dC - b[0]*b[0] - b[1]*b[1] + c[0]*c[0] + c[1]*c[1]

	x = (W*(c[1]-b[1]) - Z*(b[1]-a[1])) / (2 * ((b[0]-a[0])*(c[1]-b[1]) - (c[0]-b[0])*(b[1]-a[1])))
	y = (W - 2*x*(b[0]-a[0])) / (2 * (b[1] - a[1]))
	// y2 es una segunda medida de y para mitigar errores
	y2 = (Z - 2*x*(c[0]-b[0])) / (2 * (c[1] - b[1]))

	y = (y + y2) / 2

	return
}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
func GetMessage(messages ...[]string) (message string) {
	union := [][]string{}
	union = append(union, messages...)
	message = DecryptMessage(union)
	return
}

// input: la matriz de los mensajes tal cual son recibidos en cada satélite
// output: el mensaje desifrado
func DecryptMessage(messages [][]string) (message string) {
	inResult := make(map[string]bool)
	var words []string
	var value string

	var index int = 0
	var max int = 0

	for indexMessage, message := range messages {
		if len(message) > max {
			max = len(message)
			index = indexMessage
		}
	}

	for i, text := range messages[index] {
		for j, message := range messages {
			value = text
			if j != index {
				for k, word := range message {
					if i == k {
						if value == "" && word != "" {
							value = word
						}

						if _, ok := inResult[value]; !ok && value != "" {
							inResult[value] = true

							words = append(words, value)

						}
					}
				}
			}
		}
	}

	message = strings.Join(words[:], " ")

	return
}
