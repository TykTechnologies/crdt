package crdt

import (
	"testing"
	"encoding/json"
)

func TestGCounter(t *testing.T) {
	for _, tt := range []struct {
		incsOne int
		incsTwo int
		result  int
	}{
		{5, 10, 15},
		{10, 5, 15},
		{100, 100, 200},
		{1, 2, 3},
	} {
		gOne, gTwo := NewGCounter(), NewGCounter()

		for i := 0; i < tt.incsOne; i++ {
			gOne.Inc()
		}

		for i := 0; i < tt.incsTwo; i++ {
			gTwo.Inc()
		}

		gOne.Merge(gTwo)

		if gOne.Count() != tt.result {
			t.Errorf("expected total count to be: %d, actual: %d",
				tt.result,
				gOne.Count())
		}

		gTwo.Merge(gOne)

		if gTwo.Count() != tt.result {
			t.Errorf("expected total count to be: %d, actual: %d",
				tt.result,
				gTwo.Count())
		}
	}
}

func TestGCounterInvalidInput(t *testing.T) {
	gc := NewGCounter()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("panic expected here")
		}
	}()

	gc.IncVal(-5)
}

func TestGCounterMarshaller(t *testing.T) {
	gc := NewGCounter()
	gc.Inc()
	gc.Inc()

	var asJson []byte
	var err error
	asJson, err = json.Marshal(gc); if err != nil {
		t.Fatal(err)
	}

	var gc2 GCounter
	if err := json.Unmarshal(asJson, &gc2); err != nil {
		t.Fatal(err)
	}

	gc2.Inc()
	if gc2.Count() != 3 {
		t.Fatalf("Counter should be 3! Is: %v\n", gc2.Count())
	}
}