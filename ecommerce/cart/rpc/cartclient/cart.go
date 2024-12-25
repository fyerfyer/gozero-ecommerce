// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: cart.proto

package cartclient

import (
	"context"

	"github.com/fyerfyer/gozero-ecommerce/ecommerce/cart/rpc/cart"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddItemRequest           = cart.AddItemRequest
	AddItemResponse          = cart.AddItemResponse
	CartItem                 = cart.CartItem
	CheckStockRequest        = cart.CheckStockRequest
	CheckStockResponse       = cart.CheckStockResponse
	ClearCartRequest         = cart.ClearCartRequest
	ClearCartResponse        = cart.ClearCartResponse
	GetCartRequest           = cart.GetCartRequest
	GetCartResponse          = cart.GetCartResponse
	GetSelectedItemsRequest  = cart.GetSelectedItemsRequest
	GetSelectedItemsResponse = cart.GetSelectedItemsResponse
	RemoveItemRequest        = cart.RemoveItemRequest
	RemoveItemResponse       = cart.RemoveItemResponse
	SelectAllRequest         = cart.SelectAllRequest
	SelectAllResponse        = cart.SelectAllResponse
	SelectItemRequest        = cart.SelectItemRequest
	SelectItemResponse       = cart.SelectItemResponse
	UnselectAllRequest       = cart.UnselectAllRequest
	UnselectAllResponse      = cart.UnselectAllResponse
	UnselectItemRequest      = cart.UnselectItemRequest
	UnselectItemResponse     = cart.UnselectItemResponse
	UpdateItemRequest        = cart.UpdateItemRequest
	UpdateItemResponse       = cart.UpdateItemResponse

	Cart interface {
		// 购物车操作
		AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
		UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
		RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*RemoveItemResponse, error)
		GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
		ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error)
		// 商品选择
		SelectItem(ctx context.Context, in *SelectItemRequest, opts ...grpc.CallOption) (*SelectItemResponse, error)
		UnselectItem(ctx context.Context, in *UnselectItemRequest, opts ...grpc.CallOption) (*UnselectItemResponse, error)
		SelectAll(ctx context.Context, in *SelectAllRequest, opts ...grpc.CallOption) (*SelectAllResponse, error)
		UnselectAll(ctx context.Context, in *UnselectAllRequest, opts ...grpc.CallOption) (*UnselectAllResponse, error)
		// 结算相关
		GetSelectedItems(ctx context.Context, in *GetSelectedItemsRequest, opts ...grpc.CallOption) (*GetSelectedItemsResponse, error)
		CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error)
	}

	defaultCart struct {
		cli zrpc.Client
	}
)

func NewCart(cli zrpc.Client) Cart {
	return &defaultCart{
		cli: cli,
	}
}

// 购物车操作
func (m *defaultCart) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.AddItem(ctx, in, opts...)
}

func (m *defaultCart) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.UpdateItem(ctx, in, opts...)
}

func (m *defaultCart) RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*RemoveItemResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.RemoveItem(ctx, in, opts...)
}

func (m *defaultCart) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.GetCart(ctx, in, opts...)
}

func (m *defaultCart) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.ClearCart(ctx, in, opts...)
}

// 商品选择
func (m *defaultCart) SelectItem(ctx context.Context, in *SelectItemRequest, opts ...grpc.CallOption) (*SelectItemResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.SelectItem(ctx, in, opts...)
}

func (m *defaultCart) UnselectItem(ctx context.Context, in *UnselectItemRequest, opts ...grpc.CallOption) (*UnselectItemResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.UnselectItem(ctx, in, opts...)
}

func (m *defaultCart) SelectAll(ctx context.Context, in *SelectAllRequest, opts ...grpc.CallOption) (*SelectAllResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.SelectAll(ctx, in, opts...)
}

func (m *defaultCart) UnselectAll(ctx context.Context, in *UnselectAllRequest, opts ...grpc.CallOption) (*UnselectAllResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.UnselectAll(ctx, in, opts...)
}

// 结算相关
func (m *defaultCart) GetSelectedItems(ctx context.Context, in *GetSelectedItemsRequest, opts ...grpc.CallOption) (*GetSelectedItemsResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.GetSelectedItems(ctx, in, opts...)
}

func (m *defaultCart) CheckStock(ctx context.Context, in *CheckStockRequest, opts ...grpc.CallOption) (*CheckStockResponse, error) {
	client := cart.NewCartClient(m.cli.Conn())
	return client.CheckStock(ctx, in, opts...)
}