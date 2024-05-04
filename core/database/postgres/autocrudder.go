package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/trillyai/backend-microservices/core/automapper"
)

func Create[Dest any, Source any](ctx context.Context, req any) (Dest, error) {
	connectToDB()
	defer closeDB()

	var resp Dest
	var newItem Source
	automapper.MapLoose(req, &newItem)

	db := DB.Debug().WithContext(ctx).Model(&newItem).Create(&newItem)

	if db.Error != nil {
		return resp, db.Error
	}

	automapper.MapLoose(newItem, &resp)

	return resp, nil
}

func Read[Dest any, Source any](ctx context.Context, rule any, args ...any) (Dest, error) {
	connectToDB()
	defer closeDB()

	var resp Dest
	var existingItem Source

	db := DB.Debug().WithContext(ctx).Model(&existingItem).Where(map[string]interface{}{"IsDeleted": false}).Where(rule, args...).Find(&resp)

	if db.Error != nil {
		return resp, db.Error
	}

	return resp, nil
}

func PaginatedRead[Dest any, Source any](ctx context.Context, rule any, offset uint32, limit uint32, args ...any) (Dest, error) {
	connectToDB()
	defer closeDB()

	var resp Dest
	var existingItem Source

	db := DB.Debug().WithContext(ctx).
		Model(&existingItem).
		Where(map[string]interface{}{"IsDeleted": false}).
		Where(rule, args...).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&resp)

	if db.Error != nil {
		return resp, db.Error
	}

	return resp, nil
}

func Update[Dest any, Source any](ctx context.Context, rule any, req any) (Dest, error) {
	connectToDB()
	defer closeDB()

	var resp Dest
	var existingItem Source
	automapper.MapLoose(req, &existingItem)

	db := DB.WithContext(ctx).Model(&existingItem).Where(map[string]interface{}{"IsDeleted": false}).Where(rule).Updates(&existingItem)

	if db.Error != nil {
		return resp, db.Error
	}

	automapper.MapLoose(existingItem, &resp)

	return resp, nil
}

func Delete[Dest any, Source any](ctx context.Context, rule any) (Dest, error) {
	connectToDB()
	defer closeDB()

	var resp Dest
	var existingItem Source

	now := time.Now()
	updtes := map[string]interface{}{"IsDeleted": true, "DeletedDate": &now}

	db := DB.WithContext(ctx).Model(&existingItem).Where(map[string]interface{}{"IsDeleted": false}).Where(rule).Updates(updtes)

	if db.Error != nil {
		return resp, db.Error
	}

	if db.RowsAffected <= 0 {
		return resp, errors.New("item not found")
	}

	automapper.MapLoose(existingItem, &resp)

	return resp, nil
}
