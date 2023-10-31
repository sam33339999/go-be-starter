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

		src := []byte("hello world this is src secret 1234567890")
		maxEnLen := hex.EncodedLen(len(src)) // 最大编码长度
		dst1 := make([]byte, maxEnLen)
		n := hex.Encode(dst1, src)

		logger.Infof("encode content: %s\n", dst1[:n])

		secretbite, err := hex.DecodeString(string(dst1[:]))
		// secretbite, err := hex.DecodeString("48656c6c6f20476f7068657221")

		if err != nil {
			panic(err)
		}

		logger.Infof("secretbite is : %+s", secretbite)
		return secretbite
	}
}

func NewJwtTokenGenCommand() *JwtTokenGenCommand {
	return &JwtTokenGenCommand{}
}
