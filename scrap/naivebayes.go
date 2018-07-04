package scrap

import "fmt"
import (
	"github.com/lytics/multibayes"
)

func Demo() {
	documents := []struct {
		Text    string
		Classes []string
	}{
		{
			Text:    `guru smp`,
			Classes: []string{"guru", "smp"},
		}, {
			Text:    `guru sd`,
			Classes: []string{"guru", "sd"},
		}, {
			Text:    `nurid`,
			Classes: []string{"guru"},
		}, {
			Text:    `dibutuhkan`,
			Classes: []string{"guru", "spg"},
		},
		{
			Text:    `rokok`,
			Classes: []string{"spg"},
		},
		{
			Text:    `Spg`,
			Classes: []string{"spg"},
		},
		{
			Text:    `event`,
			Classes: []string{"spg"},
		},
	}

	classifier := multibayes.NewClassifier()
	classifier.MinClassSize = 0

	// train the classifier
	for _, document := range documents {
		classifier.Add(document.Text, document.Classes)
	}

	// predict new classes
	probs := classifier.Posterior("dibutuhkan guru smp")
	fmt.Printf("Posterior Probabilities: %+v\n", probs)
}

func Demo2(txt string) {

	documents := []struct {
		Text    string
		Classes []string
	}{
		// {
		// 	Text:    `guru smp`,
		// 	Classes: []string{"guru", "smp"},
		// }, {
		// 	Text:    `guru sd`,
		// 	Classes: []string{"guru", "sd"},
		// }, {
		// 	Text:    `nurid`,
		// 	Classes: []string{"guru"},
		// }, {
		// 	Text:    `dibutuhkan`,
		// 	Classes: []string{"guru", "spg"},
		// },
		// {
		// 	Text:    `rokok`,
		// 	Classes: []string{"spg"},
		// },
		// {
		// 	Text:    `Spg`,
		// 	Classes: []string{"spg"},
		// },
		// {
		// 	Text:    `event`,
		// 	Classes: []string{"spg"},
		// }, {
		// 	Text:    `jakarta`,
		// 	Classes: []string{"guru"},
		// },
	}

	classifier := multibayes.NewClassifier()
	classifier.MinClassSize = 0

	// train the classifier
	for _, document := range documents {
		classifier.Add(document.Text, document.Classes)
	}

	// predict new classes
	probs := classifier.Posterior(txt)
	fmt.Printf("Posterior Probabilities: %+v\n", probs)
}
