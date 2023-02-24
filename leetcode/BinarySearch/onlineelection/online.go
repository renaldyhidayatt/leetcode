package onlineelection

import "sort"

type TopVotedCandidate struct {
	persons []int
	times   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	leaders, votes := make([]int, len(persons)), make([]int, len(persons))
	leader := persons[0]

	for i := 0; i < len(persons); i++ {
		p := persons[i]
		votes[p]++
		if votes[p] >= votes[leader] {
			leader = p
		}
		leaders[i] = leader
	}

	return TopVotedCandidate{persons: leaders, times: times}
}

func (tvc *TopVotedCandidate) Q(t int) int {
	i := sort.Search(len(tvc.times), func(p int) bool { return tvc.times[p] > t })
	return tvc.persons[i-1]
}
