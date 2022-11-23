package cmd

import (
	"context"
	req "cosmos-tools/client"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"os"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "generate",
		Short: "generates CSVs with specified cosmos chain data",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generate(chainName); err != nil {
				return err
			}
			return nil
		},
	})
}

type ValidatorList struct {
	ValidatorMoniker string `csv:"validator_moniker"`
	PerVotingPower   string `csv:"per_voting_power"`
	SelfDelegation   string `csv:"self_delegation"`
	TotalDelegation  int    `csv:"total_delegation"`
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

func generate(chain string) error {
	client := req.NewRequests()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	validators, err := client.GetValidatorList(ctx, chain)
	if err != nil {
		return err
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
		delegators, err := client.GetDelegators(ctx, validator.OperatorAddress, chain)
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
	}

	err = WriteCSV("validators.csv", validatorsCSV)
	if err != nil {
		return err
	}

	err = WriteCSV("delegators.csv", delegatorCSV)
	if err != nil {
		return err
	}

	err = WriteCSV("delegatorValidators.csv", delegatorValidatorsCSV)
	if err != nil {
		return err
	}

	return nil
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
