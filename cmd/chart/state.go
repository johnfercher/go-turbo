package main

type StateType string

const (
	GetBorders StateType = "get_borders"
	GetScales  StateType = "get_scale_y"
	GetScaleX  StateType = "get_scale_x"
)

type State struct {
	stateType StateType
	next      *State
	previous  *State
}

func NewState(stateType StateType) *State {
	return &State{
		stateType: stateType,
	}
}

func (s *State) SetNext(next *State) {
	next.previous = s
	s.next = next
}

func (s *State) GetNext() *State {
	return s.next
}

func (s *State) GetPrevious() *State {
	return s.previous
}

func (s *State) GetType() StateType {
	return s.stateType
}

var getBorders = NewState("get_borders")
var getScaleY = NewState("get_scale_y")
var getScaleX = NewState("get_scale_x")

func GetStateMachine() *State {
	getBorders.SetNext(getScaleY)
	getScaleY.SetNext(getScaleX)
	return getBorders
}
