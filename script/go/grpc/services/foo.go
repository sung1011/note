package service

import (
	context "context"
)

type Prod struct{}

func (this *Prod) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{Stock: 123}, nil
}

func (this *Prod) mustEmbedUnimplementedProdServiceServer() {}
