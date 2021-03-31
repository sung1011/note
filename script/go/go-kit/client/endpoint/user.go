package endpoint

type UserReq struct {
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
}

type UserResp struct {
	Name string `json:"name"`
}
