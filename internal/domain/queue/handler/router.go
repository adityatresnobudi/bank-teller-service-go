package handler

func (q *queueHandler) MapRoutes() {
	q.r.Group("/queues").
		GET("", q.GetAll).
		GET("/:id", q.GetOne).
		POST("", q.Create).
		PUT("", q.UpdateByQueueNum).
		DELETE("/:id", q.DeleteById)
}
