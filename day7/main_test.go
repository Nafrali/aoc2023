package main

import "testing"

func TestCardSorting(t *testing.T) {
	got := sortCards("95995")
	want := "99955"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	got = sortCards("9KAQT")
	want = "AKQT9"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestTieBreak(t *testing.T) {
	got := isAMoreThanB("99995", "99994")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = isAMoreThanB("AA894", "AA995")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestFiveOfAKind(t *testing.T) {
	got := isFiveOfAKind("KKKKA")
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = isFiveOfAKind("99999")
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestFourOfAKind(t *testing.T) {
	got := isFourOfAKind("KKKKA")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = isFourOfAKind("KKKAA")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestFullHouse(t *testing.T) {
	got := isFullHouse("KKKAA")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = isFullHouse("KK9AA")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsThreeOfAKind(t *testing.T) {
	got := isThreeOfAKind("999K7")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = isThreeOfAKind("KK9AA")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsTwoPairs(t *testing.T) {
	got := isTwoPairs("QQTAA")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = isThreeOfAKind("KKKAA")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsPair(t *testing.T) {
	got := isPair("QQT9A")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = isPair("KKKAA")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestFaceValueMapper(t *testing.T) {
	initiateFaceValueMapper()

	got := getFaceValueMap('A')
	var want uint16 = 0x8000

	if got != want {
		t.Errorf("got %08b, wanted %08b", got, want)
	}

	got = getFaceValueMap('2')
	want = 0b0000000000001000

	if got != want {
		t.Errorf("got %08b, wanted %08b", got, want)
	}
}

func TestGetDecimalValuesForCards(t *testing.T) {
	got := getDecimalValuesForCards("9999K")
	var want int64 = 4504630419521536

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
