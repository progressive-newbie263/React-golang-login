package controllers

import (
	"server/database"

	"time"

	"strconv"

	"server/models"

	"github.com/gofiber/fiber/v2"

	"golang.org/x/crypto/bcrypt"

	//notice: The guidance vid give dgrijalva as jwt package.
	//Should try to use golang-jwt/jwt instead.
	"github.com/dgrijalva/jwt-go" //note: don't add /v4 in the end.
)

const SecretKey = "secret"


func Register(c *fiber.Ctx) error{
	var data map[string] string

	err := c.BodyParser(&data)
	if err!= nil {
    return err
  }

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
    Email:    data["email"],
    Password: password,
	}

	/*
		note: successfully retrieved data with hashed password:
		{
    	"ID": 0,
    	"Name": "name",
    	"Email": "a@a.com",
    	"Password": "JDJhJDE0JGlsUWUzUmFrV01QQXFZTkQxMHI5M2VPamdZVUJlaWJibVk4elFWY0tQN21CcTdoc3gydmR5"
		}
	*/

	// Raw SQL query to insert user data
	query := `
		INSERT INTO users (name, email, password) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`

	// Execute the query and retrieve the new user's ID
	var userID int
	err = database.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to insert user: "+err.Error())
	}

	// Return the user information along with the new ID
	user.ID = userID

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	query := `SELECT id, name, email, password FROM users WHERE email = $1 LIMIT 1`

	// Execute the query and scan the result into the user struct
	err := database.DB.QueryRow(query, data["email"]).Scan(
		&user.ID, 
		&user.Name, 
		&user.Email, 
		&user.Password)
	
	if err != nil {
		// User not found in the database
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	
	/*
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
    c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map {
			"message": "User not found",
		})}
	*/

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
      "message": "Incorrect password",
    })
	}

	//jwt token part:
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
      "message": "Could not login",
    })
	}

	//get json cookie part:
	cookie := fiber.Cookie {
		Name : "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24), //1 day
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	//return c.JSON(user) //post user login/reg, pre jwt
	//return c.JSON(token) //post jwt, pre cookie
		/*
	this is the response from jwt.io, after cracking 64-bit string:
	{
		"exp": 1728718561, //issuer : 4
		"iss": "4"
	}

	{
		"exp": 1728718611, //issuer: 7
		"iss": "7"
	}*/ 
	return c.JSON(fiber.Map{
		"message": "success",
	})
}


func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
      "message": "Could not login",
    })
	}

	//note: Claims is an interface, with Valid() error //claims.go, line 11
	//claims := token.Claims 

	claims := token.Claims.(*jwt.StandardClaims) //made jwt.StandardClaims a pointer (interface conversion).

	var user models.User

	query := `SELECT id, name, email, password FROM users WHERE id = $1 LIMIT 1`

	// Execute the query and scan the result into the user struct
	// with this code, at the moment this will return a bcrypted string password.
	err = database.DB.QueryRow(query, claims.Issuer).Scan(
		&user.ID, 
		&user.Name, 
		&user.Email,
		&user.Password)
	
	if err != nil {
		// User not found in the database
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user) //return: {"id": 7,"name": "name","email": "aaa@a.com", "password": "JDJhJDE0JHAxYnU2cFRhRzUzZ2U5R2NyaW9EM2VTd2V6RzNzUUpiZkJJNG5OR2hxVHpDcXVNYkhuWkRh"}
	//return c.JSON(claims) //return: {"exp": 1728718611, "iss": "7"}
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
    Name: "jwt",
    Value: "",
		// minus -time.Hour means that cookie expires an hour ago.
    Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
  }

  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "message": "Logged out",
  })

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}