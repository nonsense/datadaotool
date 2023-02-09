# datadaotool

This repo/tool is supposed to interact with the proof-of-concept [DealClient.sol](https://github.com/nonsense/fevm-hardhat-kit/blob/nonsense/deal-client-contract/contracts/basic-deal-client/DealClient.sol) deal proposal contract at

# commands

- `submit-deal` -- submits a deal proposal on chain for SP's that have `AcceptContractDeals: true` config enables and are monitoring the Filecoin chain for deals
- `events-monitor` -- simple monitor tool to listen for `DealProposalCreate(bytes32)` events emitted on chain
