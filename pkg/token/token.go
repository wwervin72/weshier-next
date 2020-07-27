package token

import (
	"encoding/json"
	"fmt"
	"time"
	"weshierNext/pkg/errno"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// JWTClaims is the context of the JSON web token
type JWTClaims struct {
	ID       uint64 `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
	URL      string `json:"url"`
	Phone    uint64 `json:"phone"`
	Role     string `json:"role"`
	Age      uint8  `json:"age"`
	Status   uint8  `json:"status"`
	Resume   uint8  `json:"resume"`
	AuthID   uint64 `json:"authId"`
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
		d, err := json.Marshal(claims)
		if err != nil {
			return ctx, errno.InternalServerError
		}
		err = json.Unmarshal(d, ctx)
		if err != nil {
			return ctx, errno.InternalServerError
		}
		return ctx, nil
	} else {
		return ctx, errno.ErrTokenInvalid
	}
}

// ParseRequest gets the token from the header and
// pass it to the parse function to parses the token
func ParseRequest(c *gin.Context) (*JWTClaims, string, error) {
	header := c.Request.Header.Get("Authorization")
	// load the jwt secret from config
	secret := viper.GetString("jwt.secret")
	if len(header) == 0 {
		return nil, "", errno.ErrTokenEmpty
	}
	var t string
	// parse the header o get the token part
	fmt.Sscanf(header, "Bearer %s", &t)
	ctx, err := Parse(t, secret)
	if err != nil {
		return nil, "", err
	}
	return ctx, t, nil
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
		"iat":      nowSecond,
		"id":       c.ID,
		"userName": c.UserName,
		"email":    c.Email,
		"nickName": c.NickName,
		"bio":      c.Bio,
		"avatar":   c.Avatar,
		"url":      c.URL,
		"phone":    c.Phone,
		"role":     c.Role,
		"age":      c.Age,
		"status":   c.Status,
		"resume":   c.Resume,
		"authId":   c.AuthID,
	})
	// sign the token with specified secret
	tokenString, err = token.SignedString([]byte(secret))
	return
}
