package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	Name          string
	MatchesPlayed int
	Wins          int
	Draws         int
	Losses        int
	Points        int
}

type Match struct {
	Host   string
	Guest  string
	Result string
}

type lessFunc func(p1, p2 *Team) bool
type multiSorter struct {
	teams []Team
	less  []lessFunc
}

func (ms *multiSorter) Sort(teams []Team) {
	ms.teams = teams
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.teams)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.teams[i], ms.teams[j] = ms.teams[j], ms.teams[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.teams[i], &ms.teams[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func LinesToArray(reader io.Reader) ([]string, error) {
	// lines in slices
	var lines []string

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func MatchesToTeams(matches []Match) (map[string]Team, error) {
	var teams = make(map[string]Team)

	for _, m := range matches {
		var host, guest Team

		host, host_exists := teams[m.Host]
		if !host_exists {
			host = Team{Name: m.Host}
		}

		guest, guest_exists := teams[m.Guest]
		if !guest_exists {
			guest = Team{Name: m.Guest}
		}

		host.MatchesPlayed++
		guest.MatchesPlayed++
		if m.Result == "win" {
			host.Wins++
			host.Points += 3
			guest.Losses++
		} else if m.Result == "loss" {
			host.Losses++
			guest.Wins++
			guest.Points += 3
		} else if m.Result == "draw" {
			// a draw, hopefully
			host.Draws++
			host.Points += 1
			guest.Draws++
			guest.Points += 1
		} else {
			// invalid line
			return teams, errors.New("Invalid match")
		}

		teams[m.Host] = host
		teams[m.Guest] = guest
	}
	return teams, nil
}

func ArrayToMatches(lines []string) ([]Match, error) {
	var matches []Match
	for _, line := range lines {
		match := strings.Split(line, ";")
		if len(match) == 2 {
			return nil, errors.New("Invalid match")
		} else if len(match) == 3 {
			host := match[0]
			guest := match[1]
			result := match[2]
			matches = append(matches, Match{host, guest, result})
		}
	}
	return matches, nil
}

func PrettyPrint(teams []Team) string {
	// titles
	table := fmt.Sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range teams {
		table += fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", v.Name, v.MatchesPlayed, v.Wins, v.Draws, v.Losses, v.Points)
	}
	return table
}

func Tally(reader io.Reader, writer io.Writer) error {
	lines, err := LinesToArray(reader)
	if err != nil {
		return err
	}

	matches, err := ArrayToMatches(lines)
	if err != nil {
		return err
	}
	teams, err := MatchesToTeams(matches)
	if err != nil {
		return err
	}

	points := func(t1, t2 *Team) bool {
		return t1.Points > t2.Points
	}

	names := func(t1, t2 *Team) bool {
		return t1.Name < t2.Name
	}

	var t []Team

	for _, v := range teams {
		t = append(t, v)
	}

	OrderedBy(points, names).Sort(t)
	io.WriteString(writer, PrettyPrint(t))
	return nil
}
