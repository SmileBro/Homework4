package point

type Point struct {
	x float64
	y float64
}

/*func (p Point) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(
		struct {
			X         float64 `json:"x"`
			Y         float64 `json:"y"`
			CreatedAt string  `json:"createdAt,omitempty"`
		}{p.x, p.y, ""})

	if err != nil {
		return nil, err
	}

	return j, err
}*/

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}
