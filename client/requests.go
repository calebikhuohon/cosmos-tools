package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Requests interface {
	GetValidatorList(ctx context.Context, chain string) ([]Validator, error)
	GetDelegators(ctx context.Context, validatorAddress, chain string) ([]DelegationResponse, error)
	GetDelegatorValidators(delegatorAddress, chain string) ([]DelegationResponse, error)
}

type requests struct {
	client *http.Client
}

func NewRequests() Requests {
	return &requests{client: http.DefaultClient}
}

func (r requests) GetValidatorList(ctx context.Context, chain string) ([]Validator, error) {
	query := fmt.Sprintf("https://lcd-%s.keplr.app/cosmos/staking/v1beta1/validators", chain)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code from query %s: %d - %v", query, resp.StatusCode, string(b))
	}

	var result struct {
		Validators []Validator `json:"validators"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse json: %v", err)
	}

	return result.Validators, nil
}

func (r requests) GetDelegators(ctx context.Context, validatorAddress, chain string) ([]DelegationResponse, error) {
	query := fmt.Sprintf("https://%s-api.lavenderfive.com/cosmos/staking/v1beta1/validators/%s/delegations", chain, validatorAddress)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	// If the chain.json file doesn't exist we simply ignore it
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code from query %s: %d - %v", query, resp.StatusCode, string(b))
	}

	var result struct {
		DelegationResponses []DelegationResponse `json:"delegation_responses"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.DelegationResponses, nil
}

func (r requests) GetDelegatorValidators(delegatorAddress, chain string) ([]DelegationResponse, error) {
	query := fmt.Sprintf("https://lcd-%s.keplr.app/cosmos/staking/v1beta1/delegations/%s?pagination.limit=1000", chain, delegatorAddress)
	resp, err := r.client.Get(query)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from query %s: %d", query, resp.StatusCode)
	}

	var result struct {
		DelegationResponses []DelegationResponse `json:"delegation_responses"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.DelegationResponses, nil
}
