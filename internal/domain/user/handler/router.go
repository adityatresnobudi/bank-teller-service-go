package handler

func (u *userHandler) MapRoutes() {
	u.r.Group("/users").
		GET("", u.GetAll)
}
