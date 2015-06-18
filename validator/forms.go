package validator
import "regexp"

func ( v Validator ) Email(check string, value interface{}, t string) bool {

	if t != "string" || t == "" {
		return false
	}

	m, _ := regexp.MatchString("^[a-zA-Z0-9_.+\\-]+@[a-zA-Z0-9\\-_.]+[a-zA-Z0-9\\-]{2,6}$", value.(string))

	return m

}
