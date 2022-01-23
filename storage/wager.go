package storage

import (
	"context"
	"crud-challenge/dto"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Wager struct {
	gorm.Model
	TotalWagerValue     int64
	Odds                int64
	SellingPercentage   int64
	SellingPrice        float64
	CurrentSellingPrice float64
	PercentageSold      sql.NullInt64
	AmountSold          sql.NullFloat64
	PlacedAt            time.Time
}

func (*Wager) TableName() string {
	return "wager"
}

func (w *Wager) ToDto() *dto.Wager {
	var percentageSold *int64
	var amountSold *float64

	if w.AmountSold.Valid {
		amountSold = &w.AmountSold.Float64
	}
	if w.PercentageSold.Valid {
		percentageSold = &w.PercentageSold.Int64
	}

	return &dto.Wager{
		Id:                  w.ID,
		TotalWagerValue:     w.TotalWagerValue,
		Odds:                w.Odds,
		SellingPercentage:   w.SellingPercentage,
		SellingPrice:        w.SellingPrice,
		CurrentSellingPrice: w.CurrentSellingPrice,
		PercentageSold:      percentageSold,
		AmountSold:          amountSold,
		PlacedAt:            w.PlacedAt.Unix(),
	}
}

var WagerRef = &Wager{}

var WagerDao IWagerDAO = &WagerDAO{}

// IWagerDAO ...
//go:generate mockery --name=IWagerDAO --inpackage --case=snake
type IWagerDAO interface {
	Create(ctx context.Context, wager *dto.Wager) (*Wager, error)
	List(ctx context.Context, page, limit int) ([]*Wager, error)
	GetById(ctx context.Context, id uint) (*Wager, error)
}

type WagerDAO struct {
}

func (dao *WagerDAO) Create(ctx context.Context, wager *dto.Wager) (*Wager, error) {
	insertWager := &Wager{
		TotalWagerValue:     wager.TotalWagerValue,
		Odds:                wager.Odds,
		SellingPercentage:   wager.SellingPercentage,
		SellingPrice:        wager.SellingPrice,
		CurrentSellingPrice: wager.CurrentSellingPrice,
		PercentageSold:      sql.NullInt64{},
		AmountSold:          sql.NullFloat64{},
		PlacedAt:            time.Now().UTC(),
	}

	err := gormConn.WithContext(ctx).Create(insertWager).Error
	if err != nil {
		return nil, err
	}

	return insertWager, nil
}

func (dao *WagerDAO) List(ctx context.Context, page, limit int) ([]*Wager, error) {
	results := []*Wager{}

	err := gormConn.
		WithContext(ctx).
		Model(WagerRef).
		Order("id desc").
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (dao *WagerDAO) GetById(ctx context.Context, id uint) (*Wager, error) {
	wager := &Wager{}
	err := gormConn.
		WithContext(ctx).
		First(wager, "id", id).Error
	if err != nil {
		return nil, err
	}

	return wager, nil
}
