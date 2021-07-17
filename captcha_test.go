package main

import (
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	storage := NewStorage()
	storage.CleanUp()

	itemFa := storage.NewItem(LevelEasy, "fa", 3)
	itemAr := storage.NewItem(LevelMedium, "ar", 3)
	itemEn := storage.NewItem(LevelHard, "en", 3)

	GenerateCaptcha(itemFa)
	GenerateCaptcha(itemAr)
	GenerateCaptcha(itemEn)
}

func BenchmarkGenerateCaptcha(b *testing.B) {
	storage := NewStorage()
	for i := 0; i < b.N; i++ {
		GenerateCaptcha(storage.NewItem(LevelHard, "fa", 3))
	}
}
