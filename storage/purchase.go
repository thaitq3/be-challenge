package storage

import (
	"context"
	"crud-challenge/dto"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	WagerId     uint
	BuyingPrice float64
	BoughtAt    time.Time
}

func (*Purchase) TableName() string {
	return "purchase"
}
func (p *Purchase) ToDTO() *dto.Purchase {
	return &dto.Purchase{
		ID:          p.ID,
		WagerId:     p.WagerId,
		BuyingPrice: p.BuyingPrice,
		BoughtAt:    p.BoughtAt.Unix(),
	}
}

var PurchaseRef = &Purchase{}

var PurchaseDao IPurchaseDAO = &PurchaseDAO{}

// IPurchaseDAO ...
//go:generate mockery --name=IPurchaseDAO --inpackage --case=snake
type IPurchaseDAO interface {
	Purchase(ctx context.Context,updatedWager *Wager, buyingPrice float64) (*Purchase, error)
}

type PurchaseDAO struct {
}

func (dao *PurchaseDAO) Purchase(ctx context.Context,updatedWager *Wager, buyingPrice float64) (*Purchase, error) {
	purchase := &Purchase{
		WagerId:     updatedWager.ID,
		BuyingPrice: buyingPrice,
		BoughtAt:    time.Now().UTC(),
	}

	txErr := gormConn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Updates(updatedWager).Error
		if err != nil {
			return err
		}

		err = tx.Create(purchase).Error
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return purchase, nil
}
