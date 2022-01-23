package wager

import (
	"crud-challenge/dto"
	"crud-challenge/dto/request"
	"crud-challenge/storage"
	"crud-challenge/utils"
	"crud-challenge/utils/validator"
	"encoding/json"
	"errors"
	"net/http"
)

type CreateWagerHandler struct {
	wagerStorage storage.IWagerDAO `container:"type"`
}

func (u *CreateWagerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	request := new(request.CreateWagerRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		utils.OutputBadRequest(w, err)
		return
	}

	err := validator.Validate(request)
	if err != nil {
		utils.OutputBadRequest(w, err)
		return
	}

	if request.SellingPrice <= float64(request.TotalWagerValue) * (float64(request.SellingPercentage)/100) {
		utils.OutputBadRequest(w, errors.New("selling_price must be greater than total_wager_value * (selling_percentage / 100)"))
		return
	}

	createdWager, err := u.wagerStorage.Create(r.Context(), &dto.Wager{
		TotalWagerValue:     request.TotalWagerValue,
		Odds:                request.Odds,
		SellingPercentage:   request.SellingPercentage,
		SellingPrice:        request.SellingPrice,
		CurrentSellingPrice: request.SellingPrice,
	})
	if err != nil {
		utils.OutputInternalServerError(w, err)
		return
	}

	utils.OutputData(w, http.StatusCreated, createdWager.ToDto())
}
