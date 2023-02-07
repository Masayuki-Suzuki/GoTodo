package utils

import (
  "app/ent"
  "app/types"
)

func GenerateUserStruct(user *ent.User, token string) types.UserDataStruct {
  return types.UserDataStruct{
    Id:           user.ID,
    FirstName:    user.FirstName,
    LastName:     user.LastName,
    FullName:     user.FirstName + " " + user.LastName,
    EmailAddress: user.Email,
    Token:        token,
  }
}
