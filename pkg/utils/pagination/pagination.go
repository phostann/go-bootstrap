package pagination

func Offset(page, pageSize int) int {
	return (page - 1) * pageSize
}
