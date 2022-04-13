package resolver

import (
	"context"
	"encoding/json"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/graphql/utils"
	"github.com/graph-gophers/graphql-go"
)

type FilterPointResolver struct {
	q dia.FilterPoint
}

type SignedFilterPoint struct {
	Asset   dia.Asset
	Value   float64
	Name    string
	Time    time.Time
	Symbol  string
	Address string
}

func (qr *FilterPointResolver) Name(ctx context.Context) (*string, error) {
	return &qr.q.Name, nil
}

func (qr *FilterPointResolver) Symbol(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Symbol, nil
}

func (qr *FilterPointResolver) Time(ctx context.Context) (*graphql.Time, error) {
	return &graphql.Time{Time: qr.q.Time}, nil
}

func (qr *FilterPointResolver) Value(ctx context.Context) (*float64, error) {
	return &qr.q.Value, nil
}
func (qr *FilterPointResolver) Address(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Address, nil
}

func (qr *FilterPointResolver) Blockchain(ctx context.Context) (*string, error) {
	return &qr.q.Asset.Blockchain, nil
}
func (qr *FilterPointResolver) Sign(ctx context.Context) (*string, error) {
	su, err := utils.NewSignatureUtil()
	if err != nil {
		return nil, err
	}

	sfp := SignedFilterPoint{Symbol: qr.q.Asset.Symbol, Address: qr.q.Asset.Address, Value: qr.q.Value, Time: qr.q.Time, Name: qr.q.Name}

	messageTosign, err := json.Marshal(sfp)
	if err != nil {
		return nil, err
	}

	signedMessage := su.Sign(messageTosign)

	return &signedMessage, nil
}
