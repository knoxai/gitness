// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package user

import (
	"context"

	"github.com/harness/gitness/types"
)

// RegisterCheck checks the DB and env config flag to return boolean
// which represents if a user sign-up is allowed or not

func (c *Controller) RegisterCheck(ctx context.Context,
	config *types.Config) (*bool, error) {

	usrCount, err := c.principalStore.CountUsers(ctx)
	if err != nil {
		return nil, err
	}

	check := false
	if usrCount == 0 || config.AllowSignUp {
		check = true
	}

	return &check, nil
}
