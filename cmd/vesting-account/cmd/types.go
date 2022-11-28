package cmd

import "time"

type ConsensusParams struct {
	Block     ConsensusParamsBlock     `json:"block"`
	Evidence  ConsensusParamsEvidence  `json:"evidence"`
	Validator ConsensusParamsValidator `json:"validator"`
	Version   ConsensusParamsVersion   `json:"version"`
}

type ConsensusParamsBlock struct {
	MaxBytes   string `json:"max_bytes"`
	MaxGas     string `json:"max_gas"`
	TimeIotaMs string `json:"time_iota_ms"`
}

type ConsensusParamsEvidence struct {
	MaxAgeNumBlocks string `json:"max_age_num_blocks"`
	MaxAgeDuration  string `json:"max_age_duration"`
	MaxBytes        string `json:"max_bytes"`
}

type ConsensusParamsValidator struct {
	PubKeyTypes []string `json:"pub_key_types"`
}

type ConsensusParamsVersion struct{}

type AppStateAuthParams struct {
	MaxMemoCharacters      string `json:"max_memo_characters"`
	TxSigLimit             string `json:"tx_sig_limit"`
	TxSizeCostPerByte      string `json:"tx_size_cost_per_byte"`
	SigVerifyCostEd25519   string `json:"sig_verify_cost_ed25519"`
	SigVerifyCostSecp256K1 string `json:"sig_verify_cost_secp256k1"`
}

type BaseAccount struct {
	Address       string      `json:"address"`
	PubKey        interface{} `json:"pub_key"`
	AccountNumber string      `json:"account_number"`
	Sequence      string      `json:"sequence"`
}

type OriginalVesting struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type BaseVestingAccount struct {
	BaseAccount      BaseAccount       `json:"base_account"`
	OriginalVesting  []OriginalVesting `json:"original_vesting"`
	DelegatedFree    []interface{}     `json:"delegated_free"`
	DelegatedVesting []interface{}     `json:"delegated_vesting"`
	EndTime          string            `json:"end_time"`
}

type AppStateAuthAccount struct {
	Type               string             `json:"@type"`
	BaseVestingAccount BaseVestingAccount `json:"base_vesting_account,omitempty"`
	Address            string             `json:"address,omitempty"`
	PubKey             interface{}        `json:"pub_key,omitempty"`
	AccountNumber      string             `json:"account_number,omitempty"`
	Sequence           string             `json:"sequence,omitempty"`
	StartTime          string             `json:"start_time,omitempty"`
}

type Auth struct {
	Params   AppStateAuthParams    `json:"params"`
	Accounts []AppStateAuthAccount `json:"accounts"`
}

type Authz struct {
	Authorization []interface{} `json:"authorization"`
}

type BankParams struct {
	SendEnabled        []interface{} `json:"send_enabled"`
	DefaultSendEnabled bool          `json:"default_send_enabled"`
}

