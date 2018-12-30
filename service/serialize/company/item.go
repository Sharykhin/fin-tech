package company

import (
	"encoding/json"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/service/serialize"
)

func Item(c entity.Company, group string) (json.RawMessage, error) {
	if group == serialize.PUBLIC {
		return json.Marshal(struct {
			entity.Company
		}{
			c,
		})
	}

	panic("err group has not been provided")
}
