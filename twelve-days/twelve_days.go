package twelve

import (
	"fmt"
	"strings"
)

var nth = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var obj = []string{"twelve Drummers Drumming,", "eleven Pipers Piping,", "ten Lords-a-Leaping,", "nine Ladies Dancing,", "eight Maids-a-Milking,", "seven Swans-a-Swimming,", "six Geese-a-Laying,", "five Gold Rings,", "four Calling Birds,", "three French Hens,", "two Turtle Doves, and", "a Partridge in a Pear Tree."}

func Verse(n int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s", nth[n-1], strings.Join(obj[12-n:], " "))
}

func Song() (song string) {
	for n := 1; n < 13; n++ {
		song += Verse(n) + "\n"
	}
	return
}
