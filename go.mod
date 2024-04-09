module github.com/xiaokatech/go-lab

go 1.22.0

// replace is a hack for local development
replace github.com/xiaokatech/go-lab-package v0.0.0 => ../go-lab-package

require github.com/xiaokatech/go-lab-package v0.0.0

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/wagslane/go-tinytime v0.0.2 // indirect
)
