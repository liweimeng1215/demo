package arglist

type ArgList struct {
	args         []string
	currentIndex int
}

func New(args []string) *ArgList {
	return &ArgList{args: args, currentIndex: -1}
}

func (al *ArgList) HasNext() bool {
	return al.currentIndex < len(al.args)-1
}

func (al *ArgList) Next() string {
	al.currentIndex++
	if al.currentIndex >= len(al.args) {
		panic("no next arg")
	}
	return al.args[al.currentIndex]
}
