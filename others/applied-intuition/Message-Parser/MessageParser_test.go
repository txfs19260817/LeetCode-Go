package appliedintuition

import "testing"

func BuildParser(schema string, primitives map[string]int) (func(string) int, func(string) string) {
	parser := NewMessageParser(primitives)
	parser.Parse(schema)
	return parser.GetSize, parser.GetType
}

func TestMessageParserCases(t *testing.T) {
	prim := map[string]int{"int": 4, "float": 8, "char": 1}

	t.Run("case1_basic", func(t *testing.T) {
		schema1 := `
Message:
a int
b float
`
		getSize, getType := BuildParser(schema1, prim)

		if got := getType("a"); got != "int" {
			t.Fatalf("get_type(a) = %s, want int", got)
		}
		if got := getSize("a"); got != 4 {
			t.Fatalf("get_size(a) = %d, want 4", got)
		}
		if got := getSize("Message"); got != 12 {
			t.Fatalf("get_size(Message) = %d, want 12", got)
		}
		if got := getSize("float"); got != 8 {
			t.Fatalf("get_size(float) = %d, want 8", got)
		}
	})

	t.Run("case2_nested", func(t *testing.T) {
		schema2 := `
Header:
x int
y char

Message:
h Header
z float
`
		getSize, getType := BuildParser(schema2, prim)

		if got := getType("h"); got != "Header" {
			t.Fatalf("get_type(h) = %s, want Header", got)
		}
		if got := getSize("Header"); got != 5 {
			t.Fatalf("get_size(Header) = %d, want 5", got)
		}
		if got := getSize("Message"); got != 13 {
			t.Fatalf("get_size(Message) = %d, want 13", got)
		}
	})

	t.Run("case3_unknown_type", func(t *testing.T) {
		schema3 := `
Message:
a int
b Unknown
`
		getSize, _ := BuildParser(schema3, prim)

		if got := getSize("Message"); got != -1 {
			t.Fatalf("get_size(Message) = %d, want -1", got)
		}
	})

	t.Run("case4_cycle", func(t *testing.T) {
		schema4 := `
A:
b B

B:
a A
`
		getSize, _ := BuildParser(schema4, prim)

		if got := getSize("A"); got != -1 {
			t.Fatalf("get_size(A) = %d, want -1", got)
		}
	})

	t.Run("case5_unknown_field", func(t *testing.T) {
		schema5 := `
Message:
a int
`
		getSize, getType := BuildParser(schema5, prim)

		if got := getType("nope"); got != "" {
			t.Fatalf("get_type(nope) = %s, want empty", got)
		}
		if got := getSize("nope"); got != -1 {
			t.Fatalf("get_size(nope) = %d, want -1", got)
		}
	})

	t.Run("case6_super_complex", func(t *testing.T) {
		// 超复杂案例：深层嵌套、多重字段、混合基础类型和自定义类型
		schema6 := `
Metadata:
version int
timestamp int

Payload:
data float
count int

Header:
meta Metadata
priority char

Body:
payload Payload
signature char
checksum int

Footer:
timestamp int
validator char

Message:
header Header
body Body
footer Footer
extra char
reserved int
`
		getSize, getType := BuildParser(schema6, prim)

		// 验证基础字段类型
		if got := getType("header"); got != "Header" {
			t.Fatalf("get_type(header) = %s, want Header", got)
		}
		if got := getType("body"); got != "Body" {
			t.Fatalf("get_type(body) = %s, want Body", got)
		}
		if got := getType("footer"); got != "Footer" {
			t.Fatalf("get_type(footer) = %s, want Footer", got)
		}

		// 验证嵌套结构体大小：Metadata = int(4) + int(4) = 8
		if got := getSize("Metadata"); got != 8 {
			t.Fatalf("get_size(Metadata) = %d, want 8", got)
		}

		// 验证嵌套结构体大小：Payload = float(8) + int(4) = 12
		if got := getSize("Payload"); got != 12 {
			t.Fatalf("get_size(Payload) = %d, want 12", got)
		}

		// 验证嵌套结构体大小：Header = Metadata(8) + char(1) = 9
		if got := getSize("Header"); got != 9 {
			t.Fatalf("get_size(Header) = %d, want 9", got)
		}

		// 验证嵌套结构体大小：Body = Payload(12) + char(1) + int(4) = 17
		if got := getSize("Body"); got != 17 {
			t.Fatalf("get_size(Body) = %d, want 17", got)
		}

		// 验证嵌套结构体大小：Footer = int(4) + char(1) = 5
		if got := getSize("Footer"); got != 5 {
			t.Fatalf("get_size(Footer) = %d, want 5", got)
		}

		// 验证最终消息大小：Message = Header(9) + Body(17) + Footer(5) + char(1) + int(4) = 36
		if got := getSize("Message"); got != 36 {
			t.Fatalf("get_size(Message) = %d, want 36", got)
		}

		// 验证查询基础类型
		if got := getSize("int"); got != 4 {
			t.Fatalf("get_size(int) = %d, want 4", got)
		}
		if got := getSize("float"); got != 8 {
			t.Fatalf("get_size(float) = %d, want 8", got)
		}
		if got := getSize("char"); got != 1 {
			t.Fatalf("get_size(char) = %d, want 1", got)
		}
	})
}
