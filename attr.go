package attr

// Attributes is basic attr map
type Attributes map[Keyer]interface{}

// Keys return all of key list
func (a *Attributes) Keys() (list []Keyer) {
	for k := range *a {
		list = append(list, k)
	}
	return
}
