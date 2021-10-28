package error

import (
	"boilerplate/models"
	"boilerplate/models/contract"
	"boilerplate/usecase"
	"errors"
	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"net/http"
)

type errorHandlerUsecase struct {
}

func NewErrorHandlerUsecase() usecase.ErrorHandlerUsecase {
	return &errorHandlerUsecase{}
}

func (eh *errorHandlerUsecase) ResponseError(A interface{}) (int, interface{}) {
	var T interface{}
	if A.(*gin.Error).Meta != nil {
		T = A.(*gin.Error).Meta
	} else {
		T = A.(*gin.Error).Err
	}

	switch T.(type) {
	case error:
		// check error type is postgres's error
		if _, ok := T.(*pq.Error); ok {
			switch T.(*pq.Error).Code.Name() {
			case "unique_violation":
				return errcntrct.ErrorMessage(http.StatusBadRequest, "", errors.New(contract.ErrDuplicate))
			}
		}

		// check error with type common error, and check .Error()
		switch T.(error).Error() {
		case gorm.ErrRecordNotFound.Error():
			return errcntrct.ErrorMessage(http.StatusBadRequest, "", errors.New(contract.ErrRecordNotFound))
		case contract.ErrInvalidParam:
			return errcntrct.ErrorMessage(http.StatusBadRequest, "", T.(error))
		case contract.ErrInvalidRoleName:
			return errcntrct.ErrorMessage(http.StatusBadRequest, "", T.(error))
		case contract.ErrRoleName:
			return errcntrct.ErrorMessage(http.StatusBadRequest, "", T.(error))
		case contract.ErrBankCardType:
			return errcntrct.ErrorMessage(http.StatusBadRequest, "", T.(error))
		case contract.ErrBankCardNo:
			return errcntrct.ErrorMessage(http.StatusBadRequest, "", T.(error))
		case contract.ErrIdentificationType:
			return errcntrct.ErrorMessage(http.StatusBadRequest, "", T.(error))
		}
	}

	return errcntrct.ErrorMessage(http.StatusInternalServerError, "", errors.New(contract.ErrUnexpectedError))

}

func (eh *errorHandlerUsecase) ValidateRequest(T interface{}) error {
	v := validator.New()
	var errArr error
	switch T.(type) {
	case models.Role:
		err := v.Struct(T)
		if err == nil {
			return nil
		}
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Field() {

			case "RoleName":
				errArr = errors.New(contract.ErrRoleName)
			}
		}

		if errArr != nil {
			return errArr
		}

		return nil


	default:
		return errors.New(contract.ErrUnexpected)
	}

}
