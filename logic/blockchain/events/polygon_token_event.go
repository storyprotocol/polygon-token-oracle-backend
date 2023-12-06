package events

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"polygon-token-oracle-backend/internal/log"
	"polygon-token-oracle-backend/logic/blockchain/actions"
	"polygon-token-oracle-backend/logic/blockchain/eth/contracts/polygon_token"
	"polygon-token-oracle-backend/logic/polygon"
	"sync/atomic"
)

type PolygonTokenEvent struct {
	ev *Events

	Addr       common.Address
	BlockStart int64

	isStart atomic.Bool
	isClose atomic.Bool
	ctx     context.Context
	cancel  context.CancelFunc

	sol *polygon_token.PolygonToken
}

func NewPolygonTokenEvent(ev *Events, addr string) (e *PolygonTokenEvent, err error) {
	e = &PolygonTokenEvent{
		ev:         ev,
		Addr:       common.HexToAddress(addr),
		BlockStart: -1,
	}

	e.ctx, e.cancel = context.WithCancel(context.Background())
	e.sol, err = polygon_token.NewPolygonToken(e.Addr, e.ev.ethClient.Client())
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *PolygonTokenEvent) Start(rt ratelimit.Limiter) {
	if e.isStart.CompareAndSwap(false, true) {
		polygonTokenChan := make(chan *polygon_token.PolygonTokenPolygonTokenBalanceRequest, 100)
		go func() {
			defer close(polygonTokenChan)

			e.ev.filterRange(e.ctx, rt, e.BlockStart, func(watchStart, watchEnd uint64) error {
				fmt.Printf("watchStart: %d\n", watchStart)

				polygonTokenEvents, err := e.sol.FilterPolygonTokenBalanceRequest(&bind.FilterOpts{Start: watchStart, End: &watchEnd}, nil, nil)
				if err != nil {
					return err
				}
				defer polygonTokenEvents.Close()

				for polygonTokenEvents.Next() {
					if polygonTokenEvents.Error() != nil {
						return polygonTokenEvents.Error()
					}
					polygonTokenChan <- polygonTokenEvents.Event
				}

				return nil
			})
		}()

		go func() {
			e.processTransfers(polygonTokenChan)
			e.isClose.Store(true)
		}()
	}
}

func (e *PolygonTokenEvent) processTransfers(transferChan chan *polygon_token.PolygonTokenPolygonTokenBalanceRequest) {
	for transfer := range transferChan {
		err := e.processPolygonToken(transfer)
		if err != nil {
			log.GetLogger().Warn("[PolygonTokenBalanceRequest]", zap.Any("event", transfer), zap.Error(err))
		}
	}
}

func (e *PolygonTokenEvent) processPolygonToken(ev *polygon_token.PolygonTokenPolygonTokenBalanceRequest) error {
	balance, err := polygon.GetTokenBalance(ev.TokenAddress, ev.TokenOwnerAddress)
	if err != nil {
		return err
	}

	// push oracle data
	txHash, err := actions.PushPolygonTokenBalance(ev.CallbackAddr, ev.RequestId, balance)
	if err != nil {
		return err
	}

	log.GetLogger().Info("push oracle data succeed.", zap.String("txHash", txHash))
	return nil
}
