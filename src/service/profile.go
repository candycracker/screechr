package service

import (
	"errors"
	"time"
)

/*
● Id – positive integer, upper limit >10 billion, uniquely identifies the user in the system,
cannot be null.
● User name – The public user name of the user. 80 characters, unique, cannot be null or
all whitespace.
● First Name – The first name of the user. 100 characters, cannot be null or whitespace.
● Last Name – The last name of the user. 100 characters, cannot be null or whitespace.
● Secret token – String used to validate requests as coming from a specific user for this
demo. 32 characters (single byte OK), unique and cannot be null.
● Profile Image - Just a URI to a profile image. Can be null.
● Date created – The date & time the user was added to the system. Cannot be null.
● Date modified – The most recent date & time the user was added to the system.
*/
type ProfileType struct {
	ID           int64
	UserName     string
	FirstName    string
	LastName     string
	Token        string
	ImageUrl     string
	CreatedTime  string //time.RFC3339
	ModifiedTime string //time.RFC3339
}

type profile struct {
	db map[int64]*ProfileType
}

type ProfileService interface {
	GetProfile(id int64) (ProfileType, error)
	UpdatePicture(id int64, url string) error
	UpdateProfile(id int64, username, firstname, lastname, token, url string) error
}

var profileDb = map[int64]*ProfileType{
	1234567890: {
		ID:           1234567890,
		UserName:     "JohnXXX",
		FirstName:    "John",
		LastName:     "James",
		Token:        "JG3LDSkEzbQgnGcIU7o1P8p2FxuHUMg8", //generate by RandStringBytes
		ImageUrl:     "sdasdasd.jpg",
		CreatedTime:  "2022-08-11T17:41:59-07:00",
		ModifiedTime: "2022-08-11T17:41:59-07:00",
	},
	9987654321: {
		ID:           9987654321,
		UserName:     "AliceXXX",
		FirstName:    "Alice",
		LastName:     "May",
		Token:        "Tex1h8VF75cB3Y6WTRPVF6hkUUFrK9lj",
		ImageUrl:     "asqewfasd.jpg",
		CreatedTime:  "2022-08-11T17:41:59-07:00",
		ModifiedTime: "2022-08-11T17:41:59-07:00",
	},
}

func Profile() ProfileService {
	return &profile{db: profileDb}
}

func (p *profile) GetProfile(id int64) (ProfileType, error) {
	profile, ok := p.db[id]
	if !ok {
		return ProfileType{}, errors.New("unable to find profile")
	}
	return *profile, nil
}

func (p *profile) UpdatePicture(id int64, url string) error {
	profile, ok := p.db[id]
	if !ok {
		return errors.New("unable to find profile")
	}
	profile.ImageUrl = url
	profile.ModifiedTime = time.Now().Format(time.RFC3339)
	return nil
}

func (p *profile) UpdateProfile(id int64, username, firstname, lastname, token, url string) error {
	profile, ok := p.db[id]
	if !ok {
		return errors.New("unable to find profile")
	}
	profile.UserName = username
	profile.FirstName = firstname
	profile.LastName = lastname
	profile.Token = token
	profile.ImageUrl = url
	profile.ModifiedTime = time.Now().Format(time.RFC3339)
	return nil
}
