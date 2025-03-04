package config

// DefaultValues is the default configuration
const DefaultValues = `
IsTrustedSequencer = false

[Log]
Level = "debug"
Outputs = ["stdout"]

[StateDB]
User = "state_user"
Password = "state_password"
Name = "state_db"
Host = "localhost"
Port = "5432"
EnableLog = false
MaxConns = 200

[PoolDB]
User = "pool_user"
Password = "pool_password"
Name = "pool_db"
Host = "localhost"
Port = "5432"
EnableLog = false
MaxConns = 200

[Etherman]
URL = "http://localhost:8545"
PrivateKeyPath = "./test/test.keystore"
PrivateKeyPassword = "testonly"

[EthTxManager]
MaxSendBatchTxRetries = 10
MaxVerifyBatchTxRetries = 10
FrequencyForResendingFailedSendBatches = "1s"
FrequencyForResendingFailedVerifyBatch = "1s"
WaitTxToBeMined = "2m"
PercentageToIncreaseGasPrice = 10
PercentageToIncreaseGasLimit = 10

[RPC]
Host = "0.0.0.0"
Port = 8123
MaxRequestsPerIPAndSecond = 50
SequencerNodeURI = ""
BroadcastURI = "127.0.0.1:61090"
DefaultSenderAddress = "0x1111111111111111111111111111111111111111"
	[RPC.DB]
		User = "rpc_user"
		Password = "rpc_password"
		Name = "rpc_db"
		Host = "localhost"
		Port = "5432"
		EnableLog = false
		MaxConns = 200

[Synchronizer]
SyncInterval = "0s"
SyncChunkSize = 100
TrustedSequencerURI = ""

[Sequencer]
MaxSequenceSize = "2000000"
WaitPeriodPoolIsEmpty = "1s"
WaitPeriodSendSequence = "15s"
LastBatchVirtualizationTimeMaxWaitPeriod = "300s"
WaitBlocksToUpdateGER = 10
MaxTimeForBatchToBeOpen = "15s"
BlocksAmountForTxsToBeDeleted = 100
FrequencyToCheckTxsForDelete = "12h"
MaxCumulativeGasUsed = 30000000
MaxKeccakHashes = 468
MaxPoseidonHashes = 279620
MaxPoseidonPaddings = 149796
MaxMemAligns = 262144
MaxArithmetics = 262144
MaxBinaries = 262144
MaxSteps = 8388608
	[Sequencer.ProfitabilityChecker]
		SendBatchesEvenWhenNotProfitable = "true"

[PriceGetter]
Type = "default"
DefaultPrice = "2000"

[Aggregator]
IntervalFrequencyToGetProofGenerationState = "5s"
IntervalToConsolidateState = "3s"
TxProfitabilityCheckerType = "acceptall"
TxProfitabilityMinReward = "1.1"

[GasPriceEstimator]
Type = "default"
DefaultGasPriceWei = 1000000000

[Prover]
ProverURI = "0.0.0.0:50051"

[MTServer]
Host = "0.0.0.0"
Port = 50060
StoreBackend = "PostgreSQL"

[MTClient]
URI = "127.0.0.1:50061"

[Executor]
URI = "127.0.0.1:50071"

[BroadcastServer]
Host = "0.0.0.0"
Port = 61090
`
