// Package twofer ...
package twofer

// ShareWith returns String with name argument
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
