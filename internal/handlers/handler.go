package handlers

import "ef_md_test/internal/services"

type Handler interface {
}

type handler struct {
	s *services.Service
}

func NewHandler(s services.Service) Handler {
	return &handler{s: &s}
}
