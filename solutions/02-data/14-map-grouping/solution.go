//go:build ignore

package koan

type Member struct{ Name, Team string }

func Group(members []Member) map[string][]string {
	groups := make(map[string][]string)
	for _, m := range members {
		groups[m.Team] = append(groups[m.Team], m.Name)
	}
	return groups
}
