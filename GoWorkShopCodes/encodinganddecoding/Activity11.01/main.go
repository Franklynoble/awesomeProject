package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type customer struct {
	UserName      string  `json:"username"`
	Password      string  `json:"password"`
	Token         string  `json:"token"`
	ShipTo        address `json:"shipto"`
	PurchaseOrder order   `json:"order"`
}
type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}
type order struct {
	TotalPrice   int    `json:"totalprice"`
	IsPaid       bool   `json:"ispaid"`
	Fragile      bool   `json:"fragile,omitempty"`
	OrderDetail []item `json:"orderdetails"`
}
type item struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

func main() {

	jsonData := []byte(`
{
    "username":"James",
      "password":"1234",
      "token":"123456789",
       "shipto":
{  
"street":"Sulphur Springs Rd",
  "city":"Park City",
  "state":"VA",
   "zipcode":1122
},
"order":
{
 "totalprice": 475,
"ispaid":false,
 "fragile":true,
  "orderdetails":
[{"name":"A Guide to the World of Zeros and Ones",
 				 "description":"Book",
 					"quantity":3,
                    "price": 50
}]
}
}`)
	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
		os.Exit(1)
	}
	var c customer
	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	game := item{}
	game.Name = "Final Fantasy The Zodiac Age"
	game.Description = "Nintendo Switch Game"
	game.Quantity = 1
	game.Price = 50

	glass := item{}
	glass.Name = "Crystal Drinking Glass"
	glass.Quantity = 11
	glass.Price = 25


	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, game)
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, glass)
	c.Total()
	c.PurchaseOrder.IsPaid = true
	c.PurchaseOrder.Fragile = true

	customerOrder, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(customerOrder))

}

func (c *customer) Total() {
	price := 0
	for _, item := range c.PurchaseOrder.OrderDetail {
		price += item.Quantity * item.Price
	}
	c.PurchaseOrder.TotalPrice = price
}

