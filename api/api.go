package api

import (
	"encoding/json"
	"net/http"
)

type CoinBaseParams struct {
	Username string
}

type CoinBaseResponse struct {
	Code    int
	Balance int64
}

type Error struct {
	Code    int
	Message string
}
