package delete

import (
	"fmt"
)

// HandleDeleteByName 处理只通过 name 操作的场景
func HandleDeleteByName(names []string, tags []string, addr string, accessToken string) error {
	// 如果提供了 tags，则删除指定名称的标签
	if len(tags) > 0 {
		if len(names) > 1 {
			return fmt.Errorf("only one name can be specified for this operation")
		}

		name := names[0]
		fmt.Printf("Deleting tags %v from link with name: %s\n", tags, name)
		return DeleteTagsByName(name, tags, addr, accessToken)
	}

	// 没有 tags，删除所有指定的名称对应的链接
	if len(names) == 1 {
		fmt.Printf("Deleting link by name: %s\n", names[0])
		return DeleteByName(names[0], addr, accessToken)
	}

	// 如果有多个名称，逐个删除
	fmt.Printf("Deleting %d links by names\n", len(names))
	for _, name := range names {
		fmt.Printf("  - Deleting link by name: %s\n", name)
		if err := DeleteByName(name, addr, accessToken); err != nil {
			return fmt.Errorf("failed to delete link by name '%s': %w", name, err)
		}
	}

	fmt.Println("Successfully deleted all specified links by names.")
	return nil
}

// HandleDeleteByLink 处理只通过 link 操作的场景
func HandleDeleteByLink(links []string, tags []string, addr string, accessToken string) error {
	// 如果提供了 tags，则删除指定链接的标签
	if len(tags) > 0 {
		if len(links) > 1 {
			return fmt.Errorf("only one link can be specified when deleting tags")
		}

		link := links[0]
		fmt.Printf("Deleting tags %v from link: %s\n", tags, link)
		// Note: DeleteTagsByLink 需要 userId 参数，这里传 0，可能需要从 token 中解析
		return DeleteTagsByLink(link, 0, tags, addr, accessToken)
	}

	// 没有 tags，直接删除链接
	fmt.Printf("Deleting %d link(s)\n", len(links))
	return DeleteLinks(links, addr, accessToken)
}

// HandleDeleteNamesFromLink 处理删除链接的名称
func HandleDeleteNamesFromLink(links []string, names []string, addr string, accessToken string) error {
	if len(links) > 1 {
		return fmt.Errorf("only one link can be specified when deleting names")
	}

	link := links[0]
	fmt.Printf("Deleting names %v from link: %s\n", names, link)
	return DeleteNamesByLink(link, names, addr, accessToken)
}
