package filters

type OrderFilter struct {
	ID string `json:"id"`
}

type SystemFieldFilter struct {
	ID string `json:"id"`
}
type SystemGroupFieldFilter struct {
	ID string `json:"id"`
}
type ConfigSystemMasterFileFieldFilter struct {
	ID string `json:"id"`
}
type MasterFileFilter struct {
	ID string `json:"id"`
}
type LogMasterFileFilter struct {
	ID string `json:"id"`
}

type DocumentFilter struct {
	ID string `json:"id"`
}

type DocumentTemplateFilter struct {
	ID string `json:"id"`
}
type DocumentVersionFilter struct {
	ID string `json:"id"`
}
type DocumentTemplateFieldFilter struct {
	ID string `json:"id"`
}
type DocumentVersionFieldValueFilter struct {
	ID string `json:"id"`
}
type LogDocumentVersionFieldValueFilter struct {
	ID string `json:"id"`
}
