package add

import (
	"fmt"
)

// HandleAddByName 处理只通过 name 操作的场景
func HandleAddByName(names []string, tags []string, remark string, addr string, accessToken string) error {
	if len(names) > 1 {
		return fmt.Errorf("only one name can be specified for this operation")
	}

	name := names[0]

	// 优先级：tags > remark
	if len(tags) > 0 {
		fmt.Printf("Adding tags %v to link with name: %s\n", tags, name)
		return AddTagsByName(name, tags, addr, accessToken)
	}

	if remark != "" {
		fmt.Printf("Adding remark to link with name: %s\n", name)
		return AddRemarkByName(name, remark, addr, accessToken)
	}

	return fmt.Errorf("when only name is provided, you must specify --tag or --remark")
}

// HandleAddByLink 处理只通过 link 操作的场景
func HandleAddByLink(link string, tags []string, remark string, addr string, accessToken string) error {
	// 优先级：tags > remark > 添加新链接
	if len(tags) > 0 {
		fmt.Printf("Adding tags %v to link: %s\n", tags, link)
		return AddTagsByLink(link, tags, addr, accessToken)
	}

	if remark != "" {
		fmt.Printf("Adding remark to link: %s\n", link)
		return AddRemarkByLink(link, remark, addr, accessToken)
	}

	// 只有 link，没有其他参数：添加新链接
	fmt.Printf("Adding new link: %s\n", link)
	return AddLink(link, addr, accessToken)
}

// HandleAddLinkWithNames 处理同时提供 link 和 name 的场景
func HandleAddLinkWithNames(link string, names []string, tags []string, remark string, addr string, accessToken string) error {
	fmt.Printf("Adding new link: %s\n", link)

	// 尝试添加链接，如果已存在则忽略该错误
	if err := AddLink(link, addr, accessToken); err != nil {
		// 检查是否是"链接已存在"错误
		if !isLinkExistsError(err) {
			return fmt.Errorf("failed to add link: %w", err)
		}
		fmt.Println("Link already exists, continuing to add names...")
	}

	fmt.Printf("Adding names %v to link: %s\n", names, link)
	if err := AddNames(link, names, addr, accessToken); err != nil {
		return fmt.Errorf("failed to add names: %w", err)
	}

	if len(tags) > 0 {
		fmt.Printf("Adding tags %v to link: %s\n", tags, link)
		if err := AddTagsByLink(link, tags, addr, accessToken); err != nil {
			return fmt.Errorf("failed to add tags: %w", err)
		}
	}

	if remark != "" {
		fmt.Printf("Adding remark to link: %s\n", link)
		if err := AddRemarkByLink(link, remark, addr, accessToken); err != nil {
			return fmt.Errorf("failed to add remark: %w", err)
		}
	}

	fmt.Println("Successfully completed all add operations.")
	return nil
}

// 辅助函数：检查是否是链接已存在的错误
func isLinkExistsError(err error) bool {
	// 根据你的错误响应格式来判断
	return err.Error() == "link already exists" ||
		err.Error() == "request failed [409]: Link already exists"
}
