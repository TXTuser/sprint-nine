package main

import (
	"testing"
)

func TestMaximum(t *testing.T) {
	// Тест 1: Пустой слайс
	t.Run("empty slice", func(t *testing.T) {
		result := maximum([]int{})
		expected := 0
		if result != expected {
			t.Errorf("maximum([]int{}) = %d; want %d", result, expected)
		}
	})

	// Тест 2: Слайс с одним элементом
	t.Run("single element", func(t *testing.T) {
		result := maximum([]int{42})
		expected := 42
		if result != expected {
			t.Errorf("maximum([]int{42}) = %d; want %d", result, expected)
		}
	})

	// Тест 3: Cлайс с положительными числами
	t.Run("positive numbers", func(t *testing.T) {
		result := maximum([]int{1, 5, 3, 9, 2})
		expected := 9
		if result != expected {
			t.Errorf("maximum([]int{1,5,3,9,2}) = %d; want %d", result, expected)
		}
	})

	// Тест 4: Слайс с отрицательными числами
	t.Run("negative numbers", func(t *testing.T) {
		result := maximum([]int{-5, -1, -7, -3})
		expected := -1
		if result != expected {
			t.Errorf("maximum([]int{-5,-1,-7,-3}) = %d; want %d", result, expected)
		}
	})

	// Тест 5: Слайс с одинаковыми числами
	t.Run("equal numbers", func(t *testing.T) {
		result := maximum([]int{7, 7, 7, 7})
		expected := 7
		if result != expected {
			t.Errorf("maximum([]int{7,7,7,7}) = %d; want %d", result, expected)
		}
	})

	// Тест 6: Слайс с смешанными числами (положительные + отрицательные)
	t.Run("mixed numbers", func(t *testing.T) {
		result := maximum([]int{-10, 0, 5, -3, 8})
		expected := 8
		if result != expected {
			t.Errorf("maximum([]int{-10,0,5,-3,8}) = %d; want %d", result, expected)
		}
	})
}
