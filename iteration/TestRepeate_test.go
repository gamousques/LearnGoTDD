package iteration

import "testing"

func TestRepeate(t *testing.T) {
	t.Run("Repeat five times ge aaaa", func(t *testing.T) {
		repeate := Repeate("a", 5)
		expected := "aaaaa"
	
		if repeate != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeate)
		}
	})

	t.Run("test negative interation resturns empty string", func (t *testing.T) {
		repeate := Repeate("a", -5)
		expected := ""

		if len(repeate) != len(expected) {
			t.Errorf("expected '%q' but got '%q'", expected, repeate)
		}
	})
}


func BenchmarkRepeate(t *testing.B) {
	t.Run("benchmark Repeate", func(t *testing.B) {
		for i := 0; i < t.N; i++ {
			Repeate("a", 5)
		}

	})
}


