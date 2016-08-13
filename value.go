package attr

// KeyValuer return valuer from key
type KeyValuer interface {
	KeyValue(Keyer) interface{}
}
