package app

import "gorm.io/gorm"

type Repository interface {
	Save(app User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(app User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(app User) (User, error) {
	err := r.db.Create(&app).Error
	if err != nil {
		return app, err
	}
	return app, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ? ", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ? ", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(app User) (User, error) {
	err := r.db.Save(&app).Error
	if err != nil {
		return app, err
	}
	return app, nil
}
