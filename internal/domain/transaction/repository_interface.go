package transaction

type Transaction interface {
	BelongToUser(accountNumber, transactionId string) (bool, error)
	Create(entity TransactionEntity) error
	List(accountNumber string) ([]TransactionEntity, error)
	Get(transactionNumber string) (TransactionEntity, error)
}
