package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	reader := bufio.NewReader(r)

	result := make(DomainStat)
	var user User

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return result, err
		}

		if err = json.Unmarshal(line, &user); err != nil {
			return result, err
		}

		if strings.Contains(user.Email, domain) {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
		}
	}

	return result, nil
}
