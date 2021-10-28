package usecase

import "boilerplate/config/log"

type TransferRtgsUsecaseStruct struct {
	log *log.LogCustom
}

func NewTransferRtgsUsecase(log *log.LogCustom) Example2UsecaseInterface {
	return &TransferRtgsUsecaseStruct{log}
}

func (t TransferRtgsUsecaseStruct) IniFuncSatu()  {

}