package user

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	validId          = "usr-d187b52cf4ee97e05e65a7ebd4fd7ef7"
	validName        = "John Doe"
	validLine1       = "123 Main Street"
	validLine2       = "Flat 4B"
	validTown        = "London"
	validCounty      = "Greater London"
	validPostcode    = "SW1A 1AA"
	validPhoneNumber = "+447911123456"
	validEmail       = "john.doe@example.com"
)

const (
	invalidId          = "tra-d187b52cf4ee97e05e65a7ebd4fd7ef7"
	invalidPhoneNumber = "-047911123456"
	invalidEmail       = "john.doeexample.com"
)

func TestUserEntity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User entity unit test")
}

var _ = Describe("User Entity test", func() {
	Context("Case: trying to create a user entity", func() {
		When("When entity is created with correct input data", func() {
			It("Than all getters return the correct values", func() {
				addressLine2 := validLine2
				now := time.Now()

				entity, err := New(UserInput{
					Id:          validId,
					Name:        validName,
					Line1:       validLine1,
					Line2:       &addressLine2,
					Line3:       nil,
					Town:        validTown,
					County:      validCounty,
					Postcode:    validPostcode,
					PhoneNumber: validPhoneNumber,
					Email:       validEmail,
					CreatedAt:   now,
					UpdatedAt:   now,
				})

				// Validation success
				Expect(err).ToNot(HaveOccurred())

				// All getters return correct values
				Expect(entity.Id().AsString()).To(Equal("usr-d187b52cf4ee97e05e65a7ebd4fd7ef7"))

				Expect(entity.Id().AsString()).To(Equal(validId))
				Expect(entity.Name()).To(Equal(validName))
				Expect(entity.Line1()).To(Equal(validLine1))
				Expect(entity.Line2()).To(Equal(&addressLine2))
				Expect(entity.Line3()).To(BeNil())
				Expect(entity.Town()).To(Equal(validTown))
				Expect(entity.County()).To(Equal(validCounty))
				Expect(entity.Postcode()).To(Equal(validPostcode))
				Expect(entity.PhoneNumber().AsString()).To(Equal(validPhoneNumber))
				Expect(entity.Email().AsString()).To(Equal(validEmail))
				Expect(entity.CreatedAt().Equal(now)).To(BeTrue())
				Expect(entity.UpdatedAt().Equal(now)).To(BeTrue())
			})
		})

		When("When entity is created with incorrect id", func() {
			It("Than returns error", func() {
				addressLine2 := validLine2
				now := time.Now()

				_, err := New(UserInput{
					Id:          invalidId,
					Name:        validName,
					Line1:       validLine1,
					Line2:       &addressLine2,
					Line3:       nil,
					Town:        validTown,
					County:      validCounty,
					Postcode:    validPostcode,
					PhoneNumber: validPhoneNumber,
					Email:       validEmail,
					CreatedAt:   now,
					UpdatedAt:   now,
				})

				// Validation failed
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("invalid user id"))
			})
		})

		When("When entity is created with incorrect phone number", func() {
			It("Than returns error", func() {
				addressLine2 := validLine2
				now := time.Now()

				_, err := New(UserInput{
					Id:          validId,
					Name:        validName,
					Line1:       validLine1,
					Line2:       &addressLine2,
					Line3:       nil,
					Town:        validTown,
					County:      validCounty,
					Postcode:    validPostcode,
					PhoneNumber: invalidPhoneNumber,
					Email:       validEmail,
					CreatedAt:   now,
					UpdatedAt:   now,
				})

				// Validation failed
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("invalid phone"))
			})
		})

		When("When entity is created with incorrect email", func() {
			It("Than returns error", func() {
				addressLine2 := validLine2
				now := time.Now()

				_, err := New(UserInput{
					Id:          validId,
					Name:        validName,
					Line1:       validLine1,
					Line2:       &addressLine2,
					Line3:       nil,
					Town:        validTown,
					County:      validCounty,
					Postcode:    validPostcode,
					PhoneNumber: validPhoneNumber,
					Email:       invalidEmail,
					CreatedAt:   now,
					UpdatedAt:   now,
				})

				// Validation failed
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("invalid email"))
			})
		})
	})
})
