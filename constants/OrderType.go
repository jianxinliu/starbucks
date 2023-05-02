package constants

import "github.com/pkg/errors"

const (
	CHARGE string = "charge"
	CAFE   string = "cafe"
	VIP    string = "vip"
)

var orderTypeMap = map[string]string{
	"charge": CHARGE,
	"cafe":   CAFE,
	"vip":    VIP,
}

func ParseOrderType(s string) (string, error) {
	t, ok := orderTypeMap[s]
	if !ok {
		return "", errors.New("暂不支持的类型")
	}
	return t, nil
}
