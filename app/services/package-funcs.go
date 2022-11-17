package services

func FindInCache(docPath string) ([]byte, bool) {
	downloadData, found := cache.Get(docPath)
	if found {
		return downloadData.([]byte), found
	}
	return nil, false
}
