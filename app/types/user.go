package types

type UserDataStruct struct {
  Id           int    `json:"id,omitempty"`
  FirstName    string `json:"first_name,omitempty"`
  LastName     string `json:"last_name,omitempty"`
  FullName     string `json:"full_name,omitempty"`
  EmailAddress string `json:"email,omitempty"`
  Token        string `json:"token,omitempty"`
}
