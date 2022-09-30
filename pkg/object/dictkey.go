package object

type DictKey struct {
	Type  ObjectType // the type of the origional key in DictLiteral, one of String, Integer, or Boolean
	Value uint64
}

type Hashable interface {
	Hash() DictKey
}
