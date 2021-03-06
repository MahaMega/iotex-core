// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package action

import (
	"math/big"
	"testing"

	"github.com/iotexproject/iotex-core/address"
	"github.com/iotexproject/iotex-core/test/testaddress"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeposit(t *testing.T) {
	t.Parallel()

	addr1 := address.New(1, testaddress.Addrinfo["producer"].PublicKey[:]).Bech32()
	addr2 := address.New(2, testaddress.Addrinfo["alfa"].PublicKey[:]).Bech32()

	assertDeposit := func(deposit *Deposit) {
		require.NotNil(t, deposit)
		assert.Equal(t, uint64(1), deposit.Nonce())
		assert.Equal(t, big.NewInt(1000), deposit.Amount())
		assert.Equal(t, addr1, deposit.Sender())
		assert.Equal(t, addr2, deposit.Recipient())
		assert.Equal(t, uint64(10), deposit.GasLimit())
		assert.Equal(t, big.NewInt(100), deposit.GasPrice())
	}

	deposit1 := NewDeposit(
		1,
		big.NewInt(1000),
		addr1,
		addr2,
		10,
		big.NewInt(100),
	)
	assertDeposit(deposit1)

	data := deposit1.Proto()
	require.NotNil(t, data)
	var deposit2 Deposit
	deposit2.LoadProto(data)
	assertDeposit(&deposit2)
}
