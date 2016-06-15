package main

type wordStats struct {
	word  string
	count int
}

// ByCount provides implementation hooks to sort wordStats by word count.
type ByCount []wordStats

// Len is part of the sort interface.
func (ws ByCount) Len() int {
	return len(ws)
}

// Swap is part of the sort interface.
func (ws ByCount) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

// Less is part of the sort interface.
func (ws ByCount) Less(i, j int) bool {
	// sort by count, then by word
	if ws[i].count != ws[j].count {
		return ws[i].count < ws[j].count
	}

	return ws[i].word < ws[j].word
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
