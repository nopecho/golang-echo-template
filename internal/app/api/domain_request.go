package api

type GetParams struct {
	ID int64 `param:"id"`
}

type MetaQuery struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

type CreateRequest struct {
	Name string `json:"name"`
}

type UpdateRequest struct {
	ID   int64  `param:"id"`
	Name string `json:"name"`
}
