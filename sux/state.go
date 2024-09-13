package sux

// State
func NewState() *State {
	return &State{Defined: true}
}

func (s *State) IsDefined() bool {
	return s.Defined
}
