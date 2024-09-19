// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package api

import (
	"context"
	"github.com/google/uuid"
	"github.com/synnaxlabs/synnax/pkg/distribution/ontology"
	access2 "github.com/synnaxlabs/synnax/pkg/service/access"
	rbac2 "github.com/synnaxlabs/synnax/pkg/service/access/rbac"
	"github.com/synnaxlabs/x/gorp"
	"go/types"
)

type AccessService struct {
	internal *rbac2.Service
	dbProvider
}

func NewAccessService(p Provider) *AccessService {
	return &AccessService{
		internal:   p.RBAC,
		dbProvider: p.db,
	}
}

type (
	AccessCreatePolicyRequest struct {
		Policies []rbac2.Policy `json:"policies" msgpack:"policies"`
	}
	AccessCreatePolicyResponse = AccessCreatePolicyRequest
)

func (a *AccessService) CreatePolicy(ctx context.Context, req AccessCreatePolicyRequest) (AccessCreatePolicyResponse, error) {
	if err := a.internal.Enforce(ctx, access2.Request{
		Subject: getSubject(ctx),
		Objects: []ontology.ID{{Type: rbac2.OntologyType}},
		Action:  access2.Create,
	}); err != nil {
		return AccessCreatePolicyRequest{}, err
	}
	results := make([]rbac2.Policy, len(req.Policies))
	if err := a.WithTx(ctx, func(tx gorp.Tx) error {
		w := a.internal.NewWriter(tx)
		for i, p := range req.Policies {
			if p.Key == uuid.Nil {
				p.Key = uuid.New()
			}
			if err := w.Create(ctx, &p); err != nil {
				return err
			}
			results[i] = p
		}
		return nil
	}); err != nil {
		return AccessCreatePolicyRequest{}, err
	}
	return AccessCreatePolicyResponse{Policies: results}, nil
}

type AccessDeletePolicyRequest struct {
	Keys []uuid.UUID `json:"keys" msgpack:"keys"`
}

func (a *AccessService) DeletePolicy(ctx context.Context, req AccessDeletePolicyRequest) (types.Nil, error) {
	if err := a.internal.Enforce(ctx, access2.Request{
		Subject: getSubject(ctx),
		Objects: rbac2.OntologyIDs(req.Keys),
		Action:  access2.Delete,
	}); err != nil {
		return types.Nil{}, err
	}
	return types.Nil{}, a.WithTx(ctx, func(tx gorp.Tx) error {
		w := a.internal.NewWriter(tx)
		for _, key := range req.Keys {
			if err := w.Delete(ctx, key); err != nil {
				return err
			}
		}
		return nil
	})
}

type (
	AccessRetrievePolicyRequest struct {
		Subject ontology.ID `json:"subject" msgpack:"subject"`
	}
	AccessRetrievePolicyResponse struct {
		Policies []rbac2.Policy `json:"policies" msgpack:"policies"`
	}
)

func (a *AccessService) RetrievePolicy(
	ctx context.Context,
	req AccessRetrievePolicyRequest,
) (res AccessRetrievePolicyResponse, err error) {
	res.Policies = make([]rbac2.Policy, 0)

	if err = a.internal.NewRetriever().
		WhereSubject(req.Subject).
		Entries(&res.Policies).
		Exec(ctx, nil); err != nil {
		return AccessRetrievePolicyResponse{}, err
	}
	if err = a.internal.Enforce(ctx, access2.Request{
		Subject: getSubject(ctx),
		Action:  access2.Retrieve,
		Objects: rbac2.OntologyIDsFromPolicies(res.Policies),
	}); err != nil {
		return AccessRetrievePolicyResponse{}, err
	}
	return res, nil
}
