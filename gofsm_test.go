package fsm

import (
  "testing"
)

var modeltest = []struct {
  label string
	in  Model
	expect bool
}{
	{
    "OK",
    Model{
      States: []string{"a","b"},
      Transitions: []Transition{
        Transition{"ab","a","b"},
        Transition{"ba","a","b"},
      },
    },
    true,
  },
  {
    "Invalid from",
    Model{
      States: []string{"a","b"},
      Transitions: []Transition{
        Transition{"ab","c","b"},
      },
    },
    false,
  },
  {
    "Invalid to",
    Model{
      States: []string{"a","b"},
      Transitions: []Transition{
        Transition{"ab","a","c"},
      },
    },
    false,
  },
}

func TestNewModel(t *testing.T) {
	for _, mt := range modeltest {
		t.Run(mt.label, func(t *testing.T) {
			_, err := NewModel(mt.in.States, mt.in.Transitions)
      got := err == nil
			if got != mt.expect {
				t.Errorf("Test %q => got [%v], want [%v]", mt.label, got, mt.expect)
			}
		})
	}
}
