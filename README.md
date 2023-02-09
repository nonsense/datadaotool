# datadaotool

This repo/tool is supposed to interact with the proof-of-concept [DealClient.sol](https://github.com/nonsense/fevm-hardhat-kit/blob/nonsense/deal-client-contract/contracts/basic-deal-client/DealClient.sol) deal proposal contract at

# commands

- `datadaotool submit-deal` -- submits a deal proposal on chain for SP's that have `AcceptContractDeals: true` config enables and are monitoring the Filecoin chain for deals
- `datadaomonitor run` -- simple monitor tool to listen for `DealProposalCreate(bytes32)` events emitted on chain

# examples

```
datadaomonitor run

datadaotool submit-deal --contract=0xF54d13E2AFb212Da562E3ab4BFB4372c63913eb6 --commp=bafk2bzacec3jst4tkh424chatp273o6rxvipfg54kphd56gaxobpcdtr2sgco --piece-size=2048 --car-size=1175 --payload-cid=bafk2bzacec3jst4tkh424chatp273o6rxvipfg54kphd56gaxobpcdtr2sgco

```
