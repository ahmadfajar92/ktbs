package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gcd(t *testing.T) {
	t.Run("FIND GCD", func(t *testing.T) {
		assert.Equal(t, 3, findGCD(21, 9))
	})
}

func Test_scenario(t *testing.T) {

	t.Run("there's no item can divided evenly", func(t *testing.T) {
		// there's no item can divided evenly
		box := NewBox(9, 25)
		boxes := box.HowMany()
		apple, cake := box.CountItemsEachBox()

		assert.Equal(t, 1, boxes)
		assert.Equal(t, 9, apple)
		assert.Equal(t, 25, cake)
	})

	t.Run("item can divided evenly", func(t *testing.T) {
		// there's no item can divided evenly
		box := NewBox(30, 25)
		boxes := box.HowMany()
		apple, cake := box.CountItemsEachBox()

		assert.Equal(t, 5, boxes)
		assert.Equal(t, 6, apple)
		assert.Equal(t, 5, cake)
	})
}
