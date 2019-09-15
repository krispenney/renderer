package main

type Vector struct {
	x, y, z float64
}

func (v *Vector) add(other Vector) Vector {
  return Vector {
    v.x + other.x,
    v.y + other.y,
    v.z + other.z,
  }
}

func (v *Vector) scale(scalar float64) Vector {
  return Vector {
    v.x * scalar,
    v.y * scalar,
    v.z * scalar,
  }
}

func (v *Vector) subtract(other Vector) Vector {
  inverted := other.scale(-1)

  return v.add(inverted)
}
