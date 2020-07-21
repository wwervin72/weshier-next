package token

import (
	"fmt"
	"time"
	"weshierNext/pkg/errno"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// JWTClaims is the context of the JSON web token
type JWTClaims struct {
	ID uint64 `json:"id"`
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// make sure the `alg` is what we except.
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errno.ErrTokenInvalid
		}
		return []byte(secret), nil
	}
}

// Parse parse validates the token with the specified secret
// and return the JWTClaims if the token was valid.
func Parse(tokenString string, secret string) (*JWTClaims, error) {
	ctx := &JWTClaims{}
	// parse token
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	// parse error
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		// read the token if it's valid
		ctx.ID = uint64(claims["id"].(float64))
		return ctx, nil
	} else {
		return ctx, errno.ErrTokenInvalid
	}
}

// ParseRequest gets the token from the header and
// pass it to the parse function to parses the token
func ParseRequest(c *gin.Context) (*JWTClaims, error) {
	header := c.Request.Header.Get("Authorization")
	// load the jwt secret from config
	secret := viper.GetString("jwt.secret")
	if len(header) == 0 {
		return nil, errno.ErrTokenEmpty
	}
	var t string
	// parse the header o get the token part
	fmt.Sscanf(header, "Bearer %s", &t)
	ctx, err := Parse(t, secret)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

// Sign signs the JWTClaims with the specified secret
func Sign(ctx *gin.Context, c JWTClaims, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("jwt.secret")
	}
	var nowSecond = time.Now().Unix()
	// the token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": nowSecond,
		// token createdAt
		"iat": nowSecond,
		"id":  c.ID,
	})
	// sign the token with specified secret
	tokenString, err = token.SignedString([]byte(secret))
	return
}
