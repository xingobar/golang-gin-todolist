package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"net/http"
	"os"
	"strings"
	"time"
)

type TokenDetails struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpiredAt int64 `json:"at_expired_at"`
	RfExpiredAt int64	`json:"rf_expired_at"`
	AccessUid string	`json:"access_uid"`
	RefreshUid string	`json:"refresh_uid"`
}

// 產生 jwt token
func CreateJwtToken(userid int) (*TokenDetails, error){

	td := &TokenDetails{}
	td.AtExpiredAt = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUid = uuid.NewV4().String()

	// 1個禮拜
	td.RfExpiredAt = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUid = uuid.NewV4().String()

	// 解碼後的 map
	claims := jwt.MapClaims{
		"user_id": userid,
		"exp": td.AtExpiredAt, // 過期時間
		"authorize": true,
		"access_uid": td.AccessUid,
	}

	// 設定 access token
	sign := []byte(os.Getenv("JWT_SECRET")) // 簽名的 key
	at := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	token, err := at.SignedString(sign)
	if err != nil {
		return nil, err
	}
	td.AccessToken = token

	// 設定 refresh token
	refreshClaim := jwt.MapClaims{
		"user_id": userid,
		"exp": td.RfExpiredAt,
		"refresh_uid": td.RefreshUid,
	}
	refreshSign := []byte(os.Getenv("REFRESH_JWT"))
	rf := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)

	rtoken, err := rf.SignedString(refreshSign)
	if err != nil {
		return nil, err
	}
	td.RefreshToken = rtoken

	return td, nil
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
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("========= error ========")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("======== t =========")
	fmt.Println(t)
	return t, nil
}