package repository

import (
	"errors"
	"hotel-booking/internal/domain"
	"log"

	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryById(id int) (*domain.Category, error)
	EditCategory(e *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error

	CreateRoom(e *domain.Room) error
	FindRooms() ([]*domain.Room, error)
	FindRoomById(id int) (*domain.Room, error)
	FindLessorRooms(id int) ([]*domain.Room, error)
	EditRoom(e *domain.Room) (*domain.Room, error)
	DeleteRoom(e *domain.Room) error
}

type catalogRepository struct {
	db *gorm.DB
}

func (c catalogRepository) CreateRoom(e *domain.Room) error {
	err := c.db.Model(&domain.Room{}).Create(e).Error
	if err != nil {
		log.Printf("err: %v", err)
		return errors.New("cannot create room")
	}
	return nil
}

func (c catalogRepository) FindRooms() ([]*domain.Room, error) {
	var rooms []*domain.Room
	err := c.db.Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (c catalogRepository) FindRoomById(id int) (*domain.Room, error) {
	var room *domain.Room
	err := c.db.First(&room, id).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("room does not exist")
	}
	return room, nil
}

func (c catalogRepository) FindLessorRooms(id int) ([]*domain.Room, error) {
	var rooms []*domain.Room
	err := c.db.Where("user_id=?", id).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (c catalogRepository) EditRoom(e *domain.Room) (*domain.Room, error) {
	err := c.db.Save(&e).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("fail to update room")
	}
	return e, nil
}

func (c catalogRepository) DeleteRoom(e *domain.Room) error {
	err := c.db.Delete(&domain.Room{}, e.ID).Error
	if err != nil {
		return errors.New("room cannot delete")
	}
	return nil
}

func (c catalogRepository) CreateCategory(e *domain.Category) error {
	err := c.db.Create(&e).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("create category failed")
	}
	return nil
}

func (c catalogRepository) FindCategories() ([]*domain.Category, error) {
	var categories []*domain.Category

	err := c.db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c catalogRepository) FindCategoryById(id int) (*domain.Category, error) {
	var category *domain.Category

	err := c.db.First(&category, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("category does not exist")
	}

	return category, nil
}

func (c catalogRepository) EditCategory(e *domain.Category) (*domain.Category, error) {
	err := c.db.Save(&e).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("fail to update category")
	}

	return e, nil
}

func (c catalogRepository) DeleteCategory(id int) error {

	err := c.db.Delete(&domain.Category{}, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("fail to delete category")
	}

	return nil
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}
