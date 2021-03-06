package treex

type ITree interface {
	Add(interface{}, interface{})
	Remove(interface{}) interface{}
	Contains(interface{}) bool
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	Size() int
	IsEmpty() bool
}
