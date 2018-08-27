package domain

type Paginator struct {
	PageNo   int `json:"page_no"`
	PageSize int `json:"page_size"`
}

func NewPaginator(pageNo int, pageSize int) (*Paginator) {
	if pageNo < 1 {
		pageNo = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return &Paginator{PageNo: pageNo, PageSize: pageSize}
}

func (paginator *Paginator) GetSkip() (int) {
	return (paginator.PageNo - 1) * paginator.PageSize
}

func (paginator *Paginator) GetLimit() (int) {
	return paginator.PageSize
}

func (paginator *Paginator) HasNext(totalCount int) (bool) {
	return totalCount > paginator.PageNo*paginator.PageSize
}

func (paginator *Paginator) PackResp(totalCount int) (*PaginatorResp) {
	var totalPage int
	if totalCount == 0 {
		totalPage = 0
	} else {
		totalPage = 1 + (totalCount-1)/paginator.PageSize
	}
	return &PaginatorResp{
		PageNo:     paginator.PageNo,
		PageSize:   paginator.PageSize,
		TotalPage:  totalPage,
		TotalCount: totalCount,
	}
}

type PaginatorResp struct {
	PageNo     int `json:"page_no"`
	PageSize   int `json:"page_size"`
	TotalPage  int `json:"total_page"`
	TotalCount int `json:"total_count"`
}
