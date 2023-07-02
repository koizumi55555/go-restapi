package usecase

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"koizumi55555/go-restapi/src/common"
	"koizumi55555/go-restapi/src/entitiy"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type UserUsecase struct {
	PostgresIf PostgresIf
}

type ServerUsecase struct {
	ClientID     string
	ClientSecret string
}

func NewUserUsecase(pif PostgresIf) UserUsecase {
	return UserUsecase{
		PostgresIf: pif,
	}
}

func NewServerUsecase(clientID, clientSecret string) ServerUsecase {
	return ServerUsecase{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (uuc UserUsecase) GetUser(id string) (user entitiy.User, err error) {
	return uuc.PostgresIf.GetUserDB(id)
}

func (uuc UserUsecase) DeleteUser(id string) (err error) {

	getUser, err := uuc.PostgresIf.GetUserDB(id)
	if err != nil {
		return err
	}

	return uuc.PostgresIf.DeleteUserDB(getUser.ID)
}

func (uuc UserUsecase) UpdateUser(updateUser entitiy.User) (user entitiy.User, err error) {

	getUser, err := uuc.PostgresIf.GetUserDB(updateUser.ID)
	if err != nil {
		return getUser, err
	}

	return uuc.PostgresIf.UpdateUserDB(getUser, updateUser)
}

func (uuc UserUsecase) CreateUser(createUser entitiy.User) (user entitiy.User, err error) {
	return uuc.PostgresIf.CreateUserDB(createUser)
}

func (uuc UserUsecase) ListUsers() (user []entitiy.User, err error) {
	return uuc.PostgresIf.ListUsersDB()
}

// URLの生成
func (s *ServerUsecase) CreateAuthorizationRequestURL() (*url.URL, error) {
	u, err := url.Parse(common.AuthorizationEndpoint)
	if err != nil {
		return nil, err
	}

	// URLのクエリパラメータを取得
	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", s.ClientID)
	if common.RedirectURI != "" {
		q.Set("redirect_uri", common.RedirectURI)
	}
	if len(common.Scopes) > 0 {
		q.Set("scope", strings.Join(common.Scopes, " "))
	}
	q.Set("prompt", "consent")
	// 更新されたクエリパラメータをエンコードし、元のURLに戻す。
	u.RawQuery = q.Encode()

	return u, nil
}

// アクセストークンを取得する
func (s *ServerUsecase) Exchange(ctx context.Context, code string) (*entitiy.Token, error) {
	v := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {common.RedirectURI},
		"client_id":    {s.ClientID},
	}

	// HTTPリクエストを作成
	req, err := http.NewRequest("POST", common.TokenEndpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	// ヘッダーと認証情報を設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(url.QueryEscape(s.ClientID), url.QueryEscape(s.ClientSecret))

	// HTTPリクエストの送受信
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// レスポンスボディを読み込み
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if code := resp.StatusCode; code < 200 || code > 299 {
		return nil, err
	}

	// レスポンスからentitiy.Tokenを作成
	var token *entitiy.Token
	// Content-Type の値を解析
	contentType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	if contentType != "application/json" {
		return nil, err
	}
	token = &entitiy.Token{}

	// HTTPレスポンスのbodyをJSONとして解析しtokenオブジェクトにデコード
	if err = json.Unmarshal(body, token); err != nil {
		return nil, err
	}
	if token.ExpiresIn != 0 {
		token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}

	return token, nil
}
