package client

import "time"

type Validator struct {
	OperatorAddress   string          `json:"operator_address"`
	ConsensusPubkey   ConsensusPubkey `json:"consensus_pubkey"`
	Jailed            bool            `json:"jailed"`
	Status            string          `json:"status"`
	Tokens            string          `json:"tokens"`
	DelegatorShares   string          `json:"delegator_shares"`
	Description       Description     `json:"description"`
	UnbondingHeight   string          `json:"unbonding_height"`
	UnbondingTime     string          `json:"unbonding_time"`
	Commission        Commission      `json:"commission"`
	MinSelfDelegation string          `json:"min_self_delegation"`
}

type CommissionRates struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}

type Commission struct {
	CommissionRates CommissionRates `json:"commission_rates"`
	UpdateTime      time.Time       `json:"update_time"`
}

type ConsensusPubkey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}

type Uptime struct {
	Address      string `json:"address"`
	MissedBlocks int    `json:"missed_blocks"`
	OverBlocks   int    `json:"over_blocks"`
}

type Delegator struct {
	DelegatorAddress string `json:"delegator_address"`
	Amount           string `json:"amount"`
}

type Delegation struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Shares           string `json:"shares"`
}

type DelegationBalance struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type DelegationResponse struct {
	Delegation Delegation        `json:"delegation"`
	Balance    DelegationBalance `json:"balance"`
}