type BankBalance struct {
	Address string `json:"address"`
	Coins   []Coin `json:"coins"`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type DenomMetadata struct {
	Description string      `json:"description"`
	DenomUnits  []DenomUnit `json:"denom_units"`
	Base        string      `json:"base"`
	Display     string      `json:"display"`
	Name        string      `json:"name"`
	Symbol      string      `json:"symbol"`
}

type DenomUnit struct {
	Denom    string   `json:"denom"`
	Exponent int      `json:"exponent"`
	Aliases  []string `json:"aliases"`
}

type Bank struct {
	Params        BankParams      `json:"params"`
	Balances      []BankBalance   `json:"balances"`
	Supply        []Coin          `json:"supply"`
	DenomMetadata []DenomMetadata `json:"denom_metadata"`
}

type Capability struct {
	Index  string        `json:"index"`
	Owners []interface{} `json:"owners"`
}

type Crisis struct {
	ConstantFee ConstantFee `json:"constant_fee"`
}

type ConstantFee struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Distribution struct {
	Params                          DistributionParams `json:"params"`
	FeePool                         FeePool            `json:"fee_pool"`
	DelegatorWithdrawInfos          []interface{}      `json:"delegator_withdraw_infos"`
	PreviousProposer                string             `json:"previous_proposer"`
	OutstandingRewards              []interface{}      `json:"outstanding_rewards"`
	ValidatorAccumulatedCommissions []interface{}      `json:"validator_accumulated_commissions"`
	ValidatorHistoricalRewards      []interface{}      `json:"validator_historical_rewards"`
	ValidatorCurrentRewards         []interface{}      `json:"validator_current_rewards"`
	DelegatorStartingInfos          []interface{}      `json:"delegator_starting_infos"`
	ValidatorSlashEvents            []interface{}      `json:"validator_slash_events"`
}

type FeePool struct {
	CommunityPool []interface{} `json:"community_pool"`
}

type DistributionParams struct {
	CommunityTax        string `json:"community_tax"`
	BaseProposerReward  string `json:"base_proposer_reward"`
	BonusProposerReward string `json:"bonus_proposer_reward"`
	WithdrawAddrEnabled bool   `json:"withdraw_addr_enabled"`
}

type Evidence struct {
	Evidence []interface{} `json:"evidence"`
}

type FeeGrant struct {
	Allowances []interface{} `json:"allowances"`
}

type Body struct {
	Messages                    []Message     `json:"messages"`
	Memo                        string        `json:"memo"`
	TimeoutHeight               string        `json:"timeout_height"`
	ExtensionOptions            []interface{} `json:"extension_options"`
	NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
}

type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}

type Commission struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}

type Pubkey struct {
	Type string `json:"@type"`
	Key  string `json:"key"`
}

type Message struct {
	Type              string      `json:"@type"`
	Description       Description `json:"description,omitempty"`
	Commission        Commission  `json:"commission,omitempty"`
	MinSelfDelegation string      `json:"min_self_delegation,omitempty"`
	DelegatorAddress  string      `json:"delegator_address,omitempty"`
	ValidatorAddress  string      `json:"validator_address,omitempty"`
	Pubkey            Pubkey      `json:"pubkey,omitempty"`
	Value             Coin        `json:"value,omitempty"`
	Validator         string      `json:"validator,omitempty"`
	Orchestrator      string      `json:"orchestrator,omitempty"`
	EthAddress        string      `json:"eth_address,omitempty"`
}

type SignerInfo struct {
	PublicKey Pubkey   `json:"public_key"`
	ModeInfo  ModeInfo `json:"mode_info"`
	Sequence  string   `json:"sequence"`
}

type ModeInfo struct {
	Single ModeInfoSingle `json:"single"`
}

type ModeInfoSingle struct {
	Mode string `json:"mode"`
}

type Fee struct {
	Amount   []interface{} `json:"amount"`
	GasLimit string        `json:"gas_limit"`
	Payer    string        `json:"payer"`
	Granter  string        `json:"granter"`
}

type AuthInfo struct {
	SignerInfos []SignerInfo `json:"signer_infos"`
	Fee         Fee          `json:"fee"`
}

type GenTx struct {
	Body       Body     `json:"body"`
	AuthInfo   AuthInfo `json:"auth_info"`
	Signatures []string `json:"signatures"`
}

type GenUtil struct {
	GenTxs []GenTx `json:"gen_txs"`
}

type Gov struct {
	StartingProposalID string          `json:"starting_proposal_id"`
	Deposits           []interface{}   `json:"deposits"`
	Votes              []interface{}   `json:"votes"`
	Proposals          []interface{}   `json:"proposals"`
	DepositParams      DepositParams   `json:"deposit_params"`
	VotingParams       GovVotingParams `json:"voting_params"`
	TallyParams        GovTallyParams  `json:"tally_params"`
}

type GovVotingParams struct {
	VotingPeriod string `json:"voting_period"`
}

type GovTallyParams struct {
	Quorum        string `json:"quorum"`
	Threshold     string `json:"threshold"`
	VetoThreshold string `json:"veto_threshold"`
}

type DepositParams struct {
	MinDeposit       []Coin `json:"min_deposit"`
	MaxDepositPeriod string `json:"max_deposit_period"`
}

