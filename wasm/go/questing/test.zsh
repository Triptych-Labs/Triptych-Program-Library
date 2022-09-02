pushd ./integrations/questing

export ORACLE="GbfoTncFrg8PxS2KY9mmCHz73Bv9cXUxsr7Q66y5SUDo"
export HOLDER="DhHMVCTDvdcVYokuskrCAnM7LSzivzBsze5SxUCEe26z"

export PATH=$PATH:/usr/local/go/misc/wasm/
GOOS=js GOARCH=wasm go test -v

popd
