package dao

type TagDao struct {
	Dao

	Name string
	CreatedBy string
	ModifiedBy string
	State int
}

func GetTags(rowStart int, pageSize int, maps interface{}) (tags []TagDao) {
	db.Where(maps).Offset(rowStart).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag TagDao
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

func ExistTagById(id int) bool {
	var tag TagDao
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&TagDao{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&TagDao{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&TagDao{})
	return true
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ?", 0).Delete(&TagDao{})
	return true
}
