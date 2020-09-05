package idgen

type IDGenerator interface {
	NextID() (int64, error)
	GetID() int64
}
