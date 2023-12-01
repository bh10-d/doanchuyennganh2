package handlers

// "github.com/HelloWorld/goProductAPI/entity"

// // GetProductHandler is used to get data inside the products defined on our product catalog
// func GetProductHandler() http.HandlerFunc {
// 	return func(rw http.ResponseWriter, r *http.Request) {
// 		// Read JSON file
// 		data, err := os.ReadFile("../entity/data.json")
// 		if err != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		// // Write the body with JSON data
// 		rw.Header().Add("content-type", "application/json")
// 		rw.WriteHeader(http.StatusFound)
// 		// testdata := "buiduchieu"
// 		rw.Write([]byte(data))
// 		// fmt.Println(string(testdata))
// 	}
// }

// // CreateProductHandler is used to create a new product and add to our product store.
// func CreateProductHandler() http.HandlerFunc {
// 	return func(rw http.ResponseWriter, r *http.Request) {
// 		// Read incoming JSON from request body
// 		fmt.Println(r.Body)
// 		data, err := ioutil.ReadAll(r.Body)
// 		fmt.Println(data)
// 		// If no body is associated return with StatusBadRequest
// 		if err != nil {
// 			rw.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		// Check if data is proper JSON (data validation)
// 		var product entity.Product
// 		err = json.Unmarshal(data, &product)
// 		if err != nil {
// 			rw.WriteHeader(http.StatusExpectationFailed)
// 			rw.Write([]byte("Invalid Data Format"))
// 			return
// 		}
// 		// Load existing products and append the data to product list
// 		var products []entity.Product
// 		data, err = os.ReadFile("../entity/data.json")
// 		if err != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		// Load our JSON file to memory using array of products
// 		err = json.Unmarshal(data, &products)
// 		if err != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		// Add new Product to our list
// 		products = append(products, product)

// 		// Write Updated JSON file
// 		updatedData, err := json.Marshal(products)
// 		if err != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		err = os.WriteFile("../entity/data.json", updatedData, os.ModePerm)
// 		if err != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		// return after writing Body
// 		rw.WriteHeader(http.StatusCreated)
// 		rw.Write([]byte("Added New Product"))
// 	}
// }
