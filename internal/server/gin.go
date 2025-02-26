package server

func (s *server) runGinServer() error {
	return s.r.Run(s.cfg.Http.Port)
}
