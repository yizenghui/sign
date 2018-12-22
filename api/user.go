package api

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	cpi "github.com/yizenghui/sign/core"
	"github.com/yizenghui/sign/db"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	OpenID  string `json:"open_id"`
	Code    string `json:"code"`
	Session string `json:"session"`
	jwt.StandardClaims
}

// GetToken 获取 jwt token
func GetToken(c echo.Context) error {
	code := c.QueryParam("code")
	ret, _ := cpi.GetOpenID(code)
	if code != "" && ret.OpenID != "" {

		log.Println(ret)
		// Set custom claims
		claims := &JwtCustomClaims{
			ret.OpenID,
			code,
			ret.SessionKey,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// 获取签名里面的信息
func getOpenID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims.OpenID
}

// 获取用户信息
func getUser(openID string) (*db.Fans, error) {
	fans, err := cpi.GetFansByOpenID(openID)
	return fans, err
}

// Crypt 解密同步用户信息
func Crypt(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	type Data struct {
		Ed string `json:"ed"`
		Iv string `json:"iv"`
	}
	data := new(Data)
	if err := c.Bind(data); err != nil {
		return err
	}

	// sessionKey := c.QueryParam("sk")
	// encryptedData := c.QueryParam("ed")
	// iv := c.QueryParam("iv")
	// log.Println(claims.Code, claims.Session, encryptedData, iv)
	// ret, _ := cpi.GetCryptData(claims.Session, encryptedData, iv)

	ret, _ := cpi.GetCryptData(claims.Session, data.Ed, data.Iv)
	return c.JSON(http.StatusOK, ret)
}

// CheckOpenID 解密同步用户信息
func CheckOpenID(c echo.Context) error {

	return c.JSON(http.StatusOK, getOpenID(c))
}
