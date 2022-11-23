package main

import (
	"context"
	req "cosmos-tools/client"
	"fmt"
	"github.com/gocarina/gocsv"
	"golang.org/x/exp/slices"
	"log"
	"os"
)

type ValidatorList struct {
	ValidatorMoniker string `json:"validator_moniker" ,csv:"validator_moniker"`
	PerVotingPower   string `json:"per_voting_power" ,csv:"per_voting_power"`
	SelfDelegation   string `json:"self_delegation" ,csv:"self_delegation"`
	TotalDelegation  int    `json:"total_delegation" ,csv:"total_delegation"`
}

type DelegatorList struct {
	Delegator   string `csv:"delegator"`
	VotingPower string `csv:"voting_power"`
}

type DelegatorValidatorsList struct {
	DelegatorAcc string `csv:"delegator_acc"`
	ValidatorAcc string `csv:"validator_acc"`
	BondedToken  string `csv:"bonded_token"`
}

func main() {
	//take in a cosmos chain name
	client := req.NewRequests()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	validators, err := client.GetValidatorList(ctx, os.Args[1])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	slices.SortFunc(validators, func(a, b req.Validator) bool {
		return a.Tokens > b.Tokens
	})

	var (
		delegatorCSV           []DelegatorList
		validatorsCSV          []ValidatorList
		delegatorValidatorsCSV []DelegatorValidatorsList
	)

	for _, validator := range validators {
		delegators, err := client.GetDelegators(ctx, validator.OperatorAddress, os.Args[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		slices.SortFunc(delegators, func(a, b req.DelegationResponse) bool {
			return a.Delegation.Shares > b.Delegation.Shares
		})

		validatorsCSV = append(validatorsCSV, ValidatorList{
			ValidatorMoniker: validator.Description.Moniker,
			PerVotingPower:   validator.Tokens,
			SelfDelegation:   validator.MinSelfDelegation,
			TotalDelegation:  len(delegators),
		})

		for i := 0; i < len(delegators); i++ {
			for j := 0; j < len(delegators[i:]); j++ {
				id := delegators[i].Delegation
				jd := delegators[j].Delegation

				if id.DelegatorAddress == jd.DelegatorAddress {
					delegatorValidatorsCSV = append(delegatorValidatorsCSV, DelegatorValidatorsList{
						DelegatorAcc: id.DelegatorAddress,
						ValidatorAcc: id.ValidatorAddress,
						BondedToken:  delegators[i].Balance.Denom,
					})
				}
			}

			delegatorCSV = append(delegatorCSV, DelegatorList{
				Delegator:   delegators[i].Delegation.DelegatorAddress,
				VotingPower: delegators[i].Delegation.Shares,
			})
		}

		fmt.Println(validator.OperatorAddress)
	}

	err = WriteCSV("validators.csv", validatorsCSV)
	if err != nil {
		log.Fatalln(err)
	}

	err = WriteCSV("delegators.csv", delegatorCSV)
	if err != nil {
		log.Fatalln(err)
	}

	err = WriteCSV("delegatorValidators.csv", delegatorValidatorsCSV)
	if err != nil {
		log.Fatalln(err)
	}
}

func WriteCSV(filename string, data any) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	err = gocsv.Marshal(data, f)
	if err != nil {
		return fmt.Errorf("failed to marshal csv: %v", err)
	}
	return nil
}
