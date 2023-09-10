package app

import (
	"net/http"

	"github.com/abrarnaim015/belajar-golang-dependency-injection/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server {
		Addr: ":3000",
		Handler: authMiddleware,
	}
}