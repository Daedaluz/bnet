package bnet

import(
	auth "golang.org/x/oauth2"
)

var broken = []string{
	"https://eu.battle.net/",
	"https://us.battle.net/",
	"https://kr.battle.net/",
	"https://tw.battle.net/",
	"https://www.battlenet.com.cn/",
}

var EndpointEU = auth.Endpoint {
	AuthURL: "https://eu.battle.net/oauth/authorize",
	TokenURL: "https://eu.battle.net/oauth/token",
}

var EndpointUS = auth.Endpoint {
	AuthURL: "https://us.battle.net/oauth/authorize",
	TokenURL: "https://us.battle.net/oauth/token",
}

var EndpointKR = auth.Endpoint {
	AuthURL: "https://kr.battle.net/oauth/authorize",
	TokenURL: "https://kr.battle.net/oauth/token",
}

var EndpointTW = auth.Endpoint {
	AuthURL: "https://tw.battle.net/oauth/authorize",
	TokenURL: "https://tw.battle.net/oauth/token",
}

func init() {
	for _, broke := range broken {
		auth.RegisterBrokenAuthHeaderProvider(broke)
	}
}