type GravityParams struct {
	GravityID                    string   `json:"gravity_id"`
	ContractSourceHash           string   `json:"contract_source_hash"`
	BridgeEthereumAddress        string   `json:"bridge_ethereum_address"`
	BridgeChainID                string   `json:"bridge_chain_id"`
	SignedValsetsWindow          string   `json:"signed_valsets_window"`
	SignedBatchesWindow          string   `json:"signed_batches_window"`
	SignedLogicCallsWindow       string   `json:"signed_logic_calls_window"`
	TargetBatchTimeout           string   `json:"target_batch_timeout"`
	AverageBlockTime             string   `json:"average_block_time"`
	AverageEthereumBlockTime     string   `json:"average_ethereum_block_time"`
	SlashFractionValset          string   `json:"slash_fraction_valset"`
	SlashFractionBatch           string   `json:"slash_fraction_batch"`
	SlashFractionLogicCall       string   `json:"slash_fraction_logic_call"`
	UnbondSlashingValsetsWindow  string   `json:"unbond_slashing_valsets_window"`
	SlashFractionBadEthSignature string   `json:"slash_fraction_bad_eth_signature"`
	ValsetReward                 Coin     `json:"valset_reward"`
	BridgeActive                 bool     `json:"bridge_active"`
	EthereumBlacklist            []string `json:"ethereum_blacklist"`
}

type Gravity struct {
	Params             GravityParams `json:"params"`
	GravityNonces      GravityNonces `json:"gravity_nonces"`
	Valsets            []interface{} `json:"valsets"`
	ValsetConfirms     []interface{} `json:"valset_confirms"`
	Batches            []interface{} `json:"batches"`
	BatchConfirms      []interface{} `json:"batch_confirms"`
	LogicCalls         []interface{} `json:"logic_calls"`
	LogicCallConfirms  []interface{} `json:"logic_call_confirms"`
	Attestations       []interface{} `json:"attestations"`
	DelegateKeys       []interface{} `json:"delegate_keys"`
	Erc20ToDenoms      []interface{} `json:"erc20_to_denoms"`
	UnbatchedTransfers []interface{} `json:"unbatched_transfers"`
}

type GravityNonces struct {
	LatestValsetNonce         string `json:"latest_valset_nonce"`
	LastObservedNonce         string `json:"last_observed_nonce"`
	LastSlashedValsetNonce    string `json:"last_slashed_valset_nonce"`
	LastSlashedBatchBlock     string `json:"last_slashed_batch_block"`
	LastSlashedLogicCallBlock string `json:"last_slashed_logic_call_block"`
	LastTxPoolID              string `json:"last_tx_pool_id"`
	LastBatchID               string `json:"last_batch_id"`
}

type IBC struct {
	ClientGenesis     `json:"client_genesis"`
	ConnectionGenesis ConnectionGenesis `json:"connection_genesis"`
	ChannelGenesis    ChannelGenesis    `json:"channel_genesis"`
}

type ClientGenesis struct {
	Clients            []interface{}       `json:"clients"`
	ClientsConsensus   []interface{}       `json:"clients_consensus"`
	ClientsMetadata    []interface{}       `json:"clients_metadata"`
	Params             ClientGenesisParams `json:"params"`
	CreateLocalhost    bool                `json:"create_localhost"`
	NextClientSequence string              `json:"next_client_sequence"`
}

type ClientGenesisParams struct {
	AllowedClients []string `json:"allowed_clients"`
}

type ConnectionGenesis struct {
	Connections            []interface{}           `json:"connections"`
	ClientConnectionPaths  []interface{}           `json:"client_connection_paths"`
	NextConnectionSequence string                  `json:"next_connection_sequence"`
	Params                 ConnectionGenesisParams `json:"params"`
}

type ConnectionGenesisParams struct {
	MaxExpectedTimePerBlock string `json:"max_expected_time_per_block"`
}

