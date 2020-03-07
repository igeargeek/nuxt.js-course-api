package controllers

import (
	"net/http"

	"app/src/models"
	"app/src/utils"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

type ReservationHandler struct {
	MovieService       models.MovieReporer
	ReservationService models.ReservationReporer
	Validator          ut.Translator
}

func NewReservationHandler(movieRepository models.MovieReporer, reservationRepository models.ReservationReporer, validator ut.Translator) ReservationHandler {
	return ReservationHandler{
		MovieService:       movieRepository,
		ReservationService: reservationRepository,
		Validator:          validator,
	}
}

func (handler *ReservationHandler) ShowOneReservationGet(c *gin.Context) {
	id := c.Param("id")
	userId := "123" // mock user id
	reservation, err := handler.ReservationService.GetID(id, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseServerError("Not found!"))
		return
	}

	_, err = handler.MovieService.GetID(reservation.MovieId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ResponseServerError("Not found!"))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseObject(reservation))
}

func (handler *ReservationHandler) ShowAllReservationGet(c *gin.Context) {
	userId := "123" // mock user id
	reservations, _ := handler.ReservationService.GetAll(userId)
	c.JSON(http.StatusOK, utils.ResponseObject(gin.H{
		"message": "Data retrieval successfully",
		"total":   len(reservations),
		"data":    reservations,
	}))
}

func (handler *ReservationHandler) CreateReservationPost(c *gin.Context) {
	var reservation models.Reservation
	reservation.MovieId = c.Param("id")
	reservation.UserId = "123" // mock user id
	if err := c.ShouldBind(&reservation); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseErrorValidation(handler.Validator, err))
		return
	}
	movie, err := handler.MovieService.GetID(reservation.MovieId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	if len(utils.Difference(movie.ReservedSeat, reservation.SeatNo)) != len(movie.ReservedSeat) {
		c.JSON(http.StatusBadRequest, utils.ResponseObject(gin.H{
			"message": "This seat has been selected!",
		}))
		return
	}

	id, err := handler.ReservationService.Create(&reservation, &movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	movie.ReservedSeat = append(movie.ReservedSeat, reservation.SeatNo...)
	errReserved := handler.MovieService.Edit(reservation.MovieId, &movie)

	if errReserved != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseServerError("Something went wrong."))
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseObject(gin.H{
		"message": "Created Successful",
		"id":      id,
	}))
}
