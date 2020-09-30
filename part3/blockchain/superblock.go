package blockchain

// FindPrevLevelBlockIndex will take in a block index and
// RETURN the block index of the closest higher level block
func FindPrevLevelBlockIndex(lc []BlockHeader, n int) int {
	if n < 2 {
		return 0
	}
	result := n - 1
	for lc[n].Level > lc[result].Level {
		result--
	}
	return result
}
