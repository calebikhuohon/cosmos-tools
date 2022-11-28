package cmd

import (
	"cosmos-tools/pkg/file"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "analyze",
		Short: "produces data analysis from Umee genesis files",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := analyze(); err != nil {
				return err
			}
			fmt.Println("Vesting account analysis complete")
			return nil
		},
	})
}

func analyze() error {
	fs := file.New("genesis.json")
	b, err := fs.ReadFile()
	if err != nil {
		return err
	}

	var gen Genesis
	if err := json.Unmarshal(b, &gen); err != nil {
		return err
	}

	var result struct {
		Sch []UnlockSchedule `json:"unlock_schedule"`
	}

	for _, n := range gen.AppState.Auth.Accounts {
		st, err := strconv.ParseInt(n.StartTime, 10, 64)
		if n.StartTime != "" && err != nil {
			return err
		}
		startTime := time.Unix(st, 0).UTC()

		et, err := strconv.ParseInt(n.BaseVestingAccount.EndTime, 0, 64)
		if n.BaseVestingAccount.EndTime != "" && err != nil {
			return err
		}
		endTime := time.Unix(et, 0).UTC()

		result.Sch = append(result.Sch, UnlockSchedule{
			StartTime: startTime,
			EndTime:   endTime,
		})
	}

	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err
	}

	f, err := os.OpenFile("unlock-schedule.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
