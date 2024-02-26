package routes

import (
	"errors"
	"http-test/lib"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	name := c.QueryParam("name")
	rawAge := c.QueryParam("age")

	age, err := strconv.Atoi(rawAge)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	db, err := lib.GetDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user = lib.User{Name: name, Age: age}
	db.Create(&user)
	if db.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	db, err := lib.GetDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user *lib.User
	if err := db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No matching record found, user will be nil
			user = nil
		} else {
			// Other error occurred
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	rawID := c.Param("id")
	var (
		name string
		age  int
		err  error
	)

	id, err := strconv.Atoi(rawID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if queryName := c.QueryParam("name"); queryName != "" {
		name = queryName
	}
	if queryAge := c.QueryParam("age"); queryAge != "" {
		age, err = strconv.Atoi(queryAge)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	db, err := lib.GetDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user lib.User
	if err := db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	if name != "" {
		user.Name = name
	}
	if age != 0 {
		user.Age = age
	}

	db.Save(&user)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	db, err := lib.GetDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var user lib.User
	if err := db.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	db.Delete(&user)

	return c.JSON(http.StatusOK, "User deleted")
}
