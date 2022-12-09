package point

type Point2D struct {
	X int
	Y int
}

func (p *Point2D) InReach(target *Point2D, length int) bool {
	for dx := -length; dx <= length; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if p.X+dx == target.X && p.Y+dy == target.Y {
				return true
			}
		}
	}
	return false
}
