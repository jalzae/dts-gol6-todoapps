package models

import (
	"gorm.io/gorm"
)

type Employe struct {
	Id_employe    int    `form:"id_employe" json:"id_employe" xml:"id_employe" gorm:"column:id_employe;primary_key;auto_increment;not null;"`
	Name_employe  string `form:"name_employe" json:"name_employe" xml:"name_employe" binding:"required" gorm:"column:name_employe;type:varchar(200);not null;"`
	Status_delete int    `form:"status_delete" json:"status_delete" xml:"status_delete" gorm:"column:status_delete;type:int(11);"`
	table         string `gorm:"-"`
}

func (p Employe) TableName() string {
	// double check here, make sure the table does exist!!
	if p.table != "" {
		return p.table
	}
	return "employe"
}

func Createemploye(db *gorm.DB, user *Employe) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Getemployes(db *gorm.DB, user *[]Employe) (err error) {
	err = db.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Getemploye(db *gorm.DB, user *Employe, usersId string) (err error) {
	err = db.Where("id_employe = ?", usersId).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Updateemploye(db *gorm.DB, user *Employe) (err error) {
	db.Save(user)
	return nil
}

func Deleteemploye(db *gorm.DB, user *Employe, usersId string) (err error) {
	db.Where("id_employe = ?", usersId).Delete(user)
	return nil
}
