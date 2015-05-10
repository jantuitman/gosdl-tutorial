package basic;

type Rect struct {
  X int
  Y int
  W int 
  H int
}

// returns self, with offset determined by other.x,other.y
func (self Rect) offset(other Rect) Rect {
  return Rect{other.X + self.X , other.Y + self.Y , self.W , self.H }
}

