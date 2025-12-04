package transaction

type Transaction interface {
	Create(entity TransactionEntity) error
	List(accountNumber string) ([]TransactionEntity, error)
	Get(transactionNumber string) (TransactionEntity, error)
}
