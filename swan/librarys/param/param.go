//go:binary-only-package

package param

const (
	PARAM_TYPE_GLOBAL = "global"
)

func SetParam(paramtype string, paramkey string, paramvalue interface{}) {
}

func GetParam(paramtype string, key string) (value interface{}, exist bool) {
	return
}

func GetParamBool(paramtype string, key string) (value bool, exist bool) {
	return
}

func GetParamInt(paramtype string, key string) (value int, exist bool) {
	return
}

func GetParamFloat(paramtype string, key string) (value float64, exist bool) {
	return
}

func GetParamString(paramtype string, key string) (value string, exist bool) {
	return
}

func GetSystemParam(key string) (value interface{}) {
	return
}

func GetSystemParamInt(key string) (value int, ok bool) {
	return
}

func GetSystemParamString(key string) (value string, ok bool) {
	return
}

func GetSystemParamFloat(key string) (value float64, ok bool) {
	return
}

func GetSystemParamMap(key string) (value map[string]interface{}, ok bool) {
	return
}

func SetSystemParam(key string, value interface{}) (err error) {
	return nil
}

func ReloadSystemParam(key string) error {
	return nil
}