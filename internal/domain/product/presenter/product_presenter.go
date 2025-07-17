package presenter

import (
	"ecommerce/internal/domain/product/dto"
	"ecommerce/internal/domain/product/usecase"
	HttpResponser "ecommerce/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type IProductPresenter interface {
	GetAll(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type ProductPresenter struct {
	useCase usecase.IProductUseCase
}

func NewProductPresenter(useCase usecase.IProductUseCase) *ProductPresenter {
	return &ProductPresenter{useCase}
}

// GetAll godoc
// @Summary      Get All product
// @Description  Get All product data
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 PerPage query int true "item per page count"
// @Param 		 Page query int true "page"
// @Param 		 Sort query string true "sorting order (desc, asc)"
// @Param 		 SortBy query string true "sorting fields (default created_at)"
// @Param 		 Search query string false "product param query"
// @Success      200  {object}  response.PaginationResponse{data=[]dto.FindProductDTO}
// @Router       /products [get]
func (p *ProductPresenter) GetAll(c echo.Context) error {
	params := &dto.ProductPaginationDTO{}
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

	count, totalPage, products, err := p.useCase.FindAll(params)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, HttpResponser.NewPaginationResponse(count, totalPage, int(params.PerPage), int(params.Page), products))
}

// Get godoc
// @Summary      Get product
// @Description  Get product data
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path int true "product id"
// @Success      200  {object}  response.PaginationResponse{data=dto.FindProductDTO}
// @Router       /products/{id} [get]
func (p *ProductPresenter) Get(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	payload := &dto.ProductWithIdDTO{
		ID: id,
	}

	product, err := p.useCase.FindById(payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, HttpResponser.NewSuccessResponse("Get product success", product))
}

// Create godoc
// @Summary      Create product
// @Description  Create product data
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 request body dto.CreateProductDTO true "request body"
// @Success      200  {object}  response.PaginationResponse{data=nil}
// @Router       /products [post]
func (p *ProductPresenter) Create(c echo.Context) error {
	payload := &dto.CreateProductDTO{}
	if err := c.Bind(&payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(
			"Bad Request",
			err.(*echo.HTTPError).Message.(map[string]interface{})["errors"]),
		)
	}

	err := p.useCase.CreateProduct(payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, HttpResponser.NewSuccessResponse("Product created", nil))
}

// Update godoc
// @Summary      Update product
// @Description  Update product data
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path int true "product id"
// @Param 		 request body dto.UpdateProductDTO true "request body"
// @Success      200  {object}  response.PaginationResponse{data=nil}
// @Router       /products/{id} [patch]
func (p *ProductPresenter) Update(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	payload := &dto.UpdateProductDTO{}
	if err := c.Bind(payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(err.Error()))
	}

	payload.ID = id

	if err := c.Validate(payload); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, HttpResponser.NewErrorResponse(
			"Bad Request",
			err.(*echo.HTTPError).Message.(map[string]interface{})["errors"]),
		)
	}

	err = p.useCase.UpdateProduct(payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, HttpResponser.NewSuccessResponse("Product updated", nil))
}

// Delete godoc
// @Summary      Delete product
// @Description  Delete product data
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id path int true "product id"
// @Success      200  {object}  response.PaginationResponse{data=nil}
// @Router       /products/{id} [delete]
func (p *ProductPresenter) Delete(c echo.Context) error {
	payload := &dto.ProductWithIdDTO{}
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.Logger().Error(err)
	}

	payload.ID = id

	err = p.useCase.DeleteProduct(payload)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, HttpResponser.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, HttpResponser.NewSuccessResponse("Product deleted", nil))
}
