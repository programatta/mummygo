package states

//StateMgr ...
type StateMgr struct {
	states         map[string]*IState
	currentStateID string
}

//NewStateMgr es el constructor.
func NewStateMgr() *StateMgr {
	sm := &StateMgr{}
	sm.states = make(map[string]*IState)

	return sm
}

//AddState añade estados a la maquina de estados.
func (sm *StateMgr) AddState(name string, state IState) {
	sm.states[name] = &state
}

//GetState ...
func (sm *StateMgr) GetState(name string) IState {
	sm.currentStateID = name
	return *sm.states[name]
}

//ChangeState ...
func (sm *StateMgr) ChangeState(currentState IState) IState {
	stateID := currentState.NextState()
	state := *sm.states[stateID]
	if sm.currentStateID != stateID {
		sm.currentStateID = stateID
		state = *sm.states[stateID]
		state.Init()
	}
	return state
}
