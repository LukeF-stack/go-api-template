package utils

import (
	"github.com/gofiber/fiber/v2"
)

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}

func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for index, _ := range slice {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[index] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[index] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[index] = left[i]
			i++
		} else {
			slice[index] = right[j]
			j++
		}
	}
	return slice
}
