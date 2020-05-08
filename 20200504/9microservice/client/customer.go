package client

import (
	// "encoding/json"
	"fmt"
	"github.com/muly/go-training/20200504/9microservice/models"
	"io/ioutil"
	"net/http"
)

// jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
// jsonValue, _ := json.Marshal(jsonData)
// request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
// request.Header.Set("Content-Type", "application/json")
// client := &http.Client{}
// response, err := client.Do(request)

func GetCustomer(loc string) (models.Customer, error) {

	url := fmt.Sprintf("http://localhost:8080/customer?location=%s", loc)
	fmt.Println(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest error ", err)
		return models.Customer{}, err
	}

	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		fmt.Println("c.Do(r) error ", err)
		return models.Customer{}, err
	}

	fmt.Println(resp.Status)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll error ", err)
		return models.Customer{}, err
	}
	fmt.Println(string(b))

	cust := models.Customer{}
	// err = json.NewDecoder(resp.Body).Decode(&cust)
	// if err != nil {
	// 	fmt.Println("json.NewDecoder error ", err)
	// 	// b, _ := ioutil.ReadAll(resp.Body)
	// 	// fmt.Println(string(b))
	// 	return models.Customer{}, err
	// }

	return cust, nil

}
