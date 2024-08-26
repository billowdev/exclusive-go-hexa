package routers

import "github.com/gofiber/fiber/v2"

type RouterImpls struct {
	route fiber.Router
}

func NewRoute(r fiber.Router) RouterImpls {
	return RouterImpls{r}
}
