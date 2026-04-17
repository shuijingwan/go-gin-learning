package main

import "fmt"

// List 表示一个单向链表，可以存储任意类型的值
type List[T any] struct {
	next *List[T]
	val  T
}

// New 创建一个包含给定值的新链表节点
// 参数:
//   - val: 存储在链表节点中的初始值
//
// 返回值:
//   - 指向新创建的包含该值的链表的指针
func New[T any](val T) *List[T] {
	return &List[T]{val: val}
}

// Append 在链表末尾添加一个新元素
// 参数:
//   - val: 要追加到链表的值
func (l *List[T]) Append(val T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: val}
}

// Prepend 在链表开头添加一个新元素
// 参数:
//   - val: 要前置到链表的值
//
// 返回值:
//   - 指向新链表头节点的指针，包含前置的值
func (l *List[T]) Prepend(val T) *List[T] {
	return &List[T]{
		val:  val,
		next: l,
	}
}

// Length 返回链表中的元素数量
// 返回值:
//   - 链表中节点的总数
func (l *List[T]) Length() int {
	if l == nil {
		return 0
	}
	count := 0
	current := l
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// ToSlice 将链表转换为切片
// 返回值:
//   - 包含链表中所有元素的切片
func (l *List[T]) ToSlice() []T {
	if l == nil {
		return []T{}
	}
	result := make([]T, 0, l.Length())
	current := l
	for current != nil {
		result = append(result, current.val)
		current = current.next
	}
	return result
}

// Iterate 对链表中的每个元素应用给定的函数
// 参数:
//   - fn: 要应用于每个元素的函数
func (l *List[T]) Iterate(fn func(T)) {
	if l == nil {
		return
	}
	current := l
	for current != nil {
		fn(current.val)
		current = current.next
	}
}

// Map 通过对每个元素应用函数来创建一个新的链表
// 参数:
//   - l: 原始链表
//   - fn: 转换函数，将类型 T 转换为类型 U
//
// 返回值:
//   - 包含转换后元素的新链表
func Map[T, U any](l *List[T], fn func(T) U) *List[U] {
	if l == nil {
		return nil
	}

	newHead := New(fn(l.val))
	currentNew := newHead
	currentOld := l.next

	for currentOld != nil {
		currentNew.next = New(fn(currentOld.val))
		currentNew = currentNew.next
		currentOld = currentOld.next
	}

	return newHead
}

// Filter 创建一个只包含满足谓词条件的元素的新链表
// 参数:
//   - predicate: 判断函数，返回 true 表示保留该元素
//
// 返回值:
//   - 包含满足条件元素的新链表
func (l *List[T]) Filter(predicate func(T) bool) *List[T] {
	if l == nil {
		return nil
	}

	var newHead *List[T]
	var currentNew *List[T]

	current := l
	for current != nil {
		if predicate(current.val) {
			if newHead == nil {
				newHead = New(current.val)
				currentNew = newHead
			} else {
				currentNew.next = New(current.val)
				currentNew = currentNew.next
			}
		}
		current = current.next
	}

	return newHead
}

// Contains 检查链表是否包含给定的值
// 参数:
//   - l: 要搜索的链表
//   - val: 要查找的值
//
// 返回值:
//   - 如果找到值则返回 true，否则返回 false
func Contains[T comparable](l *List[T], val T) bool {
	_, found := Find(l, val)
	return found
}

// Find 返回值的第一次出现位置及其索引，如果未找到则返回 -1
// 参数:
//   - l: 要搜索的链表
//   - val: 要查找的值
//
// 返回值:
//   - int: 元素的索引位置，未找到时返回 -1
//   - bool: 如果找到值则返回 true，否则返回 false
func Find[T comparable](l *List[T], val T) (int, bool) {
	if l == nil {
		return -1, false
	}
	current := l
	index := 0
	for current != nil {
		if current.val == val {
			return index, true
		}
		current = current.next
		index++
	}
	return -1, false
}

// Delete 从链表中删除第一次出现的指定值
// 参数:
//   - l: 要操作的链表
//   - val: 要删除的值
//
// 返回值:
//   - *List[T]: 删除后的新链表头节点
//   - bool: 如果找到并删除了值则返回 true，否则返回 false
func Delete[T comparable](l *List[T], val T) (*List[T], bool) {
	if l == nil {
		return nil, false
	}

	if l.val == val {
		return l.next, true
	}

	current := l
	for current.next != nil {
		if current.next.val == val {
			current.next = current.next.next
			return l, true
		}
		current = current.next
	}
	return l, false
}

func main() {
	// 创建整数链表
	intList := New(10)
	intList.Append(20)
	intList.Append(30)
	intList.Append(40)

	fmt.Println("Integer List:", intList.ToSlice())
	fmt.Println("Length:", intList.Length())

	// 测试 Contains 方法
	fmt.Println("Contains 20:", Contains(intList, 20))
	fmt.Println("Contains 50:", Contains(intList, 50))

	// 测试 Find 方法
	if idx, found := Find(intList, 30); found {
		fmt.Printf("Found 30 at index %d\n", idx)
	}

	// 测试 Prepend 方法
	intList = intList.Prepend(5)
	fmt.Println("After prepend 5:", intList.ToSlice())

	// 测试 Delete 方法
	intList, deleted := Delete(intList, 20)
	fmt.Printf("Deleted 20: %v, List: %v\n", deleted, intList.ToSlice())

	// 测试 Iterate 方法
	fmt.Print("Iterate: ")
	intList.Iterate(func(val int) {
		fmt.Printf("%d ", val)
	})
	fmt.Println()

	// 测试 Map 方法
	doubled := Map(intList, func(x int) int {
		return x * 2
	})
	fmt.Println("Doubled:", doubled.ToSlice())

	// 测试 Filter 方法
	evens := intList.Filter(func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Even numbers:", evens.ToSlice())

	// 创建字符串链表
	strList := New("hello")
	strList.Append("world")
	strList.Append("go")

	fmt.Println("\nString list:", strList.ToSlice())
	fmt.Println("Length:", strList.Length())
	fmt.Println("Contains 'world':", Contains(strList, "world"))
}
