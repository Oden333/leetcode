# Правила написания тестов

## Структура

Всегда использовать table-driven тесты:

```go
func TestFoo(t *testing.T) {
    t.Parallel()

    // необязательный сетап, общий для всех кейсов
    sharedGraph := buildSomeGraph()

    tests := []struct {
        name  string
        // поля входа и ожидаемого результата
        setup func() *Node   // если нужен объект — через setup func
        want  map[int]int
    }{
        { name: "...", ... },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := Foo(tt.setup())
            assert.Equal(t, tt.want, got, "short failure hint")
        })
    }
}
```

## Правила

1. **Table-driven всегда.** Несколько последовательных `t.Run` без `[]tests` — антипаттерн.
2. **`t.Parallel()`** — в TestXxx и внутри каждого `t.Run`.
3. **Сетап** — предпочитать поле с сериализованным входом (`args [][]int`, `vals []int`) и вызывать builder в цикле (`structs.BuildGraph(tt.args)`). `setup func() T` — только когда нет удобного сериализованного формата и объект нужно собирать вручную.
4. **Сравнение** — думать под каждый тип:
   - примитивы, срезы, map — `assert.Equal` напрямую
   - указатели на структуры (GraphNode, TreeNode) — сериализовать через вспомогательную функцию (например `PrintGraph`), чтобы diff был читаемым
5. **Сообщение в assert** — одна короткая фраза, только если неочевидно что сравнивается.
6. **Имена кейсов** — строка на английском, описывает сценарий, а не входные данные (`"two sources on linear graph"`, не `"starts=[0,4]"`).
7. **Визуализация сложных кейсов** — для деревьев и графов оставлять ASCII-комментарий прямо внутри тест-кейса:
8. **Таймаут 5 секунд** — каждый subtest запускает функцию в горутине и завершается по `time.After(5 * time.Second)`. Это ловит бесконечные циклы и TLE до того, как зависнет весь suite.
9. **Лимит памяти** — перед вызовом делать `runtime.GC()` + `runtime.ReadMemStats`, после — сравнивать `TotalAlloc` delta. Лимит по умолчанию: **50 МБ** (`50 << 20`). Если задача явно требует больше — указать в комментарии почему.

### Шаблон subtest с таймаутом и памятью

```go
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()

        runtime.GC()
        var m1, m2 runtime.MemStats
        runtime.ReadMemStats(&m1)

        done := make(chan struct{}, 1)
        go func() { solve(tt.board); close(done) }()

        select {
        case <-done:
        case <-time.After(5 * time.Second):
            t.Fatal("TLE: exceeded 5s")
        }

        runtime.ReadMemStats(&m2)
        assert.LessOrEqual(t, m2.TotalAlloc-m1.TotalAlloc, uint64(50<<20), "MLE")
        assert.Equal(t, tt.want, tt.board)
    })
}
```

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
