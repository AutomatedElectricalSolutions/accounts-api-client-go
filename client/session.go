package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AutomatedElectricalSolutions/accounts-api-client-go/models"
)

// IntrospectSession fetches the session information from the account api
func IntrospectSession(sessionID string) (*models.Session, error) {
	// we rely on the client, so we need to make sure it's initialised
	// and authorised
	if client == nil {
		return nil, fmt.Errorf("Client has not been initialised")
	}

	// make the request to the accountapi
	res, err := client.Get(fmt.Sprintf("https://accounts.automated.net.au/v1/web/session/%s", sessionID))

	if err != nil {
		return nil, err
	}

	// read the body of the request
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// unmarshal the json into the struct
	session := models.Session{}

	session.IsValid = res.StatusCode == 200

	switch res.StatusCode {
	case 200:
		// populate all the values with the body
		err = json.Unmarshal(body, &session)
		if err != nil {
			return nil, err
		}
		return &session, nil
	case 404:
		return &session, nil
	default:
		return nil, fmt.Errorf("Issue requesting information about the session %s", body)
	}
}
