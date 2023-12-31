package handlers

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"golang/entity"

	"github.com/gorilla/mux"
	// "github.com/HelloWorld/goProductAPI/entity"
)

// GetProductsHandler is used to get data inside the products defined on our product catalog
func GetProductsHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := entity.GetProducts()
		// fmt.Println(string(data))
		// fmt.Println("Bui duc hieu")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Write the body with JSON data
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(data)
	}
}

// GetProductHandler is used to get data inside the products defined on our product catalog
func GetProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read product ID
		productID := mux.Vars(r)["id"]
		product, err := entity.GetProduct(productID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		responseData, err := json.Marshal(product)
		if err != nil {
			// Check if it is No product error or any other error
			if errors.Is(err, entity.ErrNoProduct) {
				// Write Header if no related product found.
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Write body with found product
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(responseData)
	}
}

// CreateProductHandler is used to create a new product and add to our product store.
func CreateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read incoming JSON from request body
		data, err := ioutil.ReadAll(r.Body)
		// If no body is associated return with StatusBadRequest
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var product entity.Product
		err = json.Unmarshal(data, &product)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}
		err = entity.AddProduct(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// return after writing Body
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte("Added New Product Successful"))
	}
}

// DeleteProductHandler deletes the product with given ID.
func DeleteProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read product ID
		productID := mux.Vars(r)["id"]
		err := entity.DeleteProduct(productID)
		if err != nil {
			// Check if it is No product error or any other error
			if errors.Is(err, entity.ErrNoProduct) {
				// Write Header if no related product found.
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Write Header with Accepted Status (done operation)
		rw.WriteHeader(http.StatusAccepted)
	}
}

// UpdateProductHandler deletes the product with given ID.
func UpdateProductHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Read product ID
		productID := mux.Vars(r)["id"]
		err := entity.DeleteProduct(productID)
		if err != nil {
			if errors.Is(err, entity.ErrNoProduct) {
				rw.WriteHeader(http.StatusNoContent)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		// Read incoming JSON from request body
		data, err := ioutil.ReadAll(r.Body)
		// If no body is associated return with StatusBadRequest
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if data is proper JSON (data validation)
		var product entity.Product
		err = json.Unmarshal(data, &product)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid Data Format"))
			return
		}
		// Addproduct with the requested body
		err = entity.AddProduct(product)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Write Header if no related product found.
		rw.WriteHeader(http.StatusAccepted)
	}
}

func AuthHandler(h http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if ok {
			username := sha256.Sum256([]byte("admin"))
			password := sha256.Sum256([]byte("admin"))
			userHash := sha256.Sum256([]byte(user))
			passHash := sha256.Sum256([]byte(pass))
			validUser := subtle.ConstantTimeCompare(userHash[:], username[:]) == 1
			validPass := subtle.ConstantTimeCompare(passHash[:], password[:]) == 1
			if validPass && validUser {
				h.ServeHTTP(rw, r)
				return
			}
		}
		http.Error(rw, "No/Invalid Credentials", http.StatusUnauthorized)
	}
}
