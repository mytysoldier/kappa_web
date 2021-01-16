module github.com/mytysoldier/kappa_web

go 1.15

replace github.com/mytysoldier/kappa_web/funcs/todo => ./funcs

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/mytysoldier/kappa_web/funcs/todo v0.0.0-00010101000000-000000000000
)
