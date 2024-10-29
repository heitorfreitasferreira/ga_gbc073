package hybrid

import "testing"

func TestArticleCase(t *testing.T) {
	inst3x3, err := GetInstanceFromFile("./../benchmark/instances/ft06") //GetInstanceFromFile("./tst03")

	if err != nil {
		t.Fatal(err)
	}
	// articleCromossome := newCromossome(*inst3x3, []int{8, 0, 2, 5, 7, 1, 4, 6, 3}, 0.5)
	articleCromossome := newCromossome(*inst3x3, []int{1, 29, 16, 18, 4, 31, 27, 25, 11, 22, 20, 17, 28, 9, 26, 14, 2, 19, 13, 33, 6, 30, 7, 23, 8, 3, 24, 15, 32, 35, 21, 12, 10, 5, 0, 34}, 0.5)
	// articleCromossome := newCromossome(*inst3x3, []int{10, 6, 0, 27, 4, 3, 25, 32, 33, 20, 31, 34, 2, 15, 30, 26, 11, 7, 24, 29, 5, 13, 18, 1, 22, 12, 14, 17, 23, 28, 16, 21, 19, 35, 9, 8}, 0.5)

	expected := infoMatrix{
		{8, 0, 2, 5, 7, 1, 4, 6, 3}, // Operação geral
		{2, 0, 0, 1, 2, 0, 1, 2, 1}, // Job ID
		{2, 0, 2, 2, 1, 1, 1, 0, 0}, // Operação do job
		{1, 1, 0, 1, 0, 2, 2, 2, 0}, //  Máquina
		{9, 3, 5, 8, 3, 7, 9, 4, 6}, // Tempo do job na máquina
	}

	if len(expected) != len(articleCromossome.infoMatrix) {
		t.Fatalf("Expected %d, got %d", len(expected), len(articleCromossome.infoMatrix))
	}

	for i := range expected {
		if len(expected[i]) != len(articleCromossome.infoMatrix[i]) {
			t.Fatalf("Expected %d, got %d", len(expected[i]), len(articleCromossome.infoMatrix[i]))
		}
		for j := range expected[i] {
			if expected[i][j] != articleCromossome.infoMatrix[i][j] {
				t.Fatalf("Expected %d, got %d", expected[i][j], articleCromossome.infoMatrix[i][j])
			}
		}
	}
}
