package remove

// Remove removes character or sequence(seq) from the string s
// If the given character or sequence does not exist it returns false
// Otherwise return a true
// o specifies the no. of occurences to remove
// if o>the actual occurence it removes mox no. of occurences
// Occurences will be removed from left to right
func Remove(s, seq string, o int) string {

	if seq == "" {
		return s
	}

	//oRem is occurences removed
	//pChar is processed characters
	//rBuf is result buffer
	//iBuf buffer reulsting from s
	//[]byte used over string for performance optimization
	var rBuf []byte
	var oRem, pCh int
	iBuf := []byte(s)
	sLen, sqLen := len(s), len(seq)
	for i := 0; i < sLen; i++ {
		if isAllOccurRemoved(o, oRem) {
			return string(rBuf) + string(iBuf[i:])
		}

		if isPotentialSeqMatch(sLen, sqLen, pCh, iBuf[i], seq[0]) {
			if isSeqMatch(iBuf[i:i+sqLen], seq) {
				oRem++
				i = i + sqLen - 1
				pCh += sqLen

				continue
			}
		}
		rBuf = append(rBuf, iBuf[i])
		pCh++
	}

	return string(rBuf)
}

func isPotentialSeqMatch(sLen, sqLen, pCh int, currIBufCh, firstSeqCh byte) bool {
	if (sLen-pCh > sqLen) && currIBufCh == firstSeqCh {
		return true
	}
	return false
}

func isSeqMatch(slice []byte, seq string) bool {
	if string(slice) == seq {
		return true
	}
	return false
}

func isAllOccurRemoved(o, oRem int) bool {
	if oRem == o {
		return true
	}
	return false
}
