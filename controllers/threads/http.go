package threads

import (
	"errors"
	"fmt"
	"infion-BE/businesses/threads"
	controller "infion-BE/controllers"
	"infion-BE/controllers/threads/request"
	"infion-BE/controllers/threads/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type ThreadsController struct {
	threadsUseCase threads.Usecase
}

func NewThreadsController(threadsUC threads.Usecase) *ThreadsController {
	return &ThreadsController{
		threadsUseCase: threadsUC,
	}
}

func (ctrl *ThreadsController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Threads{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.threadsUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ThreadsController) ReadID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	resp, err := ctrl.threadsUseCase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ThreadsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Threads{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.threadsUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *ThreadsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Threads{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.threadsUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewDeleteResponse(c, response.FromDomain(*resp))
}

func (ctrl *ThreadsController) GetThreads(c echo.Context) error {
	fmt.Println("routes")
	ctx := c.Request().Context()
	sortBy := c.QueryParam("sortBy")
	category := c.QueryParam("category")

	if sortBy == "" && category == "" {
		fmt.Println("empty sortBy & category")
		threads, err := ctrl.threadsUseCase.GetThreads(ctx)
		if err != nil {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	
		return controller.NewSuccessResponse(c, response.NewResponseArray(threads))
	}

	if sortBy != "" && category == "" {
		fmt.Println("empty category")
		threads, err := ctrl.threadsUseCase.GetThreadsBySort(ctx, sortBy)
		if err != nil {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	
		return controller.NewSuccessResponse(c, response.NewResponseArray(threads))
	}

	if sortBy == "" && category != "" {
		fmt.Println("empty sortBy")
		threads, err := ctrl.threadsUseCase.GetThreadsByCategory(ctx, category)
		if err != nil {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	
		return controller.NewSuccessResponse(c, response.NewResponseArray(threads))
	}

	threads, err := ctrl.threadsUseCase.GetThreadsBySortCategory(ctx, sortBy, category)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	return controller.NewSuccessResponse(c, response.NewResponseArray(threads))
}

func (ctrl *ThreadsController) GetThreadsByUserID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	threads, err := ctrl.threadsUseCase.GetThreadsByUserID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	return controller.NewSuccessResponse(c, response.NewResponseArray(threads))
}