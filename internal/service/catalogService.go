package service

import (
	"errors"
	"hotel-booking/config"
	"hotel-booking/internal/domain"
	"hotel-booking/internal/dto"
	"hotel-booking/internal/helper"
	"hotel-booking/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {
	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})
	return err
}

func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {
	exitCat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}

	if len(input.Name) > 0 {
		exitCat.Name = input.Name
	}

	if input.ParentId > 0 {
		exitCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		exitCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		exitCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(exitCat)
	return updatedCat, err
}

func (s CatalogService) DeleteCategory(id int) error {
	err := s.Repo.DeleteCategory(id)
	if err != nil {
		return errors.New("category does not exist to delete")
	}
	return nil
}

func (s CatalogService) GetCategories() ([]*domain.Category, error) {
	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, errors.New("categories do not exist")
	}
	return categories, err
}

func (s CatalogService) GetCategory(id int) (*domain.Category, error) {
	cat, err := s.Repo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}
	return cat, nil
}

func (s CatalogService) CreateRoom(input dto.CreateRoomRequest, user domain.User) error {
	err := s.Repo.CreateRoom(&domain.Room{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
		ImageUrl:    input.ImageUrl,
		UserId:      int(user.ID),
		Stock:       uint(input.Stock),
	})
	return err
}

func (s CatalogService) EditRoom(id int, input dto.CreateRoomRequest, user domain.User) (*domain.Room, error) {
	existRoom, err := s.Repo.FindRoomById(id)
	if err != nil {
		return nil, errors.New("room does not exist")
	}

	if existRoom.UserId != int(user.ID) {
		return nil, errors.New("you don't have manage rights of this room")
	}

	if len(input.Name) > 0 {
		existRoom.Name = input.Name
	}
	if len(input.Description) > 0 {
		existRoom.Description = input.Description
	}
	if input.Price > 0 {
		existRoom.Price = input.Price
	}
	if input.CategoryId > 0 {
		existRoom.CategoryId = input.CategoryId
	}

	updatedRoom, err := s.Repo.EditRoom(existRoom)
	return updatedRoom, err
}

func (s CatalogService) DeleteRoom(id int, user domain.User) error {
	existRoom, err := s.Repo.FindRoomById(id)
	if err != nil {
		return errors.New("room does not exist")
	}
	if existRoom.UserId != int(user.ID) {
		return errors.New("you don't have manage rights of this room")
	}
	err = s.Repo.DeleteRoom(existRoom)
	if err != nil {
		return errors.New("room cannot be deleted")
	}
	return nil
}

func (s CatalogService) GetRooms() ([]*domain.Room, error) {
	rooms, err := s.Repo.FindRooms()
	if err != nil {
		return nil, errors.New("rooms do not exist")
	}
	return rooms, err
}

func (s CatalogService) GetRoomById(id int) (*domain.Room, error) {
	room, err := s.Repo.FindRoomById(id)
	if err != nil {
		return nil, errors.New("room does not exist")
	}
	return room, nil
}

func (s CatalogService) GetSellerRooms(id int) ([]*domain.Room, error) {
	rooms, err := s.Repo.FindLessorRooms(id)
	if err != nil {
		return nil, errors.New("rooms do not exist")
	}
	return rooms, err
}

func (s CatalogService) UpdateRoomStock(e domain.Room) (*domain.Room, error) {
	room, err := s.Repo.FindRoomById(int(e.ID))
	if err != nil {
		return nil, errors.New("room not found")
	}

	if room.UserId != e.UserId {
		return nil, errors.New("you don't have manage rights of this room")
	}
	room.Stock = e.Stock
	editRoom, err := s.Repo.EditRoom(room)
	if err != nil {
		return nil, err
	}
	return editRoom, nil
}
