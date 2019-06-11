package gofsm

import (
  "errors"
)

type Transition struct {
  Name string
  From string
  To string
}

type Model struct {
  States []string
  Transitions []Transition
}

func NewModel(states []string, transitions []Transition) (Model, error) {

  var from, to bool

  for _, t := range transitions {

    from, to = false, false

    for _, s := range states {
      if t.From == s {
          from = true
          break
      }
    }

    for _, s := range states {
      if t.To == s {
          to = true
          break
      }
    }

    if !from || !to {
      return Model{}, errors.New("Invalid model")
    }
  }

  return Model{States:states, Transitions:transitions}, nil
}

func (m *Model) CanTransition(from, to string) bool {
  for _, t := range m.Transitions {
    if t.From == from && t.To == to {
      return true
    }
  }
  return false
}
