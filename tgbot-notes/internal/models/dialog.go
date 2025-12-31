package models

type Dialog struct {
	states map[string]bool
}

func NewDialog() *Dialog {
	return &Dialog{
		states: make(map[string]bool),
	}
}

func (d *Dialog) SetState(state string) {
	d.states[state] = true
}

func (d *Dialog) GetState(state string) bool {
	return d.states[state]
}

func (d *Dialog) DeleteState(state string) {
	delete(d.states, state)
}

func (d *Dialog) Length() int {
	return len(d.states)
}