type ChannelGenesis struct {
	Channels            []interface{} `json:"channels"`
	Acknowledgements    []interface{} `json:"acknowledgements"`
	Commitments         []interface{} `json:"commitments"`
	Receipts            []interface{} `json:"receipts"`
	SendSequences       []interface{} `json:"send_sequences"`
	RecvSequences       []interface{} `json:"recv_sequences"`
	AckSequences        []interface{} `json:"ack_sequences"`
	NextChannelSequence string        `json:"next_channel_sequence"`
}

type Mint struct {
	Minter Minter     `json:"minter"`
	Params MintParams `json:"params"`
}

type Minter struct {
	Inflation        string `json:"inflation"`
	AnnualProvisions string `json:"annual_provisions"`
}

type MintParams struct {
	MintDenom           string `json:"mint_denom"`
	InflationRateChange string `json:"inflation_rate_change"`
	InflationMax        string `json:"inflation_max"`
	InflationMin        string `json:"inflation_min"`
	GoalBonded          string `json:"goal_bonded"`
	BlocksPerYear       string `json:"blocks_per_year"`
}

type Slashing struct {
	Params       SlashingParams `json:"params"`
	SigningInfos []interface{}  `json:"signing_infos"`
	MissedBlocks []interface{}  `json:"missed_blocks"`
}
type SlashingParams struct {
	SignedBlocksWindow      string `json:"signed_blocks_window"`
	MinSignedPerWindow      string `json:"min_signed_per_window"`
	DowntimeJailDuration    string `json:"downtime_jail_duration"`
	SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
	SlashFractionDowntime   string `json:"slash_fraction_downtime"`
}

type Staking struct {
	Params               StakingParams `json:"params"`
	LastTotalPower       string        `json:"last_total_power"`
	LastValidatorPowers  []interface{} `json:"last_validator_powers"`
	Validators           []interface{} `json:"validators"`
	Delegations          []interface{} `json:"delegations"`
	UnbondingDelegations []interface{} `json:"unbonding_delegations"`
	Redelegations        []interface{} `json:"redelegations"`
	Exported             bool          `json:"exported"`
}

type StakingParams struct {
	UnbondingTime     string `json:"unbonding_time"`
	MaxValidators     int    `json:"max_validators"`
	MaxEntries        int    `json:"max_entries"`
	HistoricalEntries int    `json:"historical_entries"`
	BondDenom         string `json:"bond_denom"`
}

type Transfer struct {
	PortID      string         `json:"port_id"`
	DenomTraces []interface{}  `json:"denom_traces"`
	Params      TransferParams `json:"params"`
}

type TransferParams struct {
	SendEnabled    bool `json:"send_enabled"`
	ReceiveEnabled bool `json:"receive_enabled"`
}

type AppState struct {
	Auth         Auth         `json:"auth"`
	Authz        Authz        `json:"authz"`
	Bank         Bank         `json:"bank"`
	Capability   Capability   `json:"capability"`
	Crisis       Crisis       `json:"crisis"`
	Distribution Distribution `json:"distribution"`
	Evidence     Evidence     `json:"evidence"`
	Feegrant     FeeGrant     `json:"feegrant"`
	Genutil      GenUtil      `json:"genutil"`
	Gov          Gov          `json:"gov"`
	Gravity      Gravity      `json:"gravity"`
	Ibc          IBC          `json:"ibc"`
	Mint         Mint         `json:"mint"`
	Params       interface{}  `json:"params"`
	Slashing     `json:"slashing"`
	Staking      `json:"staking"`
	Transfer     `json:"transfer"`
	Upgrade      `json:"upgrade"`
	Vesting      `json:"vesting"`
}

type Upgrade struct {
}

type Vesting struct {
}

type Genesis struct {
	GenesisTime     time.Time       `json:"genesis_time"`
	ChainID         string          `json:"chain_id"`
	InitialHeight   string          `json:"initial_height"`
	ConsensusParams ConsensusParams `json:"consensus_params"`
	AppHash         string          `json:"app_hash"`
	AppState        AppState        `json:"app_state"`
}

type UnlockSchedule struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	//TODO: Token denom
	//TODO: chain level total supply
	//TODO: pubkey to unlock the token
}
