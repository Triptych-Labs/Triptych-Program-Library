#!/bin/zsh

GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o swapper.wasm
# cp questing.wasm ~/creatures/Triptych-Labs-Official-Site/public/questing.wasm
 cp swapper.wasm ~/triptych/Trriptych-Boiler-Room/public/swapper.wasm

#scp someplace.wasm ddigiacomo@10.145:/noshit/triptych_labs/homepage/public
#scp someplace.wasm ddigiacomo@10.145:/noshit/triptych_labs/homepage/public/marketplace

#scp someplace.wasm ddigiacomo@10.145:/noshit/cdn
