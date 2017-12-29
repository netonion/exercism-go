package house

var parts = []string{
	" the house that Jack built.",
	" the malt\nthat lay in",
	" the rat\nthat ate",
	" the cat\nthat killed",
	" the dog\nthat worried",
	" the cow with the crumpled horn\nthat tossed",
	" the maiden all forlorn\nthat milked",
	" the man all tattered and torn\nthat kissed",
	" the priest all shaven and shorn\nthat married",
	" the rooster that crowed in the morn\nthat woke",
	" the farmer sowing his corn\nthat kept",
	" the horse and the hound and the horn\nthat belonged to",
}

// Verse returns verse v of the song.
func Verse(v int) string {
	verse := "This is"
	for i := v - 1; i >= 0; i-- {
		verse += parts[i]
	}
	return verse
}

// Song returns the whole song.
func Song() (song string) {
	for i := range parts {
		song += Verse(i+1) + "\n\n"
	}
	return song[:len(song)-2]
}
