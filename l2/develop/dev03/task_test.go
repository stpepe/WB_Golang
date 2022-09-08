package main

import (
	"testing"
)

func TestUnpack(t *testing.T){
	t.Run("positive test", func(t *testing.T){
		var inp, outp string = "file1.txt", "addddff"
		result, err := unpack(inp)
		if err != nil{
		t.Errorf("Expected %s, but result %s", outp, result)
		}
	})

	t.Run("negative test", func(t *testing.T){
		var inp, outp string = "afg44", "Error"
		result, err := unpack(inp)
		if err == nil{
			t.Errorf("Expected %s, but result %s", outp, result)
		}
	})
}
