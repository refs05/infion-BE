package followUsers

import (
	"errors"
	"infion-BE/businesses/followUsers"
	controller "infion-BE/controllers"
	"infion-BE/controllers/followUsers/request"
	"infion-BE/controllers/followUsers/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type FollowUsersController struct {
	followUsersUseCase followUsers.Usecase
}

func NewFollowUsersController(followUsersUC followUsers.Usecase) *FollowUsersController {
	return &FollowUsersController{
		followUsersUseCase: followUsersUC,
	}
}

func (ctrl *FollowUsersController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.FollowUsers{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.followUsersUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *FollowUsersController) ReadID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	resp, err := ctrl.followUsersUseCase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *FollowUsersController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.FollowUsers{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.followUsersUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *FollowUsersController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.FollowUsers{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.followUsersUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewDeleteResponse(c, response.FromDomain(*resp))
}