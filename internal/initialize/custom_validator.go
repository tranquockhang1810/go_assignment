package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/poin4003/yourVibes_GoApi/internal/consts"
	"mime/multipart"
)

func InitCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("privacy_enum", validatePrivacy)
		v.RegisterValidation("files", validateFiles)
		v.RegisterValidation("file", validateFile)
		v.RegisterValidation("language_setting", validateLanguageSetting)
	}
}

func validatePrivacy(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == string(consts.PUBLIC) || value == string(consts.PRIVATE) || value == string(consts.FRIEND_ONLY) {
		return true
	}

	return false
}

func validateLanguageSetting(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == string(consts.VI) || value == string(consts.EN) {
		return true
	}

	return false
}

func validateFiles(fl validator.FieldLevel) bool {
	files := fl.Field().Interface().([]multipart.FileHeader)

	for _, file := range files {
		if file.Size == 0 {
			return false
		}
	}

	return true
}

func validateFile(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(multipart.FileHeader)
	if !ok {
		return false
	}

	if file.Size == 0 {
		return false
	}

	return true
}
