package dao

import (
	"fmt"
	"github.com/AiLiaa/blog-service/internal/model"
	"github.com/AiLiaa/blog-service/pkg/app"
)

/*
Common：指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小写下划线驼峰。若情况特殊，也可以编写该结构体的 TableName 方法用于指定其对应返回的表名。
Where：设置筛选条件，接受 map，struct 或 string 作为条件。
Offset：偏移量，用于指定开始返回记录之前要跳过的记录数。
Limit：限制检索的记录数。
Find：查找符合筛选条件的记录。
Update：更新所选字段。
Delete：删除数据。
Count：统计行为，用于统计模型的记录数。
*/

// GetTagCount 获取tag的数量
func (dao *Dao) GetTagCount(name string, state uint8) (int, error) {
	var count int
	tag := model.Tag{Name: name, State: state}
	db := dao.engine

	if tag.Name != "" {
		db = db.Where("name = ?", tag.Name)
	}
	db = db.Where("state = ?", tag.State)
	if err := db.Model(&tag).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetTagList 获取多个tag
func (dao *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	var tags []*model.Tag
	var err error
	tag := model.Tag{Name: name, State: state}
	db := dao.engine

	pageOffset := app.GetPageOffset(page, pageSize)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if tag.Name != "" {
		db = db.Where("name = ?", tag.Name)
	}
	db = db.Where("state = ?", tag.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// CreateTag 创建tag
func (dao *Dao) CreateTag(name string, state uint8, createdBy string) error {
	db := dao.engine
	tag := model.Tag{
		Name:   name,
		State:  state,
		Common: &model.Common{CreatedBy: createdBy},
	}
	return db.Create(&tag).Error
}

// UpdateTag 更新tag
func (dao *Dao) UpdateTag(id uint32, name string, state *uint8, modifiedBy string) error {
	db := dao.engine
	tag := model.Tag{
		Name:   name,
		State:  *state,
		Common: &model.Common{ID: id, ModifiedBy: modifiedBy},
	}

	fmt.Println(*state)
	return db.Model(&tag).Where("id = ? AND is_del = ?", tag.ID, 0).Update(tag).Error
}

// DeleteTag 删除tag
func (dao *Dao) DeleteTag(id uint32) error {
	db := dao.engine
	tag := model.Tag{Common: &model.Common{ID: id}}
	return db.Where("id = ? AND is_del = ?", tag.Common.ID, 0).Delete(&tag).Error
}
