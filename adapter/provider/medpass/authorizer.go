package medpass

import (
	"encoding/json"
	"errors"
	"path"
)

func (a *authorizer) ValidateToken(token string) error {
	rel, err := a.BaseURL.Parse(path.Join("validate-token"))
	if err != nil {
		return err
	}

	client := a.Client.R()

	client.Header.Set("AuthToken", token)

	resp, err := client.Get(rel.String())
	if err != nil {
		return err
	}
	if resp.IsError() {
		var authorizerGatewayError *AuthorizerGatewayError

		if err = json.Unmarshal(resp.Body(), &authorizerGatewayError); err != nil {
			return err
		}

		return errors.New(authorizerGatewayError.Error)
	}

	return nil
}
