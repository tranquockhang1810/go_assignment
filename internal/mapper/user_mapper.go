package mapper

import (
	"github.com/poin4003/yourVibes_GoApi/internal/consts"
	"github.com/poin4003/yourVibes_GoApi/internal/dtos/user_dto"
	"github.com/poin4003/yourVibes_GoApi/internal/model"
)

func MapUserToUserDto(user *model.User) *user_dto.UserDto {
	return &user_dto.UserDto{
		ID:           user.ID,
		FamilyName:   user.FamilyName,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Birthday:     user.Birthday,
		AvatarUrl:    user.AvatarUrl,
		CapwallUrl:   user.CapwallUrl,
		Privacy:      user.Privacy,
		Biography:    user.Biography,
		AuthType:     user.AuthType,
		AuthGoogleId: user.AuthGoogleId,
		PostCount:    user.PostCount,
		Status:       user.Status,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Setting:      MapSettingToSettingDto(&user.Setting),
	}
}

func MapUserToUserDtoWithoutSetting(
	user *model.User,
	friendStatus consts.FriendStatus,
) *user_dto.UserDtoWithoutSetting {
	return &user_dto.UserDtoWithoutSetting{
		ID:           user.ID,
		FamilyName:   user.FamilyName,
		Name:         user.Name,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Birthday:     user.Birthday,
		AvatarUrl:    user.AvatarUrl,
		CapwallUrl:   user.CapwallUrl,
		Privacy:      user.Privacy,
		Biography:    user.Biography,
		AuthType:     user.AuthType,
		AuthGoogleId: user.AuthGoogleId,
		PostCount:    user.PostCount,
		Status:       user.Status,
		FriendStatus: friendStatus,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func MapSettingToSettingDto(setting *model.Setting) user_dto.SettingDto {
	return user_dto.SettingDto{
		ID:        setting.ID,
		UserId:    setting.UserId,
		Language:  setting.Language,
		Status:    setting.Status,
		CreatedAt: setting.CreatedAt,
		UpdatedAt: setting.UpdatedAt,
	}
}

func MapUserToUserDtoShortVer(user *model.User) user_dto.UserDtoShortVer {
	return user_dto.UserDtoShortVer{
		ID:         user.ID,
		FamilyName: user.FamilyName,
		Name:       user.Name,
		AvatarUrl:  user.AvatarUrl,
	}
}

func MapToUserFromUpdateDto(
	input *user_dto.UpdateUserInput,
) map[string]interface{} {
	updateData := make(map[string]interface{})

	if input.FamilyName != nil {
		updateData["family_name"] = input.FamilyName
	}

	if input.Name != nil {
		updateData["name"] = input.Name
	}

	if input.Email != nil {
		updateData["email"] = input.Email
	}

	if input.PhoneNumber != nil {
		updateData["phone_number"] = input.PhoneNumber
	}

	if input.Birthday != nil {
		updateData["birthday"] = input.Birthday
	}

	if input.Privacy != nil {
		updateData["privacy"] = input.Privacy
	}

	if input.Biography != nil {
		updateData["biography"] = input.Biography
	}

	return updateData
}
