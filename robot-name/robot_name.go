package robotname

import (
	"fmt"
	"math/rand"
)

var used = make(map[string]bool)

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if len(r.name) == 0 {
		for {
			var b [2]byte
			b[0] = byte('A' + rand.Intn(26))
			b[1] = byte('A' + rand.Intn(26))
			r.name = fmt.Sprintf("%s%03d", b, rand.Intn(1000))

			if !used[r.name] {
				used[r.name] = true
				break
			}
		}
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}
