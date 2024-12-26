package order

import (
	"net/http"

	"github.com/fyerfyer/gozero-ecommerce/ecommerce/application/internal/logic/order"
	"github.com/fyerfyer/gozero-ecommerce/ecommerce/application/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CancelOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := order.NewCancelOrderLogic(r.Context(), svcCtx)
		err := l.CancelOrder()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
