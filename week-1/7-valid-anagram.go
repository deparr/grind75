
func sorting(s, t string) bool {
	sb, tb := []byte(s), []byte(t)
	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
	sort.Slice(tb, func(i, j int) bool { return tb[i] < tb[j] })

	for i, j := 0, 0; i < len(sb); i, j = i+1, j+1 {
		if sb[i] != tb[j] {
			return false
		}
	}

	return true

}

func hashmap(s, t string) bool {
	smap := map[rune]int{}
	for _, c := range s {
		smap[c]++
	}
	tmap := map[rune]int{}
	for _, c := range t {
		tmap[c]++
	}

	for c, scount := range smap {
		tcount, prs := tmap[c]
		if !prs || scount != tcount {
			return false
		}
	}

	for c := range tmap {
		if _, prs := smap[c]; !prs {
			return false
		}
	}

	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return sorting(s, t)
	//return hashmap(s, t)
}

