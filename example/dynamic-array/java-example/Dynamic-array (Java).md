This project provides a basic **generic Dynamic Array implementation in Java**, mimicking the behavior of dynamically re-sizable arrays (like `ArrayList`) with custom logic for resizing, adding, removing, and accessing elements.

## Features

- Generic type support (`T`)
- Auto-resizing (doubling capacity when needed)
- `add`, `get`, `set`, `removeAt`, `remove`, and `contains` operations
- Manual memory clearing (`clear()`)
- Iterate support via `Iterator<T>`
- Custom `toString()` output for visualization

## Key Components

| Method                   | Description                                                   |
| ------------------------ | ------------------------------------------------------------- |
| `add(T elem)`            | Adds an element to the end of the array, resizes if necessary |
| `get(int index)`         | Retrieves the element at the given index                      |
| `set(int index, T elem)` | Sets the value at a specified index                           |
| `removeAt(int index)`    | Removes the element at a specific index                       |
| `remove(Object obj)`     | Removes the first occurrence of the object                    |
| `contains(Object obj)`   | Checks if the object is present in the array                  |
| `clear()`                | Clears all elements from the array                            |
| `size()`                 | Returns the current number of elements                        |
| `isEmpty()`              | Checks if the array is empty                                  |

## Example

```java
DynamicArray<Integer> arr = new DynamicArray<>();
arr.add(10);
arr.add(20);
arr.add(30);

System.out.println(arr); // Output: [10, 20, 30]

arr.removeAt(1);
System.out.println(arr); // Output: [10, 30]

System.out.println(arr.contains(30)); // Output: true
```


![[dynamicArray.java]]