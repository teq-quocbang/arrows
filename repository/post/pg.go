package post

import (
	"context"
	"reflect"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/codetype"
	"github.com/teq-quocbang/arrows/model"
	"github.com/teq-quocbang/arrows/proto"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (r *pgRepository) Create(ctx context.Context, p *model.Post) error {
	return r.getDB(ctx).Create(&p).Error
}

func (r *pgRepository) Update(ctx context.Context, postID uuid.UUID, content string, privacyMode proto.Privacy) error {
	return r.getDB(ctx).Model(&model.Post{}).Where("id = ?", postID).Updates(map[string]interface{}{
		"content":      content,
		"privacy_mode": privacyMode,
	}).Error
}

func (r *pgRepository) GetByID(ctx context.Context, id uuid.UUID) (model.Post, error) {
	post := model.Post{}
	err := r.getDB(ctx).Where("id = ?", id).Take(&post).Error
	if err != nil {
		return model.Post{}, err
	}
	return post, nil
}

func (r *pgRepository) GetList(ctx context.Context, accountID uuid.UUID, order []string, paginator codetype.Paginator) ([]model.Post, int64, error) {
	var (
		db     = r.getDB(ctx).Model(&model.Post{})
		data   = make([]model.Post, 0)
		total  int64
		offset int
	)

	if !reflect.DeepEqual(accountID, uuid.UUID{}) {
		db = db.Where("create_by = ?", accountID)
	}

	for i := range order {
		db = db.Order(order[i])
	}

	if paginator.Page != 1 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if paginator.Limit != -1 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(paginator.Limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if paginator.Limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}

func (r *pgRepository) UpsertEmoji(ctx context.Context, postID uuid.UUID, react *model.React) error {
	return r.getDB(ctx).Model(&model.Post{}).Where("id = ?", postID).Update("information", gorm.Expr("JSON_MERGE_PATCH(information, ?)", react)).Error
}

func (r *pgRepository) Delete(ctx context.Context, postID uuid.UUID) error {
	return r.getDB(ctx).Where("id = ?", postID).Delete(&model.Post{}).Error
}
