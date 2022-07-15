package rest

type PageData struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}
