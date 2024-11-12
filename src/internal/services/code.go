package services

import (
	"math/rand"
	"strconv"
)

type CodeService struct {}

func NewCodeService() CodeService {
  return CodeService{}
}

func (s *CodeService) GenerateCode() (code string) {
  codeLen := 6

  for i := 0; i < codeLen; i++ {
    code += strconv.Itoa(rand.Intn(9))
  }

  return
}
