package articledto

type CreateArticleRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  []int  `json:"categoryId" form:"categoryId" validate:"required"`
	UserID      int    `json:"userId" validate:"required"`
}

type UpdateArticleRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  []int  `json:"categoryId" form:"categoryId" validate:"required"`
	UserID      int    `json:"userId" validate:"required"`
}
