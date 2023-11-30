package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/job-application/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers(context.Context) ([]model.User, error)
	FindUserById(context.Context, uint) (*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	NewUser(context.Context, model.User)(*model.User, error)
}

type userRepository struct{
	db *gorm.DB
}

func (u *userRepository) FindUsers(ctx context.Context) (users []model.User, err error) {
	err = u.db.WithContext(ctx).Table("users").Find(&users).Error
	if err != nil {
		return nil, apperror.ErrFindUsersQuery
	}
	return users, nil
}

func (u *userRepository) FindUserById(ctx context.Context, id uint) (user *model.User, err error) {
	result := u.db.WithContext(ctx).Table("users").Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, apperror.ErrFindUserByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (user *model.User, err error) {
	result := u.db.WithContext(ctx).Table("users").Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, apperror.ErrFindUserByEmail
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrUserNotFound
	}
	return user, nil
}

func (u *userRepository) NewUser(ctx context.Context, user model.User) (newUser *model.User, err error){
	err = u.db.WithContext(ctx).Table("users").Create(&user).Error
	if err != nil {
		return nil, apperror.ErrNewUserQuery
	}
	return &user, nil
}

