package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) (freq FreqMap) {
	c := make(chan FreqMap)
	for _, text := range texts {
		go func(text string) {
			c <- Frequency(text)
		}(text)
	}

	freq = make(FreqMap)
	for range texts {
		for k, v := range <-c {
			freq[k] += v
		}
	}
	return
}
