package pagination

func Offset(page, pageSize int32) int32 {
	return (page - 1) * pageSize
}
