package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AutomatedElectricalSolutions/accounts-api-client-go/models"
)

// GetUser gets a user from account api given a user id
func GetUser(userID string) (*models.User, error) {
	if client == nil {
		return nil, fmt.Errorf("Client has not been initialised")
	}

	res, err := client.Get(fmt.Sprintf("https://accounts.automated.net.au/v1/users/%s", userID))

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
