package user

import (
	"encoding/json"
	"time"

	userDomain "eaglebank/internal/domain/user"
)

func (h *Handler) outboundMappingTranslator(userEntity userDomain.UserEntity) ([]byte, error) {
	response := createUserResponse{
		Id:   userEntity.Id().AsString(),
		Name: userEntity.Name(),
		Address: addressDTO{
			Line1:    userEntity.Line1(),
			Line2:    userEntity.Line2(),
			Line3:    userEntity.Line3(),
			Town:     userEntity.Town(),
			County:   userEntity.County(),
			Postcode: userEntity.Postcode(),
		},
		PhoneNumber:      userEntity.PhoneNumber().AsString(),
		Email:            userEntity.Email().AsString(),
		CreatedTimestamp: userEntity.CreatedAt().Format(time.DateTime),
		UpdatedTimestamp: userEntity.UpdatedAt().Format(time.DateTime),
	}

	return json.Marshal(response)
}
