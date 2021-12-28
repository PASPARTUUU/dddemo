package server

// Domain - мета данные домена
type Domain interface {
	DomainName() string
	RootFolderPath() string
	RootTemplatesFolder() string // format - ./domains/shop/web/templates
}
