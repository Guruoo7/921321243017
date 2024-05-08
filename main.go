package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
)

type Product struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Rating   float64 `json:"rating"`
    Company  string  `json:"company"`
    Discount float64 `json:"discount"`
}

type ProductsResponse struct {
    Products []Product `json:"products"`
    Page     int       `json:"page"`
}

func getproduct(w http.ResponseWriter, r *http.Request) {
    category := r.URL.Path[len("/categories/"):]
    n, _ := strconv.Atoi(r.URL.Query().Get("n"))
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    products, err := fetchProducts(category, n, page)
    if err != nil {
        http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
        return
    }
    jsonResponse, err := json.Marshal(products)
    if err != nil {
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}


func fetchapiproducts(category string, n, page int) ([]Product, error) {
    mockProducts := []Product{
        {ID: "1", Name: "Product 1", Price: 50.0, Rating: 4.5, Company: "Company A", Discount: 10.0},
        {ID: "2", Name: "Product 2", Price: 30.0, Rating: 4.2, Company: "Company B", Discount: 5.0},
    }
    return mockProducts, nil
}


func fetchProducts(category string, n, page int) (ProductsResponse, error) {
    products, err := fetchapiproducts(category, n, page)
    if err != nil {
        return ProductsResponse{}, err
    }
    
    startIndex := page * n
    endIndex := (page + 1) * n
    if endIndex > len(products) {
        endIndex = len(products)
    }

    return ProductsResponse{
        Products: products[startIndex:endIndex],
        Page:     page,
    }, nil
}

func getProductDetails(w http.ResponseWriter, r *http.Request) {
    category := r.URL.Path[len("/categories/"):]
    productID := r.URL.Path[len("/categories/"+category+"/products/"):]
    product := fetchProductDetails(productID)
    jsonResponse, err := json.Marshal(product)
    if err != nil {
        http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func fetchProductDetails(productID string) Product {
    return Product{
        ID:       productID,
        Name:     "Product " + productID,
        Price:    50.0,
        Rating:   4.5,
        Company:  "Company A",
        Discount: 10.0,
    }
}

func main() {
    http.HandleFunc("/categories/", getproduct)
    http.HandleFunc("/categories/:categoryname/products/:productid", getProductDetails)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
