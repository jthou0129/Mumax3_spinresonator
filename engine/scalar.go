package engine

type scalar struct {
	nComp     int
	name      string
	value     []float64
	timestamp int
	armed     bool
	updateFn  func() []float64
	UpdCount  int
	// todo: deps: interface{arm}
}

func newScalar(nComp int, name string, updateFn func() []float64) *scalar {
	s := new(scalar)
	s.nComp = nComp
	s.name = name
	s.updateFn = updateFn
	return s
}

func (s *scalar) Get() []float64 {
	if s.timestamp != itime {
		s.value = s.updateFn()
		s.UpdCount++
		s.timestamp = itime
	}
	return s.value
}

func (s *scalar) touch(good bool) {
	if s.armed {
		_ = s.Get() // update s.value
		s.armed = false
	}
}

// when armed, updateFn will fire upon next touch
func (s *scalar) arm() {
	s.armed = true
}