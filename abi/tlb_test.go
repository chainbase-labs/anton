package abi_test

import (
	"encoding/json"
	"testing"

	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"

	"github.com/stretchr/testify/require"

	"github.com/tonindexer/anton/abi"
)

func makeOperationDesc(t *testing.T, x any) string {
	d, err := abi.NewOperationDesc(x)
	require.Nil(t, err)

	n, err := d.New()
	require.Nil(t, err)

	nd, err := abi.NewOperationDesc(n)
	nd.Name = d.Name
	require.Nil(t, err)
	require.Equal(t, d, nd)

	j, err := json.Marshal(nd)
	require.Nil(t, err)

	return string(j)
}

func loadOperation(t *testing.T, schema, boc string) string {
	var d abi.OperationDesc

	err := json.Unmarshal([]byte(schema), &d)
	require.Nil(t, err)

	op, err := d.New()
	require.Nil(t, err)

	c, err := cell.FromBOC(mustBase64(t, boc))
	require.Nil(t, err)

	err = tlb.LoadFromCell(op, c.BeginParse())
	require.Nil(t, err)

	j, err := json.Marshal(op)
	require.Nil(t, err)

	return string(j)
}
