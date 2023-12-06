package polygon

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"polygon-token-oracle-backend/logic/config"
)

var (
	DefaultClient *Client
)

func init() {
	client, err := NewClient(config.Get().Polygon.Endpoint)
	if err != nil {
		panic(err)
	}

	DefaultClient = client
}

type Client struct {
	client *ethclient.Client
}

func NewClient(rawurl string) (*Client, error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.client.SuggestGasPrice(ctx)
}

func (c *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return c.client.EstimateGas(ctx, msg)
}

func (c *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.client.HeaderByNumber(ctx, number)
}

func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.client.BlockNumber(ctx)
}

func (c *Client) ChainID(ctx context.Context) (*big.Int, error) {
	return c.client.ChainID(ctx)
}

func (c *Client) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return c.client.FilterLogs(ctx, query)
}

func (c *Client) Nonce(ctx context.Context, fromAddress common.Address) (uint64, error) {
	return c.client.PendingNonceAt(context.Background(), fromAddress)
}

func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.client.NetworkID(context.Background())
}

func (c *Client) TransactionReceipt(txHash common.Hash) (*types.Receipt, error) {
	return c.client.TransactionReceipt(context.Background(), txHash)
}

func (c *Client) TransactionByHash(txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return c.client.TransactionByHash(context.Background(), txHash)
}

func (c *Client) BlockByHash(hash common.Hash) (*types.Block, error) {
	return c.client.BlockByHash(context.Background(), hash)
}

func (c *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.client.SubscribeFilterLogs(ctx, q, ch)
}

func (c *Client) Client() *ethclient.Client {
	return c.client
}
