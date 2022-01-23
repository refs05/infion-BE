package reports

import (
	"errors"
	"infion-BE/businesses/reports"
	controller "infion-BE/controllers"
	"infion-BE/controllers/reports/request"
	"infion-BE/controllers/reports/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type ReportsController struct {
	reportsUseCase reports.Usecase
}

func NewReportsController(reportsUC reports.Usecase) *ReportsController {
	return &ReportsController{
		reportsUseCase: reportsUC,
	}
}

func (ctrl *ReportsController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Reports{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.reportsUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ReportsController) ReadID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	resp, err := ctrl.reportsUseCase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *ReportsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Reports{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.reportsUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *ReportsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Reports{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.reportsUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewDeleteResponse(c, response.FromDomain(*resp))
}

func (ctrl *ReportsController) GetReports(c echo.Context) error {
	ctx := c.Request().Context()
	
	reports, err := ctrl.reportsUseCase.GetReports(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.NewResponseArray(reports))
}

func (ctrl *ReportsController) GetReportsByUserID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	reports, err := ctrl.reportsUseCase.GetReportsByUserID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	return controller.NewSuccessResponse(c, response.NewResponseArray(reports))
}