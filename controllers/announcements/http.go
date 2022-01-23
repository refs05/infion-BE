package announcements

import (
	"errors"
	"infion-BE/businesses/announcements"
	controller "infion-BE/controllers"
	"infion-BE/controllers/announcements/request"
	"infion-BE/controllers/announcements/response"
	"net/http"
	"strconv"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type AnnouncementsController struct {
	announcementsUseCase announcements.Usecase
}

func NewAnnouncementsController(announcementsUC announcements.Usecase) *AnnouncementsController {
	return &AnnouncementsController{
		announcementsUseCase: announcementsUC,
	}
}

func (ctrl *AnnouncementsController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Announcements{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.announcementsUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *AnnouncementsController) ReadID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	resp, err := ctrl.announcementsUseCase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *AnnouncementsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Announcements{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.announcementsUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}

func (ctrl *AnnouncementsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	req := request.Announcements{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = id
	resp, err := ctrl.announcementsUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewDeleteResponse(c, response.FromDomain(*resp))
}

func (ctrl *AnnouncementsController) GetAnnouncements(c echo.Context) error {
	ctx := c.Request().Context()
	
	announcements, err := ctrl.announcementsUseCase.GetAnnouncements(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.NewResponseArray(announcements))
}

func (ctrl *AnnouncementsController) GetAnnouncementsByUserID(c echo.Context) error {
	ctx := c.Request().Context()

	idstr := c.Param("id")
	if strings.TrimSpace(idstr) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	id, _ := strconv.Atoi(idstr)

	announcements, err := ctrl.announcementsUseCase.GetAnnouncementsByUserID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	return controller.NewSuccessResponse(c, response.NewResponseArray(announcements))
}