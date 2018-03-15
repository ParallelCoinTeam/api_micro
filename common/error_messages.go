package common

func DatabaseError() map[string]string {
	return map[string]string{"code": "2004", "message": "database error"}
}
func CommonError(errStr string) map[string]string {
	return map[string]string{"code": "2056", "message": errStr}
}
