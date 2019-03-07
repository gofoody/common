package generator

// SeqID return closure which generate unique id in seq 1,2,3 likewise.
func SeqID() func() int {
	nextID := 0
	return func() int {
		nextID++
		return nextID
	}
}

// FactorID return 100 times index plus id
func FactorID(index, id int) int {
	return (index * 100) + id
}
