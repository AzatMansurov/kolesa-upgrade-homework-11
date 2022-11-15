package main

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/require"
)

// Задание
// 1. Дописать существующие тесты.
// 2. Придумать один новый тест.
//
// Дополнительное задание: использовать табличные тесты и testify.require
//
// Документация по тестируемой функции:
// Делает слова в строке из первого аргумента с большой буквы, кроме слов из строки из второго аргумента
// Первое слово всегда с большой буквы
//
// TitleCase(str, minor) returns a str string with all words capitalized except minor words.
// The first word is always capitalized.
//
// E.g.
// TitleCase("the quick fox in the bag", "") = "The Quick Fox In The Bag"
// TitleCase("the quick fox in the bag", "in the") = "The Quick Fox in the Bag"

type testCase struct {
	str   string
	minor string
	want  string
}

func TestEmpty(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%v, %v) = %v; want %v", str, minor, got, want)
	}
}

func TestWithoutMinor(t *testing.T) {
	cases := []testCase{
		{
			str:   "the quick fox in the bag",
			minor: "",
			want:  "The Quick Fox In The Bag",
		},
		{
			str:   "kompot",
			minor: "",
			want:  "Kompot",
		},
	}
	for _, value := range cases {
		got := titlecase.TitleCase(value.str, value.minor)
		require.Equal(t, value.want, got)
	}
}

func TestWithMinorInFirst(t *testing.T) {
	cases := []testCase{
		{
			str:   "the quick fox in the bag",
			minor: "",
			want:  "The Quick Fox In the Bag",
		},
		{
			str:   "ilon the musk",
			minor: "",
			want:  "Ilon the Musk",
		},
	}
	for _, value := range cases {
		got := titlecase.TitleCase(value.str, value.minor)
		require.Equal(t, value.want, got)
	}
}