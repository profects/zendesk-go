package zendesk

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type UserApiHandler struct {
	client Client
}

type SingleUser struct {
	Response User `json:"user"`
}

type MultipleUser struct {
	Response []User `json:"users"`
}

func (u UserApiHandler) GetById(id int) (User, error) {
	response, err := u.client.get(
		fmt.Sprintf("/users/%d.json", id),
		nil,
	)

	if err != nil {

	}

	return u.parseSingleObject(response), err
}

func (u UserApiHandler) GetAll() ([]User, error) {
	response, err := u.client.get(
		"/users.json",
		nil,
	)

	if err != nil {

	}

	return u.parseMultiObjects(response), err
}

func (u UserApiHandler) Create(v User) (User, error) {
	var object SingleUser

	_, err := u.client.post(
		"/users.json",
		map[string]User{"user": v},
		&object,
	)

	return object.Response, err
}

func (u UserApiHandler) CreateOrUpdate(v User) (User, error) {
	var object SingleUser

	_, err := u.client.post(
		"/users/create_or_update.json",
		map[string]User{"user": v},
		&object,
	)

	return object.Response, err
}

func (u UserApiHandler) CreateOrUpdateMany(v []User) (Job, error) {
	var object Job

	_, err := u.client.post(
		"/users/create_or_update_many.json",
		map[string][]User{"users": v},
		&object,
	)

	return object, err
}

func (u UserApiHandler) Update(v User) (User, error) {
	var object SingleUser

	_, err := u.client.put(
		fmt.Sprintf("/users/%d.json", v.Id),
		map[string]User{"user": v},
		&object,
	)

	return object.Response, err
}

func (u UserApiHandler) Delete(id int) (int, error) {
	response, err := u.client.delete(
		fmt.Sprintf("/users/%d.json", id),
	)

	return response.StatusCode(), err
}

func (u UserApiHandler) parseMultiObjects(response *resty.Response) []User {
	var object MultipleUser

	u.client.parseResponseToInterface(response, &object)

	return object.Response
}

func (u UserApiHandler) parseSingleObject(response *resty.Response) User {
	var object SingleUser

	u.client.parseResponseToInterface(response, &object)

	return object.Response
}