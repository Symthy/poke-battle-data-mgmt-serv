package command

type PagenationCommand struct {
	next     int
	pageSize int
}

func NewPaginationCommand(next int, pageSize int) PagenationCommand {
	start := 1
	size := 100
	if next != 0 {
		start = next
	}
	if pageSize != 0 {
		size = pageSize
	}
	return PagenationCommand{next: start, pageSize: size}
}

func (p PagenationCommand) Next() int {
	return p.next
}

func (p PagenationCommand) PageSize() int {
	return p.pageSize
}
