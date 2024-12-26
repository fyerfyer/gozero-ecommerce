package cart

import (
	"net/http"

	"github.com/fyerfyer/gozero-ecommerce/ecommerce/application/internal/logic/cart"
	"github.com/fyerfyer/gozero-ecommerce/ecommerce/application/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSelectedItemsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewGetSelectedItemsLogic(r.Context(), svcCtx)
		resp, err := l.GetSelectedItems()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
