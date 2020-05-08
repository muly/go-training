package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Item : Non discript data type
type Item struct {
	Title string
	Body  string
}

// API : we are declaring this bcz for RPC a func should satisfy 1)Methods 2)Return error 3)two arguments one being pointer
type API int

var dB []Item

// GetDB : grab database throug the client
func (a *API) GetDB(title string, reply *[]Item) error { // title is not necessary but rpc signature demands somethig
	*reply = dB
	return nil
}

// GetByName : read through dB and get the Items by their names/title
func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item
	for _, val := range dB {
		if val.Title == title {
			getItem = val
		}
	}
	*reply = getItem
	return nil
}

// AddItem  : takes in an item and pass an item
func (a *API) AddItem(item Item, reply *Item) error {
	dB = append(dB, item)
	*reply = item
	return nil
}

// EditItem : modify the item fields
func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for i, val := range dB {
		if val.Title == edit.Title {
			dB[i] = Item{edit.Title, edit.Body}
			changed = dB[i]
		}
	}
	*reply = changed
	return nil
}

// DeleteItem : deletes the item from dB
func (a *API) DeleteItem(del Item, reply *Item) error {
	var deleted Item
	for i, val := range dB {
		if val.Title == del.Title && val.Body == del.Body {
			dB = append(dB[:i], dB[i+1:]...)
			deleted = del
			break
		}
	}
	*reply = deleted
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatalln("error registering API :", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalln("listener error :", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalln("error serving :", err)
	} //

	/* testing database
	a := Item{Title: "first", Body: "first test item"}
	b := Item{Title: "second", Body: "second item"}
	c := Item{Title: "first", Body: "third test"}

	fmt.Println("Initial database:\n", dB)

	AddItem(a)
	AddItem(b)
	AddItem(c)

	fmt.Println("After adding items to database:\n", dB)

	EditItem("second", Item{"second", "second test item"})
	EditItem("third", Item{"third", "third test item"})

	fmt.Println("After editing items:\n", dB)

	DeleteItem(c)

	fmt.Println("After deleting one item:\n", dB)

	x := GetByName("first")
	fmt.Println("First item:\n", x) */
}
