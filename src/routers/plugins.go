package routers

import (
	"net/http"

	"github.com/blackcrw/wprecon-api/src/controllers"
)

var routersPlugin = []Router{
	{
		URI:          "/v1/api/vulnerable/plugin",
		Method:       http.MethodGet,
		Func:         controllers.GetPluginNameList,
		RequiredAuth: false,
	},
	{
		URI:          "/v1/api/vulnerable/plugin/{Plugin}/{Version}",
		Method:       http.MethodGet,
		Func:         controllers.GetPluginVuln,
		RequiredAuth: false,
	},
}
