package handlers

import (
	"hotel-booking/internal/api/rest"
	"hotel-booking/internal/domain"
	"hotel-booking/internal/dto"
	"hotel-booking/internal/repository"
	"hotel-booking/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {

	app := rh.App

	// create an instance of user service & inject to handler
	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := CatalogHandler{
		svc: svc,
	}

	// public
	// listing rooms and categories
	app.Get("/rooms", handler.GetRooms)
	app.Get("/rooms/:id", handler.GetRoom)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryById)

	// private
	// manage rooms and categories
	selRoutes := app.Group("/lessor", rh.Auth.AuthorizeLessor)
	// Categories
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// Rooms
	selRoutes.Post("/rooms", handler.CreateRooms)
	selRoutes.Get("/rooms", handler.GetRooms)
	selRoutes.Get("/rooms/:id", handler.GetRoom)
	selRoutes.Put("/rooms/:id", handler.EditRoom)
	selRoutes.Patch("/rooms/:id", handler.UpdateStock) // update stock
	selRoutes.Delete("/rooms/:id", handler.DeleteRoom)

}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {

	cats, err := h.svc.GetCategories()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.SuccessResponse(ctx, "categories", cats)
}
func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	cat, err := h.svc.GetCategory(id)
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.SuccessResponse(ctx, "category", cat)
}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {

	req := dto.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)

	if err != nil {
		return rest.BadRequestError(ctx, "create category request is not valid")
	}

	err = h.svc.CreateCategory(req)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "category created successfully", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)

	if err != nil {
		return rest.BadRequestError(ctx, "update category request is not valid")
	}

	updatedCat, err := h.svc.EditCategory(id, req)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "edit category", updatedCat)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	err := h.svc.DeleteCategory(id)
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "category deleted successfully", nil)
}

func (h CatalogHandler) CreateRooms(ctx *fiber.Ctx) error {

	req := dto.CreateRoomRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "create room request is not valid")
	}

	user := h.svc.Auth.GetCurrentUser(ctx)
	err = h.svc.CreateRoom(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "room created successfully", nil)
}

func (h CatalogHandler) GetRooms(ctx *fiber.Ctx) error {

	rooms, err := h.svc.GetRooms()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "rooms", rooms)
}

func (h CatalogHandler) GetRoom(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	room, err := h.svc.GetRoomById(id)
	if err != nil {
		return rest.BadRequestError(ctx, "room not found")
	}

	return rest.SuccessResponse(ctx, "room", room)
}

func (h CatalogHandler) EditRoom(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	req := dto.CreateRoomRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "edit room request is not valid")
	}
	user := h.svc.Auth.GetCurrentUser(ctx)
	room, err := h.svc.EditRoom(id, req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "edit room", room)
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	req := dto.UpdateStockRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "update stock request is not valid")
	}
	user := h.svc.Auth.GetCurrentUser(ctx)

	room := domain.Room{
		ID:     uint(id),
		Stock:  uint(req.Stock),
		UserId: int(user.ID),
	}

	updatedRoom, err := h.svc.UpdateRoomStock(room)

	return rest.SuccessResponse(ctx, "update stock ", updatedRoom)
}

func (h CatalogHandler) DeleteRoom(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	// need to provide user id to verify ownership
	user := h.svc.Auth.GetCurrentUser(ctx)
	err := h.svc.DeleteRoom(id, user)

	return rest.SuccessResponse(ctx, "Delete room ", err)
}
