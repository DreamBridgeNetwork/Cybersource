module github.com/DreamBridgeNetwork/Go-Cybersource/cmd/httpserver

go 1.18

require (
	github.com/DreamBridgeNetwork/Go-Cybersource v0.0.0-20220910172943-373af88c681f
	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/cybersourcerest v0.0.0-20220910172943-373af88c681f
	github.com/DreamBridgeNetwork/Go-Utils v0.0.0-20220910173339-3b42edec4cdf
)

require (
	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/utils v0.0.0-20220910161331-81445e38ada5 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace (
	github.com/DreamBridgeNetwork/Go-Cybersource v0.0.0-20220910172943-373af88c681f => ../../
	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/cybersourcerest v0.0.0-20220910172943-373af88c681f => ../../pkg/cybersourcerest
	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/utils v0.0.0-20220910161331-81445e38ada5 => ../../pkg/utils
	github.com/DreamBridgeNetwork/Go-Utils v0.0.0-20220910173339-3b42edec4cdf => ../../../Go-Utils
)
