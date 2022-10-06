package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fulfillWordInformation(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		exampleRow := []string{"Achievement", "Başarı", "you should see my achievement"}
		dto := fulfillWordInformation(exampleRow)

		expectedDTO := EnglishWordDTO{
			Word:            "Achievement",
			Meaning:         "Başarı",
			ExampleSentence: "you should see my achievement",
		}

		assert.Equal(t, expectedDTO, dto)
	})

	t.Run("Failed Case", func(t *testing.T) {
		exampleRow := []string{"Achievement", "Başarı", ""}
		dto := fulfillWordInformation(exampleRow)

		expectedDTO := EnglishWordDTO{
			Word:            "Achievement",
			Meaning:         "Başarı",
			ExampleSentence: " ",
		}

		assert.NotEqual(t, expectedDTO, dto)
	})
}
