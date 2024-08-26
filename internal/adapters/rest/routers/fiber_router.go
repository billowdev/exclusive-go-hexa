package routers

import handlers "github.com/billowdev/document-system-field-manager/internal/adapters/rest/handlers/system_fields"

func (r RouterImpls) CreateSystemFieldRoute(h handlers.ISystemFieldHandler) {
	r.route.Get("/system-fields",
		h.HandleGetSystemFields)
	r.route.Get("/system-fields/:id",
		h.HandleGetSystemField)
}
