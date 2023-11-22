package contexts

import (
	"context"

	"github.com/teq-quocbang/arrows/delivery/http/auth"
	"github.com/teq-quocbang/arrows/util/token"
)

func GetUserPrincipleByContext(ctx context.Context) *token.JWTClaimCustom {
	reply := ctx.Value(auth.UserPrincipleKey)
	if values, ok := reply.(*token.JWTClaimCustom); ok {
		return values
	}
	return nil
}
