package progressbar

import "testing"

func TestPbarSizeFiveHundredThousand(t *testing.T) {
	pb, err := NewPbar(50000000)
	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < 50000000; i++ {
		pb.IncreaseBar()
	}
}

func TestPbarFromArray(t *testing.T) {
	arr := make([]int, 1000000)
	pb, err := NewPbar(arr)
	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < len(arr); i++ {
		pb.IncreaseBar()
	}
}

func TestPbarCustomWidth(t *testing.T) {
	arr := make([]int, 4352120)
	pb, err := NewPbar(arr, 5)
	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < len(arr); i++ {
		pb.IncreaseBar()
	}
}

func TestPbarAesthetics(t *testing.T) {
	pb, err := NewPbar(400000000, 50)
	pb.SetGraphics("\u001b[1mX", " ")
	pb.ToggleColor([]string{"cyan", "magenta", "green"})
	pb.SetPrintNumbers("%")
	pb.SetIncrementAmt(2)

	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < 400000000; i += 2 {
		pb.IncreaseBar()
	}
}

func TestPbarFailure(t *testing.T) {
	_, err := NewPbar("poop")
	if err == nil {
		t.Errorf("How did this happen")
	}

	_, err = NewPbar(100, "ok")
	if err == nil {
		t.Errorf("How did this happen")
	}

}
