package hybrid

import "testing"

func TestCrossoverCorrecting(t *testing.T) {
	instance, err := GetInstanceFromFile("./tst03")
	if err != nil {
		t.Fatal(err)
	}
	//2 6 4 7 3 5 8 9 1
	//4 5 2 1 8 7 6 9 3
	p1 := newCromossome(*instance, []int{1, 5, 3, 6, 2, 4, 7, 8, 0}, 0.5)
	p2 := newCromossome(*instance, []int{3, 4, 1, 0, 7, 6, 5, 8, 2}, 0.5)

	off1, off2 := crossover(p1, p2, 4, *instance, 0.5)

	// 2 5 4 1 8 7 6 9 3
	// 4 6 2 7 3 5 8 9 1
	// expectedOff1 := []int{1, 4, 3, 0, 7, 6, 5, 8, 2}
	// expectedOff2 := []int{3, 5, 1, 6, 2, 4, 7, 8, 0}

	count1 := make(map[int]int)
	count2 := make(map[int]int)
	for i := 0; i < len(off1.infoMatrix[0]); i++ {
		count1[off1.infoMatrix[0][i]]++
		count2[off2.infoMatrix[0][i]]++
	}
	for i := 0; i < len(p1.infoMatrix[0]); i++ {
		if count1[i] > 1 {
			t.Errorf("off1 has duplicate %d", i)
		}
		if count2[i] > 1 {
			t.Errorf("off2 has duplicate %d", i)
		}
	}
}
