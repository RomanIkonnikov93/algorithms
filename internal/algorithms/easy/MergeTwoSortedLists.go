package easy

/*
Вам даны заголовки двух отсортированных связанных списков list1 и list2.
Объедините два списка в один отсортированный список.
Список должен быть составлен путем соединения узлов первых двух списков.
Возвращает заголовок объединенного связанного списка.
*/

/*
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: list1 = [], list2 = []
Output: []

Input: list1 = [], list2 = [0]
Output: [0]
*/

func MergeTwoSortedListsRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoSortedListsRecursive(l1.Next, l2)
		return l1
	}
	l2.Next = MergeTwoSortedListsRecursive(l1, l2.Next)
	return l2
}
