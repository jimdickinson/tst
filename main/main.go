package main

import (
	"fmt"

	ternary "github.com/jimdickinson/tst"
)

func main() {
	tst := ternary.NewTST()

	tst.Put("zebra", 1)
	tst.Put("apple", 1)
	tst.Put("banana", 2)
	tst.Put("grape", 3)
	tst.Put("mango", 4)
	tst.Put("melon", 5)
	tst.Put("kiwi", 6)
	tst.Put("grape", 30)

	fmt.Println("apple:", tst.Get("apple"))
	fmt.Println("banana:", tst.Get("banana"))
	fmt.Println("grape:", tst.Get("grape"))
	fmt.Println("mango:", tst.Get("mango"))
	fmt.Println("orange:", tst.Get("orange"))
	fmt.Println("zebraaa:", tst.Get("zebraaa"))

	fmt.Println("All Keys:", tst.AllKeys())

	for iter := ternary.NewTSTIterator(tst); iter.HasNext(); {
		k, v, _ := iter.Next()
		fmt.Println("found key and value while iterating", k, v)
	}

	tst.Delete("grape")

	fmt.Println("All Keys after removing grape:", tst.AllKeys())
	fmt.Println("grape:", tst.Get("grape"))

	fmt.Println("Keys with the prefix ban:", tst.KeysWithPrefix("ban"))
	fmt.Println("Keys with the prefix m:", tst.KeysWithPrefix("m"))

	fmt.Println("between a and c", tst.RangeCollect("a", "c"))
	fmt.Println("between az and mb", tst.RangeCollect("az", "mb"))

	var pattern *ternary.WildcardPattern
	var matchingKeys []string

	pattern = ternary.NewWildcardPattern("mango", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for mango:", matchingKeys)

	pattern = ternary.NewWildcardPattern("mang?", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for mang?:", matchingKeys)

	pattern = ternary.NewWildcardPattern("?ango", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for ?ango:", matchingKeys)

	pattern = ternary.NewWildcardPattern("m?ng?", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for m?ng?:", matchingKeys)

	pattern = ternary.NewWildcardPattern("??ng?", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for ??ng?:", matchingKeys)

	pattern = ternary.NewWildcardPattern("m*", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for m*:", matchingKeys)

	pattern = ternary.NewWildcardPattern("b*", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for b*:", matchingKeys)

	pattern = ternary.NewWildcardPattern("*an*", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for *an*:", matchingKeys)

	pattern = ternary.NewWildcardPattern("*a", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for *a:", matchingKeys)

	pattern = ternary.NewWildcardPattern("*a*", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for *a*:", matchingKeys)

	pattern = ternary.NewWildcardPattern("?????", '?', '*')
	matchingKeys = tst.SearchWildcard(pattern)
	fmt.Println("Matching keys for ?????:", matchingKeys)

}
