package repository

import (
	"slack-api/domain"
	"time"

	"github.com/guregu/dynamo"
)

type declRepo struct {
	table *dynamo.Table
}

type DeclarationRepository interface {
	Store(decl *domain.Declaration) error
}

func NewDeclarationRepository(table *dynamo.Table) DeclarationRepository {
	return &declRepo{table}
}

func (r *declRepo) Store(decl *domain.Declaration) error {
	decl.CreatedAt = time.Now().UTC()
	err := r.table.Put(decl).Run()
	return err
}
