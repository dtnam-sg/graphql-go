package directive

import (
	"context"
	"fmt"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
)

const REGEX_EMAIL = `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`

func ValidateEmail(ctx context.Context, obj any, next graphql.Resolver, message *string) (res any, err error) {
	resp, err := next(ctx)
	if err != nil {
		return resp, err
	}

	email, ok := resp.(string)
	if !ok || email == "" {
		return nil, fmt.Errorf(*message)
	}

	re := regexp.MustCompile(REGEX_EMAIL)
	if !re.MatchString(email) {
		return nil, fmt.Errorf(*message)
	}
	return resp, nil
}

func Length(ctx context.Context, obj any, next graphql.Resolver, min int32, max *int32, message *string) (res any, err error) {
	resp, err := next(ctx)
	if err != nil {
		return resp, err
	}

	str, ok := resp.(string)
	if !ok {
		return resp, nil
	}

	l := len(str)
	if l < int(min) || l > int(*max) {
		return nil, fmt.Errorf(*message)
	}
	return resp, nil
}
