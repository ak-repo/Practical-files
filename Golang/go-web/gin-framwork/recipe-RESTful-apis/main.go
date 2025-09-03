package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json: "id"`
	Name         string    `json: "name"`
	Tags         []string  `json: "tags"`
	Ingredients  []string  `json: "ingredients"`
	Instructions []string  `json: "instructions"`
	PublishedAt  time.Time `json: "publishedAt"`
}

// for recipe
var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

// POST method

func NewRecipeHamdler(c *gin.Context) {

	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)

}

// GET method

func ListRecipesHandler(c *gin.Context) {

	c.JSON(http.StatusOK, recipes)
}

// PUT method

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	// fmt.Println(c)
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	index := -1

	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return
	}

	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)

}

//DELETE method

func DeleteRecipeHandler(c *gin.Context) {

	id := c.Param("id")

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not founc"})
	}

	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "recipe has been deleted"})

}

//search

func SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)

	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])
		}
	}

	

	c.JSON(http.StatusOK, listOfRecipes)
}

func main() {

	router := gin.Default()

	//routes
	router.POST("/recipes", NewRecipeHamdler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.GET("/recipes/search",SearchRecipesHandler)

	router.Run()
}
