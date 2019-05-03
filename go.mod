module github.com/rxc-wujianhua/go-utils

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190426145343-a29dc8fdc734
	golang.org/x/image => github.com/golang/image v0.0.0-20190501045829-6d32002ffd75
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190501004415-9ce7a6920f09
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190502143002-c0b26311cb83
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190501045030-23463209683d
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
)

require (
	github.com/go-log/log v0.1.0
	github.com/go-xorm/core v0.6.2
	github.com/go-xorm/xorm v0.7.1
	github.com/joho/godotenv v1.3.0
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/rs/xid v1.2.1
)
