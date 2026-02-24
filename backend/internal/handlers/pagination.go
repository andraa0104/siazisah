package handlers

import (
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const defaultPageSize = 5

func parsePagination(c *gin.Context) (int, int, bool) {
	page := 1
	pageSize := defaultPageSize
	isAll := false

	if rawPage := strings.TrimSpace(c.Query("page")); rawPage != "" {
		if parsed, err := strconv.Atoi(rawPage); err == nil && parsed > 0 {
			page = parsed
		}
	}

	rawSize := strings.TrimSpace(c.Query("page_size"))
	if rawSize != "" {
		if strings.EqualFold(rawSize, "all") {
			isAll = true
		} else if parsed, err := strconv.Atoi(rawSize); err == nil {
			if parsed <= 0 {
				isAll = true
			} else {
				pageSize = parsed
			}
		}
	}

	if isAll {
		page = 1
		pageSize = 0
	}

	return page, pageSize, isAll
}

func buildPagination(total, page, pageSize int, isAll bool) map[string]interface{} {
	if total < 0 {
		total = 0
	}

	if isAll {
		pageSize = total
		return map[string]interface{}{
			"page":        1,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": 1,
			"is_all":      true,
		}
	}

	if pageSize <= 0 {
		pageSize = defaultPageSize
	}
	if page <= 0 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	if totalPages == 0 {
		totalPages = 1
	}
	if page > totalPages {
		page = totalPages
	}

	return map[string]interface{}{
		"page":        page,
		"page_size":   pageSize,
		"total":       total,
		"total_pages": totalPages,
		"is_all":      false,
	}
}
