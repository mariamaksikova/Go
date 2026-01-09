package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var result []BrainrotMeme
	for _, meme := range memes {
		if meme.Views > minViews {
			result = append(result, meme)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].TrendLevel > result[j].TrendLevel
	})
	return result
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)
	for _, meme := range memes {
		impact[meme.Category] += meme.Views
	}
	return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var names []string
	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			names = append(names, meme.Name)
		}
	}
	return names
}

func main() {
	memes := []BrainrotMeme{
		{"Meme1", 8, "Subo Bratik", 45.5},
		{"Meme2", 10, "TUNTUNTUNSAHUR", 120.0},
		{"Meme3", 6, "Sigma", 60.2},
		{"Meme4", 9, "Skibidi", 30.1},
		{"Meme5", 7, "Mewing", 55.0},
		{"Meme6", 5, "Other", 10.3},
		{"Meme7", 9, "Sigma", 70.4},
	}

	topTrending := FindTopTrending(memes, 50.0)
	fmt.Println("Top Trending Memes:")
	for _, m := range topTrending {
		fmt.Printf("%s (Trend: %d, Views: %.1f)\n", m.Name, m.TrendLevel, m.Views)
	}

	categoryImpact := CalculateCategoryImpact(memes)
	fmt.Println("\nCategory Impact:")
	for cat, views := range categoryImpact {
		fmt.Printf("%s: %.1f\n", cat, views)
	}

	filteredNames := FilterByComplexCondition(memes)
	fmt.Println("\nFiltered Memes:")
	for _, name := range filteredNames {
		fmt.Println(name)
	}
}
