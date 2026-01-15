// utils/slice.go
package utils

// Contains checks if a value exists in a slice
// T must be comparable (== supported)
func Contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
