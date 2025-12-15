package transaction

import (
	"fmt"
	"testing"
	"time"

	"atybank/internal/domain/account"
	accountMock "atybank/internal/infrastructure/persistence/account/mock"
	transactionMock "atybank/internal/infrastructure/persistence/transaction/mock"
	transactionWoUMock "atybank/internal/infrastructure/workofunits/mock"

	"github.com/shopspring/decimal"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	accountNumber           = "01000004"
	transactionNumber       = "tan-DZuYEQ"
	userId                  = "usr-d187b52cf4ee97e05e65a7ebd4fd7ef7"
	transactionTypeDeposit  = "deposit"
	transactionTypeWithdraw = "withdrawal"
	currency                = "GBP"
	sortCode                = "10-10-10"
	accountType             = "personal"
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

	Context("Case: Create transaction", func() {
		When("When: calling the function, balance is positive, updates new balance, records transaction", func() {
			It("Than: runs without error, repository called", func() {
				// Assume
				returnAccountEntity, err := account.New(
					account.Input{
						AccountNumber: accountNumber,
						UserId:        userId,
						SortCode:      sortCode,
						Name:          "John Doe",
						AccountType:   accountType,
						Balance:       decimal.NewFromFloat(100),
						Currency:      currency,
						CreatedAt:     time.Now(),
						UpdatedAt:     time.Now(),
					},
				)

				// Assert
				Expect(err).NotTo(HaveOccurred())

				accountRepositoryMock.
					EXPECT().
					Get(accountNumber).
					Return(
						returnAccountEntity,
						nil,
					)

				transactionWoURepositoryMock.
					EXPECT().
					Create(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)

				// Act
				err = transactionService.Create(
					decimal.NewFromFloat(100),
					userId,
					currency,
					transactionTypeDeposit,
					accountNumber,
					nil,
				)

				// Assert
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("When: calling the function, balance is positive, but user tries to withdraw more", func() {
			It("Than: it returns error", func() {
				// Assume
				returnAccountEntity, err := account.New(
					account.Input{
						AccountNumber: accountNumber,
						UserId:        userId,
						SortCode:      sortCode,
						Name:          "John Doe",
						AccountType:   accountType,
						Balance:       decimal.NewFromFloat(100),
						Currency:      currency,
						CreatedAt:     time.Now(),
						UpdatedAt:     time.Now(),
					},
				)

				// Assert
				Expect(err).NotTo(HaveOccurred())

				accountRepositoryMock.
					EXPECT().
					Get(accountNumber).
					Return(
						returnAccountEntity,
						nil,
					)

				// Act
				err = transactionService.Create(
					decimal.NewFromFloat(200),
					userId,
					currency,
					transactionTypeWithdraw,
					accountNumber,
					nil,
				)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Balance cannot be negative"))
			})
		})

		When("When: calling with negative amount", func() {
			It("Than: returns the correct error", func() {
				// Assume
				returnAccountEntity, err := account.New(
					account.Input{
						AccountNumber: accountNumber,
						UserId:        userId,
						SortCode:      sortCode,
						Name:          "John Doe",
						AccountType:   accountType,
						Balance:       decimal.NewFromFloat(100),
						Currency:      currency,
						CreatedAt:     time.Now(),
						UpdatedAt:     time.Now(),
					},
				)

				// Assert
				Expect(err).NotTo(HaveOccurred())

				accountRepositoryMock.
					EXPECT().
					Get(accountNumber).
					Return(
						returnAccountEntity,
						nil,
					)

				// Act
				err = transactionService.Create(
					decimal.NewFromFloat(-100), // negative value
					userId,
					currency,
					transactionTypeDeposit,
					accountNumber,
					nil,
				)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("amount cannot must be between 0 and 10000"))
			})
		})

		When("When: calling with too large amount", func() {
			It("Than: returns the correct error", func() {
				// Assume
				returnAccountEntity, err := account.New(
					account.Input{
						AccountNumber: accountNumber,
						UserId:        userId,
						SortCode:      sortCode,
						Name:          "John Doe",
						AccountType:   accountType,
						Balance:       decimal.NewFromFloat(100),
						Currency:      currency,
						CreatedAt:     time.Now(),
						UpdatedAt:     time.Now(),
					},
				)

				// Assert
				Expect(err).NotTo(HaveOccurred())

				accountRepositoryMock.
					EXPECT().
					Get(accountNumber).
					Return(
						returnAccountEntity,
						nil,
					)

				// Act
				err = transactionService.Create(
					decimal.NewFromFloat(10001), // negative value
					userId,
					currency,
					transactionTypeDeposit,
					accountNumber,
					nil,
				)

				// Assert
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("amount cannot must be between 0 and 10000"))
			})
		})

		// TODO, we could add further test, with all the validation rules working on domain entity level
		// Testing 0, and max 10000 amount boundary is working, not just bigger, smaller
		// amd more...
	})
})
