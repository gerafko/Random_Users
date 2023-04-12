package model

import "test/pkg/randomuser"

type User struct {
	ID                     int    `db:"id"`
	Gender                 string `db:"gender"`
	Email                  string `db:"email"`
	Phone                  string `db:"phone"`
	Cell                   string `db:"cell"`
	Nat                    string `db:"nat"`
	NameTitle              string `db:"name_title"`
	NameFirst              string `db:"name_first"`
	NameLast               string `db:"name_last"`
	LocationStreetNumber   int    `db:"location_street_number"`
	LocationStreetName     string `db:"location_street_name"`
	LocationCity           string `db:"location_city"`
	LocationState          string `db:"location_state"`
	LocationCountry        string `db:"location_country"`
	LocationPostcode       int    `db:"location_postcode"`
	LocationLatitude       string `db:"location_latitude"`
	LocationLongitude      string `db:"location_longitude"`
	LocationTimezoneOffset string `db:"location_timezone_offset"`
	LocationTimezoneDesc   string `db:"location_timezone_description"`
	LoginUUID              string `db:"login_uuid"`
	LoginUsername          string `db:"login_username"`
	LoginPassword          string `db:"login_password"`
	LoginSalt              string `db:"login_salt"`
	LoginMD5               string `db:"login_md5"`
	LoginSHA1              string `db:"login_sha1"`
	LoginSHA256            string `db:"login_sha256"`
	DobDate                string `db:"dob_date"`
	DobAge                 int    `db:"dob_age"`
	RegisteredDate         string `db:"registered_date"`
	RegisteredAge          int    `db:"registered_age"`
	IDName                 string `db:"id_name"`
	IDValue                string `db:"id_value"`
	PictureLarge           string `db:"picture_large"`
	PictureMedium          string `db:"picture_medium"`
	PictureThumbnail       string `db:"picture_thumbnail"`
}

func MapGetRandomUserResponseToUsers(response *randomuser.GetRandomUserResponse) []User {
	users := make([]User, 0, len(response.Results))
	for _, user := range response.Results {
		userSQL := User{
			Gender:                 user.Gender,
			Email:                  user.Email,
			Phone:                  user.Phone,
			Cell:                   user.Cell,
			Nat:                    user.Nat,
			NameTitle:              user.Name.Title,
			NameFirst:              user.Name.First,
			NameLast:               user.Name.Last,
			LocationStreetNumber:   user.Location.Street.Number,
			LocationStreetName:     user.Location.Street.Name,
			LocationCity:           user.Location.City,
			LocationState:          user.Location.State,
			LocationCountry:        user.Location.Country,
			LocationPostcode:       user.Location.Postcode,
			LocationLatitude:       user.Location.Coordinates.Latitude,
			LocationLongitude:      user.Location.Coordinates.Longitude,
			LocationTimezoneOffset: user.Location.Timezone.Offset,
			LocationTimezoneDesc:   user.Location.Timezone.Description,
			LoginUUID:              user.Login.UUID,
			LoginUsername:          user.Login.Username,
			LoginPassword:          user.Login.Password,
			LoginSalt:              user.Login.Salt,
			LoginMD5:               user.Login.MD5,
			LoginSHA1:              user.Login.SHA1,
			LoginSHA256:            user.Login.SHA256,
			DobDate:                user.Dob.Date,
			DobAge:                 user.Dob.Age,
			RegisteredDate:         user.Registered.Date,
			RegisteredAge:          user.Registered.Age,
			IDName:                 user.ID.Name,
			IDValue:                user.ID.Value,
			PictureLarge:           user.Picture.Large,
			PictureMedium:          user.Picture.Medium,
			PictureThumbnail:       user.Picture.Thumbnail,
		}
		users = append(users, userSQL)
	}

	return users
}
