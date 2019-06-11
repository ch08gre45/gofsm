package gofsm

import (
  "testing"
)

var modelTests = []struct {
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
	for _, mt := range modelTests {
		t.Run(mt.label, func(t *testing.T) {
			_, err := NewModel(mt.in.States, mt.in.Transitions)
      got := err == nil
			if got != mt.expect {
				t.Errorf("Test %q => got [%v], want [%v]", mt.label, got, mt.expect)
			}
		})
	}
}

var testModel = Model{
  States: []string{"a","b","c"},
  Transitions: []Transition{
    Transition{"ab","a","b"},
    Transition{"bc","b","c"},
  },
}

var transitionTests = []struct {
  label string
	from  string
  to    string
	expect bool
}{
	{"OK", "a","b", true},
  {"Transition does not exist", "a","c", false},
  {"Initial state does not exist", "x","c", false},
  {"End state does not exist", "a","x", false},
}

func TestTransitions(t *testing.T) {
	for _, tt := range transitionTests {
		t.Run(tt.label, func(t *testing.T) {

      got := testModel.CanTransition(tt.from, tt.to)
			if got != tt.expect {
				t.Errorf("Test %q => got [%v], want [%v]", tt.label, got, tt.expect)
			}
		})
	}
}
