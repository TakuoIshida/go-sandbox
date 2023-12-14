package service

import (
	todo_model "go-sandbox/domain/model"
	"go-sandbox/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type TagsServiceImpl struct {
	TagsRepository repository.ITodoRepository
}

func NewTagsServiceImpl(tagRepository repository.ITodoRepository) ITodoService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
	}
}

// Create implements TagsService
func (t *TagsServiceImpl) Create(ctx *gin.Context, todo todo_model.Todo) {
	t.TagsRepository.Save(ctx, todo)
}

// Delete implements TagsService
func (t *TagsServiceImpl) Delete(ctx *gin.Context, id int64) {
	t.TagsRepository.Delete(ctx, id)
}

// FindAll implements TagsService
func (t *TagsServiceImpl) FindAll(ctx *gin.Context) []todo_model.Todo {
	return t.TagsRepository.FindAll(ctx)
}

// FindById implements TagsService
func (t *TagsServiceImpl) FindById(ctx *gin.Context, id int64) todo_model.Todo {
	return t.TagsRepository.FindById(ctx, id)
}

// // Update implements TagsService
// func (t *TagsServiceImpl) Update(ctx *gin.Context, tags request.UpdateTagsRequest) {
// 	tagData, err := t.TagsRepository.FindById(tags.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = tags.Name
// 	t.TagsRepository.Update(tagData)
// }
