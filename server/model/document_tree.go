package model

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
)

type DocumentTree struct {
	DocumentId   int               `json:"id"`
	DocumentName string            `json:"text"`
	ParentId     interface{}       `json:"parent"`
	Identify     string            `json:"identify"`
	BookIdentify string            `json:"-"`
	Version      int64             `json:"version"`
	State        *DocumentSelected `json:"state,omitempty"`
}

type DocumentSelected struct {
	Selected bool `json:"selected"`
	Opened   bool `json:"opened"`
}

//获取书籍的文档树状结构
func (m *Document) FindDocumentTree(bookId int, selectedId int, isEdit ...bool) ([]*DocumentTree, error) {
	docs := NewDocument().getDocsByBookId(bookId, "id", "version", "name", "parent_id", "identify")

	book, _ := NewBook().GetById(bookId)

	trees := make([]*DocumentTree, len(docs))

	for index, item := range docs {
		tree := &DocumentTree{}
		if selectedId > 0 {
			if selectedId == item.Id {
				tree.State = &DocumentSelected{Selected: true, Opened: true}
			}
		} else {
			if index == 0 {
				tree.State = &DocumentSelected{Selected: true, Opened: true}
			}
		}

		tree.DocumentId = item.Id
		tree.Identify = item.Identify
		tree.Version = item.Version
		tree.BookIdentify = book.Identify
		if item.ParentId > 0 {
			tree.ParentId = item.ParentId
		} else {
			tree.ParentId = "#"
		}
		idf := item.Identify
		if idf == "" {
			idf = strconv.Itoa(item.Id)
		}
		if len(isEdit) > 0 && isEdit[0] {
			tree.DocumentName = item.Name + "<small class='text-danger'>(" + idf + ")</small>"
		} else {
			tree.DocumentName = item.Name
		}

		trees[index] = tree
	}

	return trees, nil
}

func (m *Document) CreateDocumentTreeForHtml(bookId, selectedId int) (string, error) {
	trees, err := m.FindDocumentTree(bookId, selectedId)
	if err != nil {
		return "", err
	}
	parentId := getSelectedNode(trees, selectedId)

	buf := bytes.NewBufferString("")

	getDocumentTree(trees, 0, selectedId, parentId, buf)

	return buf.String(), nil
}

// 使用递归的方式获取指定ID的最顶层ID
func getSelectedNode(array []*DocumentTree, selectedId int, level ...int) int {
	lv := 0
	if len(level) > 0 {
		lv = level[0]
	}
	if lv > 20 {
		return 0
	}
	for _, item := range array {
		if _, ok := item.ParentId.(string); ok && item.DocumentId == selectedId {
			return item.DocumentId
		} else if pid, ok := item.ParentId.(int); ok && item.DocumentId == selectedId {
			lv++
			return getSelectedNode(array, pid, lv)
		}
	}
	return 0
}

func getDocumentTree(array []*DocumentTree, parentId int, selectedId int, selectedParentId int, buf *bytes.Buffer) {
	buf.WriteString("<ul>")

	for _, item := range array {
		pid := 0

		if p, ok := item.ParentId.(int); ok {
			pid = p
		}
		if pid == parentId {

			selected := ""
			if item.DocumentId == selectedId {
				selected = ` class="jstree-clicked"`
			}
			selectedLi := ""
			if item.DocumentId == selectedParentId {
				selectedLi = ` class="jstree-open"`
			}
			buf.WriteString("<li id=\"")
			buf.WriteString(strconv.Itoa(item.DocumentId))
			buf.WriteString("\"")
			buf.WriteString(selectedLi)
			buf.WriteString("><a href=\"")
			if item.Identify != "" {
				uri := item.Identify
				buf.WriteString(uri)
			} else {
				uri := fmt.Sprint(item.DocumentId)
				buf.WriteString(uri)
			}
			buf.WriteString("\" title=\"")
			buf.WriteString(template.HTMLEscapeString(item.DocumentName) + "\"")
			buf.WriteString(selected + ">")
			buf.WriteString(template.HTMLEscapeString(item.DocumentName) + "</a>")

			for _, sub := range array {
				if p, ok := sub.ParentId.(int); ok && p == item.DocumentId {
					getDocumentTree(array, p, selectedId, selectedParentId, buf)
					break
				}
			}
			buf.WriteString("</li>")
		}
	}
	buf.WriteString("</ul>")
}
