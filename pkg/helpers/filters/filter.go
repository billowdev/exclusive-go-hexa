package filters

type OrderFilter struct {
	ID uint `json:"id"`
}

type SystemFieldFilter struct {
	ID uint `json:"id"`
}
type SystemGroupFieldFilter struct {
	ID uint `json:"id"`
}
type ConfigSystemMasterFileFieldFilter struct {
	ID uint `json:"id"`
}
type MasterFileFilter struct {
	ID uint `json:"id"`
}
type LogMasterFileFilter struct {
	ID uint `json:"id"`
}

type DocumentFilter struct {
	ID uint `json:"id"`
}

type DocumentTemplateFilter struct {
	ID uint `json:"id"`
}
type DocumentVersionFilter struct {
	ID uint `json:"id"`
}
type DocumentTemplateFieldFilter struct {
	ID uint `json:"id"`
}
type DocumentVersionFieldValueFilter struct {
	ID uint `json:"id"`
}
type LogDocumentVersionFieldValueFilter struct {
	ID uint `json:"id"`
}
