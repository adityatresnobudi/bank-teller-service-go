package handler

func (s *serviceHandler) MapRoutes() {
	s.r.Group("/services").
		GET("", s.GetAll).
		GET("/:id", s.GetOne).
		POST("", s.Create).
		PUT("/:id", s.UpdateById).
		DELETE("/:id", s.DeleteById)
}
