package jwts

import (
	"fmt"
	"github.com/adminwjp/infrastructure-go/utils"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)
var JwtInstance=&JwtHelper{}
//自定义一个字符串
var Jwtkey = []byte("wjp123456WJP")
var JwtRefreshkey = []byte("---wjp123456WJP---")

type Claims struct {
	UserId string `json:"u_id"`
	Phone string `json:"p"`
	Email string `json:"e"`
	UserName string `json:"un"`
	NickName string `json:"nn"`
	Pic string `json:"p"`
	Ip string `json:"i"`
	Token string `json:"t"`
	jwt.StandardClaims
}
type RefreshClaims struct {
	UserId string `json:"u_id"`
	Token string `json:"t"`
	Ip string `json:"i"`
	jwt.StandardClaims
}
type JwtHelper struct {

}

type JwtDto struct {
	UserId string
	Phone string
	Email string
	UserName string
	NickName string
	Pic string
	Ip string
	ExpiresAt int64
	CreateAt int64
	Issuer string
	Subject string
}

//颁发token token refresh
func (j *JwtHelper) CreateTokenAndRefreshToken(dto JwtDto)(string,string,error) {
	if dto.Phone == "" {
		dto.Phone = "pe"
	}
	if dto.Email == "" {
		dto.Email = "ee"
	}
	if dto.UserName == "" {
		dto.Email = "ue"
	}
	if dto.NickName == "" {
		dto.Email = "用户"
	}
	token1 := utils.SeurityInstance.AesEncrypt(dto.UserId + " " +
		strconv.FormatInt(dto.CreateAt, 10) + " " +
		strconv.FormatInt(dto.ExpiresAt, 10))
	claims := &Claims{
		UserId:   dto.UserId,
		Phone:    utils.SeurityInstance.AesEncrypt(dto.Phone),
		Email:    utils.SeurityInstance.AesEncrypt(dto.Email),
		UserName: utils.SeurityInstance.AesEncrypt(dto.UserName),
		NickName: dto.NickName,
		Pic:      dto.Pic,
		Token:    token1,
		Ip:       utils.SeurityInstance.AesEncrypt(dto.Ip),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dto.ExpiresAt, //过期时间
			IssuedAt:  dto.CreateAt,
			Issuer:    dto.Issuer,  // 签名颁发者
			Subject:   dto.Subject, //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(Jwtkey)
	if err != nil {
		fmt.Printf("create token err,%s", err)
		return "", "", err
	}
	refershClaims := &RefreshClaims{
		UserId: dto.UserId,
		Token:  token1,
		Ip:     utils.SeurityInstance.AesEncrypt(dto.Ip),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dto.ExpiresAt, //过期时间
			IssuedAt:  dto.CreateAt,
			Issuer:    dto.Issuer,  // 签名颁发者
			Subject:   dto.Subject, //签名主题
		},
	}
	refershToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refershClaims)
	// fmt.Println(token)
	refershTokenString, err := refershToken.SignedString(JwtRefreshkey)
	if err != nil {
		fmt.Printf("create refresh token err,%s", err)
		return "", "", err
	}
	return tokenString, refershTokenString, err
}
func (j *JwtHelper) RefreshToken(dto JwtDto)(string,error) {
	if dto.Phone==""{
		dto.Phone="pe"
	}
	if dto.Email==""{
		dto.Email="ee"
	}
	if dto.UserName==""{
		dto.Email="ue"
	}
	if dto.NickName==""{
		dto.Email="用户"
	}
	token1 := utils.SeurityInstance.AesEncrypt(dto.UserId + " " +
		strconv.FormatInt(dto.CreateAt, 10) + " " +
		strconv.FormatInt(dto.ExpiresAt, 10))
	claims := &Claims{
		UserId: dto.UserId,
		Phone: utils.SeurityInstance.AesEncrypt(dto.Phone),
		Email: utils.SeurityInstance.AesEncrypt(dto.Email),
		UserName: utils.SeurityInstance.AesEncrypt(dto.UserName),
		NickName: dto.NickName,
		Pic: dto.Pic,
		Token:  token1,
		Ip: utils.SeurityInstance.AesEncrypt(dto.Ip),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dto.ExpiresAt, //过期时间
			IssuedAt:  dto.CreateAt,
			Issuer:    dto.Issuer,  // 签名颁发者
			Subject:   dto.Subject, //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(Jwtkey)
	if err != nil {
		fmt.Printf("create token err,%s",err)
		return "", err
	}
	return tokenString, err
}
//解析token


func (j *JwtHelper) ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return Jwtkey, nil
	})
	return token, claims, err
}

func (j *JwtHelper) ParseRefreshToken(tokenString string) (*jwt.Token, *RefreshClaims, error) {
	claims := &RefreshClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtRefreshkey, nil
	})
	return token, claims, err
}