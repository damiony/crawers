package model

type Item struct {
	Url    string
	Type   string
	Id     int
	Upload Profile
}

type SearchResult struct {
	Hits  int64
	Start int
	Query string
	Items []Item
}

type Profile struct {
	MemberId      int    // ID
	NickName      string // 姓名
	Sex           int    // 性别 0男 1女
	Height        int    // 身高
	Age           int    // 年龄
	Constellation string // 星座
	Education     string // 教育
	Marriage      string // 婚姻
}
