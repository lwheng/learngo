// Package dog allows you to learn more about dogs
package dog

// Years converts human years to dog years
func Years(h int) int {
	return h * 7
}

// YearsTwo converts human years to dog years
func YearsTwo(h int) int {
	count := 0
	for i := 0; i < h; i++ {
		count += 7
	}
	return count
}
