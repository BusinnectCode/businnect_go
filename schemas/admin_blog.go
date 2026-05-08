package schemas

type AdminBlogArticle map[string]any

type AdminBlogCategory string

const (
	AdminBlogCategoryAnnouncement AdminBlogCategory = "ANNOUNCEMENT"
	AdminBlogCategoryUpdate       AdminBlogCategory = "UPDATE"
	AdminBlogCategoryNews         AdminBlogCategory = "NEWS"
)

type AdminBlogListItem struct {
	Body        string            `json:"body"`
	Category    AdminBlogCategory `json:"category"`
	Cover       string            `json:"cover"`
	IsFeatured  bool              `json:"is_featured"`
	PublishedAt string            `json:"published_at"`
	Slug        string            `json:"slug"`
	Summary     string            `json:"summary"`
	Title       string            `json:"title"`
}

type AdminBlogListResponse struct {
	List       []AdminBlogListItem `json:"list"`
	TotalList  int64               `json:"total_list"`
	TotalPages int64               `json:"total_pages"`
}
