module github.com/profects/zendesk-go

go 1.12

require (
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/hellofresh/zendesk-go v0.0.0-20160531211648-717dd547778c
	github.com/ttacon/builder v0.0.0-20170518171403-c099f663e1c2 // indirect
	github.com/ttacon/libphonenumber v1.0.1
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	gopkg.in/resty.v0 v0.4.1
	gopkg.in/resty.v1 v1.12.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/micro/go-micro => github.com/micro/go-micro v1.2.0

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0
