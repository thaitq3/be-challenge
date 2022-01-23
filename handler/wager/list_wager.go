package wager

import (
	"crud-challenge/dto"
	"errors"
	"net/http"
	"strconv"

	"crud-challenge/storage"
	"crud-challenge/utils"
)

type ListWagerHandler struct {
	wagerStorage storage.IWagerDAO `container:"type"`
}

func (h *ListWagerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var page, limit int
	var err error

	 urlParams := r.URL.Query()


	pageStr := urlParams.Get("page")
	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			utils.OutputBadRequest(w, errors.New("couldn't get page"))
			return
		}
	}

	limitStr := urlParams.Get("limit")
	if limitStr == "" {
		limit = 20
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			utils.OutputBadRequest(w, errors.New("couldn't get limit"))
			return
		}
	}

	wagers, err := h.wagerStorage.List(r.Context(), page, limit)
	if err != nil {
		utils.OutputInternalServerError(w, err)
	}

	results := make([]*dto.Wager, len(wagers))
	for idx, wager := range wagers {
		results[idx] = wager.ToDto()
	}

	utils.OutputData(w,200, results)
}
