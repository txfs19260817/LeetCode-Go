package appliedintuition

import "strings"

// Field 表示结构体中的单个字段定义
type Field struct {
	Name string // 字段名称
	Type string // 字段类型
}

// StructDef 表示一个结构体定义，包含名称和所有字段
type StructDef struct {
	Name   string  // 结构体名称
	Fields []Field // 结构体包含的字段列表
}

// MessageParser 是消息解析器，用于解析消息模式定义和计算字段大小
type MessageParser struct {
	primitiveSizes map[string]int        // 基础数据类型的大小映射（如 int、float等） name to size
	structDefs     map[string]*StructDef // 已解析的结构体定义映射
	messageFields  map[string]string     // Message name to type
	sizeCache      map[string]int        // 类型大小的缓存 name to size
}

// NewMessageParser 创建并返回一个新的 MessageParser 实例
func NewMessageParser(primitives map[string]int) MessageParser {
	return MessageParser{
		primitiveSizes: primitives,
		structDefs:     make(map[string]*StructDef),
		messageFields:  make(map[string]string),
		sizeCache:      make(map[string]int),
	}
}

// Parse 解析给定的模式字符串，提取结构体定义和字段信息
func (mp *MessageParser) Parse(schema string) {
	// 重置所有映射
	mp.structDefs = make(map[string]*StructDef)
	mp.messageFields = make(map[string]string)
	mp.sizeCache = make(map[string]int)

	lines := strings.Split(schema, "\n")
	var currentStruct *StructDef

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 检测结构体定义行（以冒号结尾）
		if strings.HasSuffix(line, ":") {
			structName := strings.TrimSpace(strings.TrimSuffix(line, ":"))
			currentStruct = &StructDef{Name: structName, Fields: []Field{}}
			mp.structDefs[structName] = currentStruct
			continue
		}

		// 解析字段定义行（格式：字段名 字段类型）
		if parts := strings.Fields(line); len(parts) == 2 && currentStruct != nil {
			fieldName, fieldType := parts[0], parts[1]
			currentStruct.Fields = append(currentStruct.Fields, Field{Name: fieldName, Type: fieldType})
			// 如果当前结构体是 Message，将字段添加到 messageFields
			if currentStruct.Name == "Message" {
				mp.messageFields[fieldName] = fieldType
			}
		}
	}
}

// GetType 获取 Message 结构体中指定字段的类型
func (mp *MessageParser) GetType(fieldName string) string {
	if t, ok := mp.messageFields[fieldName]; ok {
		return t
	}
	return ""
}

// GetSize 获取指定查询的大小，可以是基础类型、Message 字段或自定义结构体
func (mp *MessageParser) GetSize(query string) int {
	// 首先检查是否是基础数据类型
	if size, ok := mp.primitiveSizes[query]; ok {
		return size
	}
	// 其次检查是否是 Message 中的字段
	if typeName, ok := mp.messageFields[query]; ok {
		return mp.calculateTypeSize(typeName, map[string]bool{})
	}
	// 最后尝试作为自定义类型计算大小
	return mp.calculateTypeSize(query, map[string]bool{})
}

// calculateTypeSize 递归计算给定类型的总大小，支持检测循环依赖
func (mp *MessageParser) calculateTypeSize(typeName string, visiting map[string]bool) int {
	// 检查是否是基础数据类型
	if size, ok := mp.primitiveSizes[typeName]; ok {
		return size
	}
	// 检查缓存中是否已有该类型的大小
	if size, ok := mp.sizeCache[typeName]; ok {
		return size
	}
	// 检测循环依赖（当前正在访问该类型）
	if visiting[typeName] {
		return -1
	}
	// 获取结构体定义
	def, exists := mp.structDefs[typeName]
	if !exists {
		return -1
	}

	// 标记当前类型正在访问，以检测循环
	visiting[typeName] = true
	totalSize := 0
	// 累计所有字段的大小
	for _, field := range def.Fields {
		fieldSize := mp.calculateTypeSize(field.Type, visiting)
		if fieldSize == -1 {
			return -1
		}
		totalSize += fieldSize
	}
	// 取消访问标记
	visiting[typeName] = false

	// 缓存计算结果以优化性能
	mp.sizeCache[typeName] = totalSize
	return totalSize
}
