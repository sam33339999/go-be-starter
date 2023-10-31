package commands

import (
	"encoding/hex"

	"github.com/sam33339999/go-be-starter/lib"
	"github.com/spf13/cobra"
)

type JwtTokenGenCommand struct{}

func (s *JwtTokenGenCommand) Short() string {
	return "generate jwt token"
}

func (s *JwtTokenGenCommand) SetUp(cmd *cobra.Command) {}

func (s *JwtTokenGenCommand) Run() lib.CommandRunner {
	return func(
		env lib.Env,
		logger lib.Logger,
	) []byte {
		logger.Infof("jwt secret is : %+s", env.JWTSecret)

		src := []byte("hello world this is jwt secret")
		maxEnLen := hex.EncodedLen(len(src)) // 最大編碼長度
		dst1 := make([]byte, maxEnLen)
		n := hex.Encode(dst1, src)

		logger.Infof("encode content: %s\n", dst1[:n])

		secretBite, err := hex.DecodeString(string(dst1[:]))
		if err != nil {
			panic(err)
		}

		logger.Infof("secret bite is : %+s", secretBite)
		return secretBite
	}
}

func NewJwtTokenGenCommand() *JwtTokenGenCommand {
	return &JwtTokenGenCommand{}
}
