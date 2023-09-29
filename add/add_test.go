package add

import "testing"

func TestAddPostive(t *testing.T){
	if Add(10,20) != 30 {
		t.Error("Harusnya 10 + 20 itu adalah 30")
	}
}
