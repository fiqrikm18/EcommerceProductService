package presenter

import (
	"ecommerce/internal/domain/brand/dto"
	"ecommerce/internal/domain/brand/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	HttpResponser "ecommerce/pkg/response"
)

type IBrandPresenter interface {
	GetAll(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type BrandPresenter struct {
	useCase usecase.IBrandUseCase
}

func NewBrandPresenter(useCase usecase.IBrandUseCase) *BrandPresenter {
	return &BrandPresenter{
		useCase: useCase,
	}
}

// GetAll godoc
// @Summary      Get All brand
// @Description  Get All brand data
// @Tags         brand
// @Accept       json
// @Produce      json
// @Param 		 PerPage query int true "item per page count"
// @Param 		 Page query int true "page"
// @Param 		 Sort query string true "sorting order (desc, asc)"
// @Param 		 SortBy query string true "sorting fields (default created_at)"
// @Param 		 Search query string false "brand param query"
// @Success      200  {object}  response.PaginationResponse{data=[]dto.FindBrandDTO}
// @Router       /brands [get]
func (presenter *BrandPresenter) GetAll(c echo.Context) error {
	params := &dto.BrandPaginationDTO{}
	perPageParam := c.QueryParam("PerPage")
	pageParam := c.QueryParam("Page")
	sortParam := c.QueryParam("Sort")
	sortByParam := c.QueryParam("SortBy")
	searchParam := c.QueryParam("Search")

	perPage, err := strconv.ParseInt(perPageParam, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	params.Sort = sortParam
	params.SortBy = sortByParam
	params.Search = searchParam
	params.PerPage = perPage
	params.Page = page

	if err := c.Validate(params); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(
			"Bad Request",
			err.(*echo.HTTPError).Message.(map[string]interface{})["errors"]),
		)
	}

	count, totalPage, brands, err := presenter.useCase.FindAll(params)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, HttpResponser.NewPaginationResponse(count, totalPage, int(params.PerPage), int(params.Page), brands))
}

// Get godoc
// @Summary      Get brand
// @Description  Get brand data
// @Tags         brand
// @Accept       json
// @Produce      json
// @Param 		 id path int true "brand id"
// @Success      200  {object}  response.SuccessResponse{data=dto.FindBrandDTO}
// @Router       /brands/{id} [get]
func (presenter *BrandPresenter) Get(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	payload := &dto.BrandWithIdDTO{
		ID: id,
	}

	brand, err := presenter.useCase.FindById(payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, HttpResponser.NewSuccessResponse("Get brand success", brand))
}

// Create godoc
// @Summary      Create brand
// @Description  Create new brand data
// @Tags         brand
// @Accept       json
// @Produce      json
// @Param 		 request body dto.CreateBrandDTO true "request body"
// @Success      200  {object}  response.SuccessResponse{data=nil}
// @Router       /brands [post]
func (presenter *BrandPresenter) Create(c echo.Context) error {
	payload := dto.CreateBrandDTO{}
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	err := c.Validate(&payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(
			"Bad Request",
			err.(*echo.HTTPError).Message.(map[string]interface{})["errors"]),
		)
	}

	err = presenter.useCase.CreateBrand(&payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, HttpResponser.NewSuccessResponse("Brand created", nil))
}

// Update godoc
// @Summary      Update brand
// @Description  Update brand data
// @Tags         brand
// @Accept       json
// @Produce      json
// @Param 		 id path int true "brand id"
// @Param 		 request body dto.CreateBrandDTO true "request body"
// @Success      200  {object}  response.SuccessResponse{data=nil}
// @Router       /brands/{id} [patch]
func (presenter *BrandPresenter) Update(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	payload := dto.UpdateBrandDTO{}
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	payload.ID = id

	if err := c.Validate(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(
			"Bad Request",
			err.(*echo.HTTPError).Message.(map[string]interface{})["errors"]),
		)
	}

	if err := presenter.useCase.UpdateBrand(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, HttpResponser.NewSuccessResponse("Brand updated", nil))
}

// Delete godoc
// @Summary      Delete brand
// @Description  Delete brand data
// @Tags         brand
// @Accept       json
// @Produce      json
// @Param 		 id path int true "brand id"
// @Success      200  {object}  response.SuccessResponse{data=nil}
// @Router       /brands/{id} [delete]
func (presenter *BrandPresenter) Delete(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	payload := dto.BrandWithIdDTO{
		ID: id,
	}

	if err := c.Validate(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(
			"Bad Request",
			err.(*echo.HTTPError).Message.(map[string]interface{})["errors"]),
		)
	}

	if err := presenter.useCase.DeleteBrand(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, HttpResponser.NewSuccessResponse("Brand deleted", nil))
}
