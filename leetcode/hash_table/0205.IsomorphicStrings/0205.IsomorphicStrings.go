package hash

func isIsomorphic(s string, t string) bool {
	l := len(s)
	mapS := [128]byte{}
	mapT := [128]byte{}

	for i := 0; i < l; i++ {
		if mapS[s[i]] != mapT[t[i]] {
			return false
		}
		mapS[s[i]] = byte(i + 1)
		mapT[t[i]] = byte(i + 1)
	}

	return true
}
