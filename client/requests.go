package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	//registryUrl      = "https://raw.githubusercontent.com/cosmos/chain-registry/master"
	validatorBaseUrl = "https://api.mintscan.io/v1"
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

//func (r requests) GetChain(name string) (Chain, error) {
//	query := fmt.Sprintf("%s/%s/chain.json", registryUrl, name)
//	resp, err := r.client.Get(query)
//	if err != nil {
//		return Chain{}, err
//	}
//
//	// If the chain.json file doesn't exist we simply ignore it
//	if resp.StatusCode == http.StatusNotFound {
//		return Chain{}, nil
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return Chain{}, fmt.Errorf("unexpected status code from query %s: %d", query, resp.StatusCode)
//	}
//
//	var chain Chain
//	err = json.NewDecoder(resp.Body).Decode(&chain)
//	if err != nil {
//		return Chain{}, err
//	}
//
//	return chain, nil
//}

func (r requests) GetValidatorList(ctx context.Context, chain string) ([]Validator, error) {
	query := fmt.Sprintf("https://lcd-%s.keplr.app/cosmos/staking/v1beta1/validators", chain)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "PostmanRuntime/7.29.2")

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
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "PostmanRuntime/7.29.2")

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
