package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	Id_task     int       `form:"id_task" json:"id_task" xml:"id_task" gorm:"column:id_task;primary_key;auto_increment;not null;"`
	Id_employe  int       `form:"id_employe" json:"id_employe" xml:"id_employe" binding:"required" gorm:"column:id_employe;type:int(11);not null;"`
	Name_task   string    `form:"name_task" json:"name_task" xml:"name_task" binding:"required" gorm:"column:name_task;type:varchar(200);not null;"`
	Status_task int       `form:"status_task" json:"status_task" xml:"status_task" gorm:"column:status_task;type:int(11);"`
	Entry_user  string    `form:"entry_user" json:"entry_user" xml:"entry_user" gorm:"column:entry_user;type:varchar(200);"`
	Entry_date  time.Time `form:"entry_date" json:"entry_date" xml:"entry_date" gorm:"column:entry_date;type:datetime;"`
	Update_user string    `form:"update_user" json:"update_user" xml:"update_user" gorm:"column:update_user;type:varchar(200);"`
	Update_date time.Time `form:"update_date" json:"update_date" xml:"update_date" gorm:"column:update_date;type:datetime;"`
	Delete_date time.Time `form:"delete_date" json:"delete_date" xml:"delete_date" gorm:"column:delete_date;type:datetime;"`
	table       string    `gorm:"-"`
}

func (p Task) TableName() string {
	// double check here, make sure the table does exist!!
	if p.table != "" {
		return p.table
	}
	return "task"
}

func Createtask(db *gorm.DB, user *Task) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Gettasks(db *gorm.DB, user *[]Task) (err error) {
	err = db.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Gettask(db *gorm.DB, user *Task, usersId string) (err error) {
	err = db.Where("id_task = ?", usersId).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Updatetask(db *gorm.DB, user *Task) (err error) {
	db.Save(user)
	return nil
}

func Deletetask(db *gorm.DB, user *Task, usersId string) (err error) {
	db.Where("id_task = ?", usersId).Delete(user)
	return nil
}
