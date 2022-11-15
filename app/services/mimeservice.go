package services

func DetectMimeType(fileType string) string {
	switch fileType {
	case `js`:
		return `text/javascript`

	case `css`:
		return `text/css`

	case `html`:
		return `text/html`

	case `ico`:
		return `image/vnd.microsoft.icon`

	default:
		return "text/plain"
	}
}
