package gobpcrypt

import (
	"bytes"
	"encoding/json"

	"golang.org/x/net/html"
)

const (
	// ElementTypeRoot is a root type Element struct
	ElementTypeRoot = iota
	// ElementTypeFolder is a type folder for Element struct
	ElementTypeFolder
	// ElementTypeItem is a type item for Element struct
	ElementTypeItem
)

// Element is a main struct for loaded tree
type Element struct {
	Name         string     `json:"name"`
	Parent       *Element   `json:"parent"`
	Childs       []*Element `json:"childs"`
	PasswordHash string     `json:"pwd_hash"`
	Type         int        `json:"type"`
	Login        string     `json:"login"`
	// Name         string     `xml:"name"`
	// Parent       *Element   `xml:"parent"`
	// Childs       []*Element `xml:"childs"`
	// PasswordHash string     `xml:"pwd_hash"`
	// Type         int        `xml:"type"`
	// Login        string     `xml:"login"`
}

// NewElement method return new Element pointer
func NewElement() *Element {
	elem := new(Element)
	elem.Childs = make([]*Element, 0)
	return elem
}

// LoadData function for load XML data to our new tree view
func LoadData(data []byte) (rootElement *Element, err error) {
	rootElement = NewElement()
	rootElement.Name = "ROOT"
	rootElement.Type = ElementTypeRoot

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return
	}

	doNode(doc, rootElement)
	return
}

func doNode(node *html.Node, elem *Element) {
	var current *Element
	if node.Type == html.ElementNode || node.Type == html.DocumentNode {
		current = NewElement()
		switch node.Data {
		case "folder":
			current.Type = ElementTypeFolder
			current.Parent = elem
			elem.Childs = append(elem.Childs, current)
			for _, attr := range node.Attr {
				if attr.Key == "name" {
					current.Name = attr.Val
				}
			}
		case "item":
			current.Type = ElementTypeItem
			current.Parent = elem
			elem.Childs = append(elem.Childs, current)
			for _, attr := range node.Attr {
				if attr.Key == "name" {
					current.Name = attr.Val
				}
			}
		case "login":
			elem.Login = getElementData(node)
			return
		case "pwd":
			elem.PasswordHash = getElementData(node)
			return
		default:
			current = elem
		}
	} else if node.Type == html.TextNode && node.Data == "" {
		return
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		doNode(c, current)
	}
}

func getElementData(node *html.Node) (data string) {
	if node.Type != html.ElementNode {
		panic("node is not element!")
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			data = c.Data
		}
	}
	return
}

// SetMap function sets map from Element
func (elem *Element) SetMap() (result map[string]interface{}) {
	result = make(map[string]interface{})

	result["name"] = elem.Name
	result["pwd_hash"] = elem.PasswordHash
	result["type"] = elem.Type
	result["login"] = elem.Login

	var childMaps []map[string]interface{}
	for _, child := range elem.Childs {
		ch := child.SetMap()
		childMaps = append(childMaps, ch)
	}
	result["childs"] = childMaps

	return
}

//Unmap method for convert map to Element pointer
func Unmap(m map[string]interface{}) *Element {
	res := NewElement()
	res.Name = m["name"].(string)
	res.PasswordHash = m["pwd_hash"].(string)
	res.Type = int(m["type"].(float64))
	res.Login = m["login"].(string)

	if m["childs"] == nil {
		return res
	}
	list := m["childs"].([]interface{})

	for _, ch := range list {
		elem := Unmap(ch.(map[string]interface{}))
		res.Childs = append(res.Childs, elem)
	}

	return res
}

// Marshal method marshal Element to json
func (elem *Element) Marshal() (data []byte, err error) {
	bigs := elem.SetMap()

	return json.MarshalIndent(&bigs, "", "  ")
}

// Unmarshal method unmarshal json to Element tree
func Unmarshal(data []byte) (*Element, error) {
	var rootMap map[string]interface{}

	err := json.Unmarshal(data, &rootMap)
	if err != nil {
		return nil, err
	}

	rootElement := Unmap(rootMap)

	return rootElement, nil
}
