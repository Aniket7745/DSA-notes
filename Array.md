# Static and Dynamic Array

---

- **What is an Array?**
  what is an static array ?

  > A static array is a fixed length container and containing n elements
  > index-able form the range [0, n-1].
  > ==index-able== : it means each slot in the array can be referenced with a number .

- **When and where is a Array used ?**

  1. Storing and accessing sequential data
  2. Temporally storing objects
  3. Used by IO routines a buffers
  4. Lookup table and inverse lookup table
  5. Can be used to return multiple values form a function
  6. Used in dynamic programming to catch ans to sub problems

- Time Complexity: Static Array vs Dynamic Array

| **Operation** | **Static Array** | **Dynamic Array** |
| ------------- | ---------------- | ----------------- |
| **Access**    | O(1)             | O(1)              |
| **Search**    | O(n)             | O(n)              |
| **Insertion** | N/A              | O(n)              |
| **Appending** | N/A              | O(1)              |
| **Deletion**  | N/A              | O(n)              |
![[array.go]]
- **Static array usage example**

Array `A`:

| Index       | 0   | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   |
| ----------- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Value(A[i]) | 44  | 12  | -5  | 17  | 6   | 0   | 3   | 9   | 100 |

Elements in `A` are referenced by there index . There is no other way to access elements in an array index is zero-based , meaning the first element is found in position zero .
A[0] = 44
A[1] = 12
A[7] = 6
A[9] => index out of bound !!

- **Dynamic Array implementation details**
  Array `A` :

| Index     | 0   | 1   |
| --------- | --- | --- |
| Value(ai) | 34  | 4   |

A.add (-7):

| Index     | 0   | 1   | 2   |
| --------- | --- | --- | --- |
| Value(ai) | 34  | 4   | -7  |

A.add (34):

| Index     | 0   | 1   | 2   | 3   |
| --------- | --- | --- | --- | --- |
| Value(ai) | 34  | 4   | -7  | 34  |

A.remove (4):

| Index     | 0   | 1   | 2   |
| --------- | --- | --- | --- |
| Value(ai) | 34  | -7  | 34  |

- **Code implementation**
  [[Dynamic-array (Java)]]

- Problems :
  [[Top 50 leet code]]
  ![[array.go]]