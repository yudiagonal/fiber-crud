package utils

type PaginationMeta struct {
	Page       int   `json:"page"`       // Halaman saat ini
	Limit      int   `json:"limit"`      // Jumlah data per halaman
	Total      int64 `json:"total"`      // Total data yang tersedia
	TotalPages int64 `json:"totalPages"` // Total halaman yang tersedia

}

func NewPaginationMeta(page, limit int, total int64) PaginationMeta {
	totalPages := total / int64(limit)
	if total%int64(limit) != 0 {
		totalPages++
	}
	return PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: total,
	}
}
