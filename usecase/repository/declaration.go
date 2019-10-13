package repository

import "slack-api/domain"

type DeclarationRepository interface {
	Store(*domain.Declaration) error
}
