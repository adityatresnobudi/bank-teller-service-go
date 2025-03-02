package handler

func (u *userHandler) MapRoutes() {
	u.r.Group("/users").
		GET("", u.GetAll).
		GET("/:id", u.GetOne).
		POST("", u.Create).
		PUT("/:id", u.UpdateById).
		DELETE("/:id", u.DeleteById)
}
