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

func (r *pgRepository) DeleteChild(ctx context.Context, cChildID uuid.UUID, cParentID uuid.UUID) error {
	// start tx
	tx := r.getDB(ctx).Begin()

	// delete child comment
	if err := tx.Where("id = ?", cChildID).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// remove comment ids in parent
	if err := tx.Model(&model.Comment{}).
		Where("id = ? and is_parent = true", cParentID).
		Update("information", gorm.Expr("JSON_REMOVE(information, replace(JSON_SEARCH(information, 'one', ?), '\"', ''))", cChildID)).Error; err != nil {
		tx.Rollback()
		return nil
	}

	// commit
	tx.Commit()

	return nil
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

// DeleteParent is delete at parent comment and then clear it child comment and remove comment id in the post
func (r *pgRepository) DeleteParent(ctx context.Context, c model.Comment) error {
	// start tx
	tx := r.getDB(ctx).Begin()

	// delete parent comment
	if err := tx.Where("id = ?", c.ID).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// delete child comment
	if err := tx.Where("id in (?) and is_parent=false", c.Information.ChildCommentIDs).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// remove comment id in post
	// remove comment_id from list of comment ids in post and update to information
	if err := tx.Model(&model.Post{}).
		Where("id = ?", c.PostID).
		Update("information", gorm.Expr("JSON_REMOVE(information, replace(JSON_SEARCH(information, 'one', ?), '\"', ''))", c.ID)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// commit
	tx.Commit()

	return nil
}

func (r *pgRepository) GetByIDs(ctx context.Context, IDs []uuid.UUID) ([]model.Comment, error) {
	comments := []model.Comment{}
	if err := r.getDB(ctx).Where("id in (?)", IDs).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
