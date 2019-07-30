package progressbar

import "testing"

func TestPbarSizeFiveHundredThousand(t *testing.T) {
	pb, err := NewPbar(500000)
	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < 500000; i++ {
		pb.increaseBar()
	}
}

func TestPbarSizeOneHundredThousand(t *testing.T) {
	arr := make([]int, 100000)
	pb, err := NewPbar(arr)
	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < len(arr); i++ {
		pb.increaseBar()
	}
}

func TestPbarCustomWidth(t *testing.T) {
	arr := make([]int, 500000)
	pb, err := NewPbar(arr, 50)
	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < len(arr); i++ {
		pb.increaseBar()
	}
}

func TestPbarAesthetics(t *testing.T) {
	pb, err := NewPbar(200000, 50, true)
	pb.setFillString("\u001b[1mX")

	if err != nil {
		t.Errorf("How did this happen")
	}
	for i := 0; i < 200000; i++ {
		pb.increaseBar()
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

	_, err = NewPbar(100, 50, "0")
	if err == nil {
		t.Errorf("How did this happen")
	}

}