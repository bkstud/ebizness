package structs

import "gorm.io/gorm"

// User is a retrieved and authentiacted user.
type Session struct {
	gorm.Model
	SessionToken  string
	ProviderToken string
}
