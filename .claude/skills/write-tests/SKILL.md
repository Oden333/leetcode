---
name: write-tests
description: Generate comprehensive unit tests for a Go function or problem file. Delegates to the test-writer subagent.
argument-hint: <path-to-file>
agent: test-writer
allowed-tools: Read Write Bash(go test *)
user-invocable: true
---

Делегируй задачу агенту `test-writer`. Файл для покрытия тестами: $ARGUMENTS

Контекст для агента:
- Модуль: `dsa`
- Тестирование: testify assert
- Вспомогательные типы: `dsa/structs` (TreeNode, ListNode, Graph и др.)
- Паттерн именования тест-файла: тот же путь с суффиксом `_test.go`

**Перед генерацией тестов прочитай файл со стилевыми правилами:**
`.claude/rules/testing.md` — там полный шаблон, правила таймаута/памяти и примеры.

## Обязательный стиль тестов (краткая выжимка, полное описание в testing.md)

**Table-driven всегда.** Несколько последовательных `t.Run` без `[]tests` — антипаттерн.

**Правила (детали в testing.md):**
1. Table-driven тесты, `t.Parallel()` везде
2. Таймаут 5 секунд через горутину + `time.After`
3. Лимит памяти 50 МБ через `runtime.MemStats`
4. ASCII-комментарии для деревьев и графов в кейсах:

```go
{
    //     1
    //    / \
    //   2   3
    //    \
    //     5
    name: "right child only on node 2",
    root: structs.BuildTree([]int{1, 2, 3, structs.Null, 5}),
    want: ...,
},
```
