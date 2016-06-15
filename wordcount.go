package main

type wordStats struct {
	word  string
	count int
}

func uniqueWords(words []string) map[string]wordStats {
	unique := map[string]wordStats{}

	for _, w := range words {
		stats, ok := unique[w]
		if !ok {
			stats.word = w
		}
		stats.count++
		unique[w] = stats
	}

	return unique
}

func values(m map[string]wordStats) []wordStats {
	allStats := make([]wordStats, 0, len(m))

	for _, stats := range m {
		allStats = append(allStats, stats)
	}

	return allStats
}
