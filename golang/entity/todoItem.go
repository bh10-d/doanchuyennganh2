package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ToDoItem struct {
	Id        int64      `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

var HOST_NODE_SERVER = "172.26.0.4"

// var HOST_NODE_SERVER = "localhost"

// Define func TableName belong to TodoItem struct
func (ToDoItem) TableName() string { return "todo_items" }

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem ToDoItem

		fmt.Println(dataItem)

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)
		fmt.Println(dataItem.Title)
		if dataItem.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
			return
		}

		// do not allow "finished" status when creating a new task
		dataItem.Status = "Doing" // set to default

		if err := db.Create(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func GetListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []ToDoItem

		if err := db.Table(ToDoItem{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func ReadItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem ToDoItem

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem})
	}
}

func EditItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func DeleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(ToDoItem{}.TableName()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func GetData() gin.HandlerFunc {
	// HOST_NODE_SERVER := "172.26.0.4"
	return func(c *gin.Context) {
		// data := "buiduchieu do an chuyen nganh 2"
		// c.JSON(http.StatusOK, gin.H{"data": data})
		fmt.Print("buiduchieu da o day\n")
		API := "http://" + HOST_NODE_SERVER + ":9090/api/test"
		// Make an HTTP GET request to the Node.js API
		resp, err := http.Get(API)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse the JSON response body
		var data interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Send the parsed data as JSON response
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func DeleteData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println("This function is working")
		// fmt.Println(c.Param("id"))
		// c.JSON(http.StatusOK, gin.H{"data": c.Param("id")})
		API := "http://" + HOST_NODE_SERVER + ":9090/api/test/" + c.Param("id")
		fmt.Println(API)
		// Create a new HTTP request object
		req, err := http.NewRequest("DELETE", API, nil)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Send the DELETE request
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse the JSON response body
		var data interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check the response status code
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error:", resp.StatusCode)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
		// Print a success message
		// fmt.Println("Data deleted successfully")
	}
}
