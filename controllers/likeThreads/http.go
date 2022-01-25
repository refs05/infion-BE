package likeThreads

import (
	"errors"
	"fmt"
	"infion-BE/businesses/likeThreads"
	controller "infion-BE/controllers"
	"infion-BE/controllers/likeThreads/request"
	"infion-BE/controllers/likeThreads/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type LikeThreadsController struct {
	likeThreadsUseCase likeThreads.Usecase
}

func NewLikeThreadsController(likeThreadsUC likeThreads.Usecase) *LikeThreadsController {
	return &LikeThreadsController{
		likeThreadsUseCase: likeThreadsUC,
	}
}

func (ctrl *LikeThreadsController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.LikeThreads{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.likeThreadsUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *LikeThreadsController) ReadID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	resp, err := ctrl.likeThreadsUseCase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *LikeThreadsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.LikeThreads{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.likeThreadsUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *LikeThreadsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.LikeThreads{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.likeThreadsUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewDeleteResponse(c, response.FromDomain(*resp))
}

func (ctrl *LikeThreadsController) GetStatus(c echo.Context) error {
	fmt.Println("routes")
	ctx := c.Request().Context()
	threadIDString := c.QueryParam("threadID")
	userIDString := c.QueryParam("userID")

	if threadIDString != "" && userIDString != "" {
		fmt.Println("good request query param")
		threadID, _ := strconv.Atoi(threadIDString)
		userID, _ := strconv.Atoi(userIDString)
		resp, err := ctrl.likeThreadsUseCase.GetStatus(ctx, threadID, userID)
		if err != nil {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
		return controller.NewSuccessResponse(c, response.FromDomain(resp))
	}
	
	return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required threadID or userID"))
}