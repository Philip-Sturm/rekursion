package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	//Fall 0: s ist leer
	if s == "" {
		return false
	}
	if seq == "" {
		return true
	}
	//Fall 1: s[0] == seq
	if s[0] == seq[0] {
		return true
	}
	//Fall 2: s[0] != seq
	return Contains(s[1:], seq)
}

func longseq()
