package sortList

type Contact struct {
	Name  string
	Phone int
}

type ContactList []Contact

func (c ContactList) Len() int {
	return len(c)
}
func (c ContactList) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}
func (c ContactList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
