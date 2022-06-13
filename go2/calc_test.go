package calc_test

func TestAdd(t *testing.T)  {
	got := calc.(2, 1)
	expected := 3

	if got != expected {
		t.Errorf("not expected result: %v expected: %v", got, expected)
	}
}