# datadaotool

This repo/tool is supposed to interact with the proof-of-concept [DealClient.sol](https://github.com/filecoin-project/fevm-hardhat-kit/pull/90) deal proposal contract

# commands

- `datadaotool submit-deal` -- submits a deal proposal on chain for SP's that have `AcceptContractDeals: true` config enables and are monitoring the Filecoin chain for deals

# examples

```
$ export PRIVATE_KEY=aef394....

$ datadaotool submit-deal \
  --contract=0x92780Bc6E860a935307f000eE90C5AA1878454D4 \
  --commp=baga6ea4seaqlnvynus6vwba7rob4tuslkutuvl6zuon46cfla4ebkxcn3yxd2ji \
  --piece-size=262144 \
  --car-size=236445 \
  --start-epoch=146694 \
  --storage-price=0 \
  --payload-cid=bafk2bzaceanzppnlffioby4nac2hhjmrstzntqie3oid4ovq6zu4qhhjs4bvy \
  --client=t410fsj4axrximcutkmd7aahosdc2ugdyivguhncj5ri \
  --skip-ipni-announce=true \
  --location_ref=http://webserver:24005/screenshot.car
```
