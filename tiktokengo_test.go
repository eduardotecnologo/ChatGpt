package tiktokengo

import "testing"

func TestHelloToMyName(t *testing.T) {
	input := "Eduardo"
	expectedOutput := "Eduardo, hello from tiktokengo!"
	output := HelloToMyName(input)
	if output != expectedOutput {
		t.Errorf("Got %s expected %s", output, expectedOutput)
	}
}
