## About
Algorithms to giveaway the first card of the deck, then pass the next to the bottom and so it goes until one remains
It returns the giveaway cards in order and the remaining card
3 approaches were chosen and will be compared here. All three approaches implement a queue in the end (it always removes cards from one end of the deck and adds to the other).
The first iteration that will convert to number of cards already starts giving away half of the cards to make the process faster.

## Slices
This approach uses a slice to store the queue data. Basically it will take the first item of the slice to be appended in the `giveawayCards` slice and then re-slice it starting from the element with index 1.
Again it will get the first item and re-slice it starting from the index 1 but this element will be appended at the end of the queue slice.

## Linked List
This approach uses a manual implementation of a singly linked list to store the queue data. Basically it will take the head of the list to be appended in the `giveawayCards`, the head will be replace with the next node. Then the new head will become the new tail of the list.

## Container List
This approach uses a Doubly Linked List from the standard go library (it doesn't have a singly linked list). Basically it will do the same as the previous Linked List approach but with the go library implementation, also changing the the pointers to the previous nodes.

## Final thoughts
For the final thoughts it was calculated the time to run each algorithm, number of loops, the total memory allocated (cumulative, it won't decrease with garbage collection) and the total heap objects allocated.

========= cards_slice =========
time: 2.6255ms
total memory allocated: 8473752
mallocs (cumulative count of heap objects allocated): 99824
total loops: 99999
========= cards_linked_list =========
time: 4.7213ms
total memory allocated: 6507680
mallocs (cumulative count of heap objects allocated): 199793
total loops: 99999
========= cards_container_list =========
time: 7.1058ms
total memory allocated: 10101696
mallocs (cumulative count of heap objects allocated): 249659
total loops: 99999

### Time complexity:
All 3 algorithms takes the same amount of loops, all O(n) time complexity, because it will iterate the same number as the number of cards.
We can think that we end one complete deck loop when we reach the bottom card. But we are jumping 2 cards every iteration. So we loop half of the deck size.
With 1024 cards, we iterate 512 times, the new deck will have 512 cards, we iterate 256, so it goes, until 1. It happens because, considering only integer values, we have a number X and divide it by 2 until it becomes 1 and sum each of these division results + 1, it will always be this number.
Y = X/2 + X/4 + X/8 + X/16 + X/32 + 1. Supose X/32 = 1, it's the last division, so Y = 16X/32 + 8X/32 + 4X/32 + 2X/32 + X/32 + X/32 = 32X/32 = X

But if we look at the absolute values, we can see that the fastest algorithm is the one using slices. That happens because it has less commands in each iteration probably than a singly linked list and a doubly linked list. The singly linked list gets the second place because the doubly linked list takes more iterations probably to change references each iteration.

### Space Complexity:
All 3 algorithms have the same space complexity O(n).

On the other hand, the best when we talk about absolute memory is the linked list approach implemented manually (without using the go standard library). That happens because we use only the space needed for the the elements. The cards_slice will need to alloc more space as it grows t array it refer.
And the approach using `containers/list` go standard library uses more space because it is a doubly linked list. So it stores more data because it has also the reference to the previous node (that's why it allocates almost the double of the space of the sinlgy linked list)

About the slices approach we have a particularity. Slices are a reference type to a underlying array. So, when we reslice the slice (e.g. slice[1:]), we are not allocating space for a new array, we are just changing our first slice pointer to the second array element. So we won't be doubling our allocated memory everytime we reslice. We are also changing the length and the capacity of the slice (reducing it). 
Go has a particular implementation to grow slices, as can be seen here: https://github.com/golang/go/blob/master/src/runtime/slice.go#L157
Basically, if the capacity of the array is less than a threshold (256 for Go 1.20), it will double the capacity when we reach it adding a new element. If it is bigger than that it will do some magic (`newcap += (newcap + 3*threshold) / 4` and more stuff after), basically will be less than the double. This is why the slice approach uses more memory than our linked list approach.
To give an example, for 1000 cards, the first slice will have a capacity of 510 and a length of 499 (because the way it grows). It's not 512 because when it grows from 256 it already do the magic operation, so it will be less.
But every iteration we are removing 2 elements and adding only 1 more to the slice. So, the capacity is reducing by 2 and the lenght by 1, at some point they will meet. This happens at 486, it will then allocate more memory until the capacity of 816. This will happen again at 156 and that's all.
