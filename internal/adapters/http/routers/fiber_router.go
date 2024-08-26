package routers

import handlers "github.com/billowdev/exclusive-go-hexa/internal/adapters/http/handlers/system_fields"

func (r RouterImpls) CreateSystemFieldRoute(h handlers.ISystemFieldHandler) {
	r.route.Get("/system-fields",
		h.HandleGetSystemFields)
	r.route.Get("/system-fields/:id",
		h.HandleGetSystemField)
}
