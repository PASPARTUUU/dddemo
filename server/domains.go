package server

import "dddemo/models"

// Domain - мета данные домена
type Domain interface {
	DomainName() string
	// RootFolderPath() string
	RootTemplatesFolder() string // format: domains/shop/web/templates
	RootStaticFolder() string    // format: domains/shop/web/static
	PathsToCSSFiles() []string   // format: /domains/shop/web/static/style.css
	PathsToJSFiles() []string    // format: /domains/shop/web/static/script.js

	SidebarMarkup() models.SidebarMarkup
}
