module github.com/Go-Cybersource

go 1.19

require (
	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/cybersourcerest v0.0.0-20220926123211-f0ca62a9b43e
	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/utils v0.0.0-20220926123211-f0ca62a9b43e
	github.com/DreamBridgeNetwork/Go-Utils v0.0.0-20220917200124-80c5468fb864
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
)

require golang.org/x/text v0.3.7 // indirect

//replace (
//	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/cybersourcerest v0.0.0-20220926123211-f0ca62a9b43e => ./pkg/cybersourcerest
//	github.com/DreamBridgeNetwork/Go-Cybersource/pkg/utils v0.0.0-20220926123211-f0ca62a9b43e => ./pkg/utils
//	github.com/DreamBridgeNetwork/Go-Utils v0.0.0-20220917200124-80c5468fb864 => ../Go-Utils
//)