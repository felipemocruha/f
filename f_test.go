package f

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Run("should return the sum of all numbers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		expected := 15

		sum := func(acc, val int) (int, error) {
			return acc + val, nil
		}

		res, _ := Reduce(sum, 0, in)
		assert.Equal(t, expected, res)
	})

	t.Run("should return the count of elements", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		expected := len(in)

		count := func(acc, val int) (int, error) {
			return acc + 1, nil
		}

		res, _ := Reduce(count, 0, in)
		assert.Equal(t, expected, res)
	})
}

func TestMap(t *testing.T) {
	t.Run("should return all uppercase letters", func(t *testing.T) {
		in := []string{"g", "e", "n", "e", "r", "i", "c", "s"}
		expected := []string{"G", "E", "N", "E", "R", "I", "C", "S"}

		upper := func(val string) (string, error) {
			return strings.ToUpper(val), nil
		}

		res, _ := Map(upper, in)
		assert.Equal(t, expected, res)
	})

	t.Run("should return all numbers as strings", func(t *testing.T) {
		in := []int{1, 2, 6}
		expected := []string{"1", "2", "6"}

		str := func(val int) (string, error) {
			return strconv.Itoa(val), nil
		}

		res, _ := Map(str, in)
		assert.Equal(t, expected, res)
	})
}

func TestFilter(t *testing.T) {
	t.Run("should return only odd numbers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := []int{1, 3, 5, 7, 9}

		odd := func(val int) (bool, error) {
			return val%2 != 0, nil
		}

		res, _ := Filter(odd, in)
		assert.Equal(t, expected, res)
	})
}
