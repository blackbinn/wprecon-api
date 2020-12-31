package routers

import (
	"net/http"

	"github.com/blackcrw/wprecon-api/pkg/controllers"
)

var routersPlugin = []Router{
	{
		URI:          "/vulnerable/plugin",
		Method:       http.MethodGet,
		Func:         controllers.GetPluginNameList,
		RequiredAuth: false,
	},
	{
		URI:          "/vulnerable/plugin/{Plugin}/{Version}",
		Method:       http.MethodGet,
		Func:         controllers.GetPluginVuln,
		RequiredAuth: false,
	},
}
