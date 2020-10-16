package models

type QueryForm struct {
	Sort    string `query:"sort" form:"sort"`
	Size    int64 `query:"size" form:"size"`
}

type SubscribeQueryForm struct {
	QueryForm
	Username string `query:"username" form:"username"`
}
