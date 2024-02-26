## Cache

### LFU (Least Frequently Used)

The LFU (Least Frequently Used) cache is a caching algorithm that keeps track of the frequency of access for each item in the cache. It aims to evict the least frequently used items when the cache reaches its capacity. The LFU cache maintains a count of accesses for each item, and when the cache is full, it removes the item with the lowest access frequency.

Here's a conceptual overview of how the LFU cache works:

1. Initialization:

- The LFU cache is initialized with a specified capacity (cap).
- It maintains a count of the number of items currently in the cache (len), the overall capacity (cap), the minimum frequency (minFreq), a map (itemMap) to quickly locate items by their keys, and a map (freqMap) to group items by their access frequency.

2. Item Representation:

- Each item in the cache is represented by the item struct, which includes a key (key), a value (value), and a frequency count (freq).

3. Insertion:

- When a new item is added to the cache using the Put method:
  - If the item already exists, update its value, increase its frequency, and adjust its position in the frequency list.
  - If the item is new:
    - If the cache is at full capacity, eliminate the least frequently used item from the cache.
    - Insert the new item into the cache with a frequency of 1.
    - Update the minimum frequency (minFreq) if necessary.

4. Frequency Management:

- The increaseFreq method is responsible for managing the frequency of an item when it is accessed:

  - Remove the item from its current position in the frequency list.
  - If the removed item had the minimum frequency and the frequency list is empty, update the minimum frequency.
  - Increment the frequency of the item.
  - Insert the item into the cache at its new frequency.

5. Access (Get) Operation:

- When an item is requested from the cache using the Get method:
  - If the item exists, increase its frequency and return its value.
  - If the item does not exist, return nil.

6. Elimination of Least Frequently Used Item:

- When the cache is at full capacity and a new item needs to be added, the eliminate method is called:
  - Identify the list of items with the minimum frequency (minFreq).
  - Remove the least recently used item from this list.
  - Delete the item from the cache.

### LRU (Least Recently Used)

The LRU (Least Recently Used) cache is a caching algorithm that removes the least recently used items when the cache reaches its capacity. It maintains a linked list to track the order of access for each item, promoting recently used items to the front of the list. When the cache is full, it evicts the item at the back of the list, which represents the least recently used item.
