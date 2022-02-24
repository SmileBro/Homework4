package point

type Point3d struct {
	x float64
	y float64
	z float64
}

func (p Point3d) Z() float64 {
	return p.z
}

func NewPoint3d(x float64, y float64, z float64) *Point3d {
	return &Point3d{x: x, y: y, z: z}
}

/*func (p Point3d) MarshalJSON() ([]byte, error) {
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

func (p Point3d) X() float64 {
	return p.x
}

func (p Point3d) Y() float64 {
	return p.y
}
