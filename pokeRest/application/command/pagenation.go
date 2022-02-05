package command

type Pagenation struct {
	page     int
	pageSize int
}

func NewPagination(page int, pageSize int) Pagenation {
	return Pagenation{page: page, pageSize: pageSize}
}

func (p Pagenation) Page() int {
	return p.page
}

func (p Pagenation) PageSize() int {
	return p.pageSize
}
