package wager

import (
	"crud-challenge/utils/validator"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"crud-challenge/dto/request"
	"crud-challenge/storage"
	"crud-challenge/utils"
	distributedlock "crud-challenge/utils/distributed-lock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type PurchaseWagerHandler struct {
	wagerStorage    storage.IWagerDAO `container:"type"`
	purchaseStorage storage.IPurchaseDAO `container:"type"`
	distributedLock distributedlock.DistributedLock `container:"type"`
}

var (
	maximumLockTimeSecond          = int(12) // 12 seconds
	waitTime = 50 // 5 seconds
)

func (h *PurchaseWagerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	request := new(request.PurchaseRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		utils.OutputBadRequest(w, err)
		return
	}

	err := validator.Validate(request)
	if err != nil {
		utils.OutputBadRequest(w, err)
		return
	}

	params := mux.Vars(r)
	wagerIdString, ok := params["wager_id"]
	if !ok {
		utils.OutputBadRequest(w, errors.New("missing wager_id"))
		return
	} else {
		wagerId, err := strconv.Atoi(wagerIdString)
		if err != nil {
			utils.OutputBadRequest(w, errors.New("couldn't get wager_id"))
			return
		}

		request.WagerId = uint(wagerId)
	}

	alreadyLocked, unlockFunc := h.acquireLock(fmt.Sprintf("purchase_wager_wager_id_%d", request.WagerId))
	if alreadyLocked {
		utils.OutputInternalServerError(w, errors.New("please try again"))
		return
	}

	defer unlockFunc()


	wager, err := h.wagerStorage.GetById(r.Context(), request.WagerId)
	if err != nil {
		utils.OutputNotFound(w, errors.New("couldn't find wager"))
		return
	}

	if wager.CurrentSellingPrice < request.BuyingPrice {
		utils.OutputBadRequest(w, errors.New("selling_price must be greater than buying_price"))
		return
	}

	wager.CurrentSellingPrice -= request.BuyingPrice
	wager.AmountSold.Float64 += request.BuyingPrice
	wager.AmountSold.Valid = true
	wager.PercentageSold.Int64 = int64(wager.AmountSold.Float64 / wager.SellingPrice * 100)
	wager.PercentageSold.Valid = true

	purchase, err := h.purchaseStorage.Purchase(r.Context(), wager, request.BuyingPrice)
	if err != nil {
		utils.OutputInternalServerError(w, err)
		return
	}

	utils.OutputData(w, http.StatusCreated, purchase.ToDTO())
}


func (h *PurchaseWagerHandler) acquireLock(lockKey string) (alreadyLocked bool, unlockFunc func()) {
	mutex := h.distributedLock.NewMutex(lockKey, maximumLockTimeSecond, waitTime)

	mutexErr := mutex.Lock()
	if mutexErr != nil {
		logrus.Warn( "acquireLock error when try to lock", "error", mutexErr)
	}
	alreadyLocked = mutexErr == distributedlock.ErrLockNotObtained

	return alreadyLocked, func() {
		if mutexErr == nil {
			if err := mutex.Unlock(); err != nil {
				logrus.Warn( "acquireLock error when unlock", "error", err)
			}
		}
	}
}
