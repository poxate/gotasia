package gotasia

import "strconv"

type idTracker struct{ int }

// increments the id and returns it
// used to generate a unique id
func (id *idTracker) gen() int {
	id.int++
	return id.int
}

// increments the id and returns the value before incrementation
func (id *idTracker) post() int {
	d := id.int
	id.int++
	return d
}

func (t *idTracker) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(t.int)), nil
}
