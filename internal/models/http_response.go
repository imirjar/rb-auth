package models

import "time"

/*HTTP/1.1 200 OK
Content-Type: application/json

{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  "token_type": "Bearer",
  "expires_in": 3600
}
*/

type Response struct {
	Access_token  string    `json:"access_token"`
	Refresh_token string    `json:"refresh_token"`
	Token_type    string    `json:"token_type"`
	Expires_in    time.Time `json:"expires_in"`
}
