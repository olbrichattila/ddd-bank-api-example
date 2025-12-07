package transaction

//go:generate mockgen -destination=../../infrastructure/persistence/transaction/mock/account-mock.go -package=mock . Transaction
type Transaction interface {
	BelongToUser(accountNumber, transactionId string) (bool, error)
	Create(entity TransactionEntity) error
	List(accountNumber string) ([]TransactionEntity, error)
	Get(transactionNumber string) (TransactionEntity, error)
}
