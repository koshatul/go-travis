// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Fragments of this file have been copied from the go-github (https://github.com/google/go-github)
// project, and is therefore licensed under the following copyright:
// Copyright 2013 The go-github AUTHORS. All rights reserved.

package travis

import (
	"fmt"
	"net/http"
)

// BranchesService handles communication with the branches
// related methods of the Travis CI API.
type BranchesService struct {
	client *Client
}

// Branch represents a Travis CI build
type Branch struct {
	Id           uint   `json:"id,omitempty"`
	RepositoryId uint   `json:"repository_id,omitempty"`
	CommitId     uint   `json:"commit_id,omitempty"`
	Number       string `json:"number,omitempty"`
	Config       Config `json:"config,omitempty"`
	State        string `json:"state,omitempty"`
	StartedAt    string `json:"started_at,omitempty"`
	FinishedAt   string `json:"finished_at,omitempty"`
	Duration     uint   `json:"duration,omitempty"`
	JobIds       []uint `json:"job_ids,omitempty"`
	PullRequest  bool   `json:"pull_request,omitempty"`
}

// ListBranchesResponse represents the response of a call
// to the Travis CI list branches endpoint.
type ListBranchesResponse struct {
	Branches []Branch `json:"branches"`
}

// GetBranchResponse represents the response of a call
// to the Travis CI get branch endpoint.
type GetBranchResponse struct {
	Branch *Branch `json:"branch"`
}

// List the branches of a given repository.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BranchesService) List(repositorySlug string) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%v/branches", repositorySlug), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branchesResp ListBranchesResponse
	resp, err := bs.client.Do(req, &branchesResp)
	if err != nil {
		return nil, resp, err
	}

	return branchesResp.Branches, resp, err
}

// Get fetches a branch based on the provided repository slug
// and it's id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BranchesService) Get(repositorySlug string, branchId uint) (*Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%v/branches/%d", repositorySlug, branchId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branchResp GetBranchResponse
	resp, err := bs.client.Do(req, &branchResp)
	if err != nil {
		return nil, resp, err
	}

	return branchResp.Branch, resp, err
}
