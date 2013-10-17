package tree

type Tree interface {
	AllWithPrefix(prefix string, limit int) (values []interface{})
	Contains(key string) bool
	Delete(key string)
	Get(key string) (val interface{})
	Set(key string, val interface{})
}

func New() Tree {
	return new(node)
}
