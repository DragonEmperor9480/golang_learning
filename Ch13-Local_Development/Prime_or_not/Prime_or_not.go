package Prime_or_not

func Prime(n int) bool {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}

	}
	if count == 2 {
		return true
	}
	return false

}
