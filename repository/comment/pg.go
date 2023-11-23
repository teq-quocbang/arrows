package comment

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/teq-quocbang/arrows/model"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (r *pgRepository) CreateCommentInPost(ctx context.Context, c *model.Comment, p *model.Post) error {
	tx := r.getDB(ctx).Begin()
	// create comment
	if err := tx.Create(&c).Error; err != nil {
		tx.Rollback()
		return err
	}

	// TODO: should put this solution in post service
	// push comment to post
	commentIDs, err := json.Marshal(model.PostInfo{CommentIDs: append(p.Information.CommentIDs, c.ID)})
	if err != nil {
		return err
	}
	if err := tx.Model(&model.Post{}).Where("id = ?", p.ID).Update("information", gorm.Expr("JSON_MERGE_PATCH(information, ?)", commentIDs)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// commit
	tx.Commit()

	return nil
}

func (r *pgRepository) Update(ctx context.Context, commentID uuid.UUID, contents string) error {
	return r.getDB(ctx).Model(&model.Comment{}).Where("id = ?", commentID).Updates(map[string]interface{}{
		"contents": contents,
	}).Error
}

func (r *pgRepository) UpsertEmoji(ctx context.Context, commentID uuid.UUID, react *model.React) error {
	reactByte, err := json.Marshal(model.CommentInfo{Reacts: *react})
	if err != nil {
		return err
	}
	return r.getDB(ctx).Model(&model.Comment{}).Where("id = ?", commentID).Update("information", gorm.Expr("JSON_MERGE_PATCH(information, ?)", reactByte)).Error
}

func (r *pgRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	return r.getDB(ctx).Where("id = ?", commentID).Delete(&model.Comment{}).Error
}

func (r *pgRepository) CreateInParentComment(ctx context.Context, cChild *model.Comment, cParent *model.Comment) error {
	// start tx
	tx := r.getDB(ctx).Begin()

	// create child comment
	if err := tx.Create(&cChild).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update comment id to parent comment
	childCommentByte, err := json.Marshal(model.CommentInfo{ChildCommentIDs: append(cParent.Information.ChildCommentIDs, cChild.ID)})
	if err != nil {
		return err
	}
	if err := tx.Model(&model.Comment{}).Where("id = ?", cParent.ID).Update("information", gorm.Expr("JSON_MERGE_PATCH(information, ?)", childCommentByte)).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *pgRepository) GetByID(ctx context.Context, commentID uuid.UUID) (model.Comment, error) {
	comment := model.Comment{}
	if err := r.getDB(ctx).Where("id = ?", commentID).Take(&comment).Error; err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}
