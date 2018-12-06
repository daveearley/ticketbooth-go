package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type transactionRepository struct {
	db *sql.DB
}

type TransactionRepository interface {
	Store(event *models.Transaction) (*models.Transaction, error)
	GetByUUID(uuid string) (*models.Transaction, error)
}

//NewTransactionRepository returns a new instance of transactionRepository
func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db}
}

//GetByUUID gets a transaction by a given UUID
func (r *transactionRepository) GetByUUID(uuid string) (*models.Transaction, error) {
	transaction, err := models.Transactions(qm.Where("uuid=?", uuid)).One(r.db)

	if err != nil {
		return nil, getErrorType(err, app.TransactionResource, uuid)
	}

	return transaction, nil
}

//Store creates a transaction
func (r *transactionRepository) Store(transaction *models.Transaction) (*models.Transaction, error) {
	err := transaction.Insert(r.db, boil.Infer())

	if err != nil {
		return nil, app.ServerError(err)
	}

	return transaction, nil
}
