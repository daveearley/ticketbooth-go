package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type transactionRepository struct {
	db *sql.DB
}

type TransactionRepository interface {
	Create(event *models.Transaction) (*models.Transaction, error)
	FindByUUID(uuid string) (*models.Transaction, error)
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) FindByUUID(uuid string) (*models.Transaction, error) {
	return models.Transactions(qm.Where("uuid=?", uuid)).One(r.db)
}

func (r *transactionRepository) Create(transaction *models.Transaction) (*models.Transaction, error) {
	err := transaction.Insert(r.db, boil.Infer())

	return transaction, err
}
