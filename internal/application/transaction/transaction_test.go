package transaction

import (
	"fmt"
	"testing"

	accountMock "eaglebank/internal/infrastructure/persistence/account/mock"
	transactionMock "eaglebank/internal/infrastructure/persistence/transaction/mock"
	transactionWoUMock "eaglebank/internal/infrastructure/workofunits/mock"

	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	accountNumber     = "01000004"
	transactionNumber = "tan-DZuYEQ"
)

func TestUserEntity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transaction service unit")
}

var _ = Describe("transaction service test", func() {
	var (
		err  error
		ctrl *gomock.Controller

		transactionRepositoryMock    *transactionMock.MockTransaction
		accountRepositoryMock        *accountMock.MockAccount
		transactionWoURepositoryMock *transactionWoUMock.MockTransaction

		transactionService Transaction
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		transactionRepositoryMock = transactionMock.NewMockTransaction(ctrl)
		accountRepositoryMock = accountMock.NewMockAccount(ctrl)
		transactionWoURepositoryMock = transactionWoUMock.NewMockTransaction(ctrl)

		transactionService, err = New(
			transactionRepositoryMock,
			accountRepositoryMock,
			transactionWoURepositoryMock,
		)

		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Case: Belongs to user", func() {
		When("When: calling the function and repository return true, no error", func() {
			It("Than: it returns correct values", func() {
				// Assert
				transactionRepositoryMock.
					EXPECT().
					BelongToUser(accountNumber, transactionNumber).
					Return(true, nil)

				// Act
				ok, err := transactionService.BelongToUser(accountNumber, transactionNumber)

				// Assert
				Expect(err).NotTo(HaveOccurred())
				Expect(ok).To(BeTrue())
			})
		})

		When("When: calling the function and repository return false, no error", func() {
			It("Than: it returns correct values", func() {
				// Assert
				transactionRepositoryMock.
					EXPECT().
					BelongToUser(accountNumber, transactionNumber).
					Return(false, nil)

				// Act
				ok, err := transactionService.BelongToUser(accountNumber, transactionNumber)

				// Assert
				Expect(err).NotTo(HaveOccurred())
				Expect(ok).To(BeFalse())
			})
		})

		When("When: calling the function and repository return false, and error", func() {
			It("Than: it returns correct values", func() {
				// Assert
				transactionRepositoryMock.
					EXPECT().
					BelongToUser(accountNumber, transactionNumber).
					Return(false, fmt.Errorf("repository error"))

				// Act
				ok, err := transactionService.BelongToUser(accountNumber, transactionNumber)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("repository error"))
				Expect(ok).To(BeFalse())
			})
		})
	})
})
