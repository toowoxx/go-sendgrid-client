package sendgrid

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

func (c *Client) Request(endpoint string, method rest.Method, obj interface{}, result interface{}) error {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return errors.Wrap(err, "could not marshal object")
	}

	request := sendgrid.GetRequest(c.ApiKey, endpoint, sendGridHost)
	request.Method = method
	request.Body = requestBody
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	} else {
		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("server returned code %d, response %s", response.StatusCode, response.Body)
		}
	}

	if result != nil {
		if err := json.Unmarshal([]byte(response.Body), result); err != nil {
			return errors.Wrap(err, "could not unmarshal response body into result")
		}
	}

	return nil
}
