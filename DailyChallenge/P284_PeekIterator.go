/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
package LeetCode

type PeekingIterator struct {
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	future := *(this.iter)
	return future.next()
}
