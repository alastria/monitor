package middleware

import (
	hydra "github.com/ory-am/hydra/sdk"

	"net/http"
	"strings"

	"github.com/astaxie/beego"
	beegoContext "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

var (
	hc *hydra.Client
)

func Init(hydraClient *hydra.Client) {
	hc = hydraClient
}

func AccessTokenFromRequest(req *http.Request) string {
	auth := req.Header.Get("Authorization")
	split := strings.SplitN(auth, " ", 2)
	if len(split) != 2 || !strings.EqualFold(split[0], "Bearer") {
		// Empty string returned if there's no such parameter
		err := req.ParseForm()
		if err != nil {
			return ""
		}
		return req.Form.Get("access_token")
	}
	return split[1]
}

func ScopesRequired(scopes ...string) beego.FilterFunc {
	return func(c *beegoContext.Context) {
		logs.Trace("ScopesRequired: ", scopes, c)
		// Esto es necesario para pre-venir verificaciones en cascada.
		if len(c.Request.Header.Get("API-Authenticate")) > 0 {
			logs.Trace("Ya lo he mirado")
			return
		}
		token := AccessTokenFromRequest(c.Request)
		logs.Debug("Token: ", token)
		if len(token) == 0 {
			logs.Trace("No hay cabecera authorization.")
			c.Abort(302, "Operación no permitida. NO se ha suministrado el token.")
		}
		ctx, err := hc.Introspection.IntrospectToken(c.Request.Context(), token, scopes...)
		logs.Debug("Introspect: ", ctx, err)
		if err != nil || (ctx != nil && !ctx.Active) {
			logs.Debug("Introspect: ", ctx, err)
			c.Abort(302, "Operación no permitida. Token incorrecto o caducado.")
		}

		c.Request.Header.Set("API-Authenticate", "Yes")
	}
}
