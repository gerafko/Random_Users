package randomuser

import "errors"

// ErrNoUsers представляет ошибку отсутствия юзеров
var ErrNoUsers = errors.New("no appointments")

type GetRandomUserResponse struct {
	Results []User `json:"results"`
}

type User struct {
	Gender     string     `json:"gender"`
	Name       Name       `json:"name"`
	Location   Location   `json:"location"`
	Email      string     `json:"email"`
	Login      Login      `json:"login"`
	Dob        Dob        `json:"dob"`
	Registered Registered `json:"registered"`
	Phone      string     `json:"phone"`
	Cell       string     `json:"cell"`
	ID         ID         `json:"id"`
	Picture    Picture    `json:"picture"`
	Nat        string     `json:"nat"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Location struct {
	Street      Street      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    int         `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MD5      string `json:"md5"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
}

type Dob struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type Registered struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type ID struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}
