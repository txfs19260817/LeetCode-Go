# Moshi Bug Review

以下是我在本项目中确认的 3 个问题（不改源码，只给出定位和修复建议）。

## 1) `ParserStack.endObject()` 未正确推进父数组下标（JSONPath 错误）
- 位置: `moshi/src/main/java/com/squareup/moshi/ParserStack.java` 的 `endObject()`
- 现象: 当对象是数组元素时，读完 `}` 后路径仍停在旧下标。
  - 例: 读取 `[{}]`，`endObject()` 后实际是 `$[0]`，预期应为 `$[1]`。
- 原因: 关闭对象时只 `stackSize--`，没有把“父数组当前元素已消费”反映到 `pathIndices`。
- 建议修复:
  - 在 `endObject()` 出栈后推进父层数组索引（若父层是数组）。

## 2) `ParserStack.endArray()` 未正确推进父数组下标（JSONPath 错误）
- 位置: `moshi/src/main/java/com/squareup/moshi/ParserStack.java` 的 `endArray()`
- 现象: 当数组是数组元素时，读完 `]` 后父数组下标不前进。
  - 例: 读取 `[[]]`，内层 `endArray()` 后实际是 `$[0]`，预期应为 `$[1]`。
- 原因: 关闭数组时只 `stackSize--`，没有把“父数组一个元素已结束”计入下标。
- 建议修复:
  - 在 `endArray()` 出栈后推进父层数组索引（若父层是数组）。

## 3) `JsonReaderPathTest.emptyObj()` 是无效测试用法，会直接失败
- 位置: `moshi/src/test/java/com/squareup/moshi/JsonReaderPathTest.java:25` (`emptyObj`)
- 现象: 对 `{}` 直接调用 `reader.endObject()`，没有先 `beginObject()`，会抛 `IllegalStateException`。
- 原因: 调用序列不符合 `JsonReader` 协议；且包含 `System.out.println` 调试输出，不应出现在单测。
- 建议修复:
  - 删除该测试，或改成标准流程: `beginObject()` -> `endObject()` -> `assertEquals("$", reader.getPath())`。
  - 移除 `System.out.println` 调试语句。

