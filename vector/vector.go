package vector

type Vector struct {
	x float64
	y float64
	z float64
}

func (v *Vector) add(other Vector) *Vector {
  return &Vector {
    v.x + other.x,
    v.y + other.y,
    v.z + other.z,
  }
}
