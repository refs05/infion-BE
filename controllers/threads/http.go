package threads

import (
	"errors"
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

	return controller.ThreadsuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ThreadsController) ReadID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, err := strconv.Atoi(idstr)

	resp, err := ctrl.threadsUseCase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.ThreadsuccessResponse(c, response.FromDomain(resp))
}

// func (ctrl *ThreadsController) Update(c echo.Context) error {
// 	ctx := c.Request().Context()

// 	id := c.QueryParam("id")
// 	if strings.TrimSpace(id) == "" {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
// 	}

// 	req := request.Threads{}
// 	if err := c.Bind(&req); err != nil {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	domainReq := req.ToDomain()
// 	idInt, _ := strconv.Atoi(id)
// 	domainReq.ID = idInt
// 	resp, err := ctrl.threadsUseCase.Update(ctx, domainReq)
// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controller.ThreadsuccessResponse(c, response.FromDomain(*resp))
// }