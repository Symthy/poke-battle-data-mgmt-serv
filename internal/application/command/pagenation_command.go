package command

import "math"

type PagenationCommand struct {
	next     uint32
	pageSize uint16
}

func NewPaginationCommand(next uint64, pageSize uint64) PagenationCommand {
	start := uint32(1)
	size := uint16(100)
	if next > uint64(0) && next < uint64(math.MaxUint32) {
		start = uint32(next)
	}
	if pageSize > uint64(0) && pageSize <= uint64(1000) {
		size = uint16(pageSize)
	}
	return PagenationCommand{next: start, pageSize: size}
}

func (p PagenationCommand) Next() uint32 {
	return p.next
}

func (p PagenationCommand) PageSize() uint16 {
	return p.pageSize
}
