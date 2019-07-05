package arraylist

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

type List struct {
	elements []interface{}
	size     int
}
