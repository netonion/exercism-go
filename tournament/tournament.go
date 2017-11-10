package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type score struct {
	win  int
	draw int
	loss int
}

func (s score) matches() int {
	return s.win + s.draw + s.loss
}

func (s score) points() int {
	return s.win*3 + s.draw
}

func Tally(r io.Reader, w io.Writer) error {
	// Build tally
	buf := bufio.NewScanner(r)
	scores := make(map[string]*score)
	for buf.Scan() {
		line := buf.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		p := strings.Split(line, ";")
		if len(p) != 3 {
			return errors.New("Wrong number of fields.")
		}
		t1, t2, res := p[0], p[1], p[2]
		if scores[t1] == nil {
			scores[t1] = &score{}
		}
		if scores[t2] == nil {
			scores[t2] = &score{}
		}
		switch res {
		case "win":
			scores[t1].win++
			scores[t2].loss++
		case "draw":
			scores[t1].draw++
			scores[t2].draw++
		case "loss":
			scores[t1].loss++
			scores[t2].win++
		default:
			return errors.New("Match result incorrect.")
		}
	}

	// Sort teams based on score
	var teams []string
	for team := range scores {
		teams = append(teams, team)
	}
	sort.Slice(teams, func(i, j int) bool {
		p1, p2 := scores[teams[i]].points(), scores[teams[j]].points()
		return p1 > p2 || (p1 == p2 && teams[i] < teams[j])
	})

	// Output result
	fmt.Fprintf(w, "Team                           | MP |  W |  D |  L |  P\n")
	for _, team := range teams {
		s := scores[team]
		fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			team, s.matches(), s.win, s.draw, s.loss, s.points())
	}
	return nil
}
