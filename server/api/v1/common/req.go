package common

type PageReq struct {
	PageNum  int `p:"pageNum"`  //当前页码
	PageSize int `p:"pageSize"` //每页数
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
