package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

type TokenDetails struct {
	AccessToken string
	RefreshToken string
	AtExpiredAt int64
	RfExpiredAt int64
}

// 產生 jwt token
func CreateJwtToken(userid int) (string, error){
	claims := jwt.MapClaims{
		"user_id": userid,
		"exp": time.Now().Add(time.Minute * 15).Unix(), // 過期時間
		"authorize": true,
	}

	sign := []byte("todolist") // 簽名的 key
	at := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	token, err := at.SignedString(sign)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(r *http.Request) (*jwt.Token, error) {

	// 判斷是否有帶入 authorized
	_, ok := r.Header["Authorization"]
	if !ok {
		return nil, fmt.Errorf("UnAuthorized")
	}

	// 取得 token 解析
	auth := r.Header.Get("Authorization")
	strarray := strings.Split(auth, "Bearer ")

	// 判斷是否有帶入 Beaer
	if len(strarray) < 2 {
		return nil, fmt.Errorf("UnAuthorized")
	}

	token := strarray[1]
	
	t, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unepxected method")
		}
		// 解析成功
		return []byte("todolist"), nil
	})

	if err != nil {
		return nil, err
	}

	return t, nil
}