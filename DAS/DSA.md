# What is a Data Structures ?

A data Structures is a way organizing data so that it can be used efficiently.

# why Data structures ?

- They essential ingredients in creating fast and powerful algorithms.
- They help to mange and organized data.
- They make code cleaner and easier to understand.

# Classification of Data structure:

| **Category**                  | **Subcategory**        | **Data Structures**                   |
| ----------------------------- | ---------------------- | ------------------------------------- |
| **Linear Data Structure**     | Static Data Structure  | [[Array]]                             |
|                               | Dynamic Data Structure | [[Queue]], [[Stack]], [[Linked List]] |
| **Non-linear Data Structure** | [[Tree]]               | [[General Tree Types]]                |
|                               | [[Graph]]              | [[Graph Structures]]                  |

# Abstract Data Type and Data Structures:

**Abstract Data type** :

- An Abstract datatype in an abstraction of a data structure which a data structure must be adhere to.
- The interface dose not give any specific details about how something should be implement or in what programming language .

**Example :**

| ADT     | DS                                                          |
| ------- | ----------------------------------------------------------- |
| List    | Daynamic Array , Link list                                  |
| Queue   | Link List based Queue ,Array based Queue ,Stack based Queue |
| Map     | Tree Map , Hash Map /Hash Table                             |
| Vehicle | Golf cart ,bicycle,car ,bike                                |

# Computational Complexity Analysis :

**Computational Analysis :**

- How much time dose this algorithm needs to finish?
- How much space does this algorithm meed for its Computation ?

# Big-O Notation :

_Big-O Notation gives an upper bound pf the Complexity in the worst case , helping to qualify performance as the input size becomes arbitrarily large_

n - The size of the input Complexities order in form in smaller to largest

| Time               | Big-O           |
| ------------------ | --------------- |
| Constant Time      | ==O(1)==        |
| Logarithemic Time  | ==O(log(n))==   |
| Linear Time        | ==O(n)==        |
| Linearithemic Time | ==O(nlog(n))==  |
| Quadric Time       | ==O(n^2)==      |
| cubic Time         | ==O(n^3)==      |
| Exponential Time   | ==O(b^n), b>1== |
| Factorial Time     | ==O(n!)==       |

**Big-O Property** :

$$
O(n + c) = O(n)
o(cn) = O(n), c > 0
$$

lets f be a function that describe the rune time of a particular algorithm for an input of size n:

$$
f(n) = 7log(n)^3 + 15n^2 + 2n^3 + 8
O(f(n)) = )(n^3)
$$

**Big-O Example:**

- The following run in constant time: O(1)

```
a := 1
b := 2
c := a + 5*b

i := 0
  while i<11 Do
    i = i + 1
```

- The following run in linear time: O(n)

```
i := 0
while i < n Do
i = i + #

        f(n) = n
        O(f(n)) = O(n)

i := 0
while i < n i+3

f(n) = n/3
O(f(n)) = O(n)
```

2. Both of the following run in quadratic time. The first may be obvious and since n works done n time is
   n\*n = 0(n^2), but what about the second one?

```
For (i := 0 ;i < n i +1)
    For (j := 0 ; j< n ; j = j+1)

f(n) = n\*n = n^2 , O(f(n)) = O(n^2)

For (i + 0; i < n; i = i + 1)
    For (j := i ; j < n ; j = j + 1)
```

3. Suppose we have a sorted array and we want to find the index of a particular value in the array , if it exists . What is the time Complexity of the following algorithm?

```
low := 0     |     Ans : O(log_2(n)) = O(log(n))
high := n-1
    while low <=high do
          mid := (low + high )/ 2
    if array [mid] == value : return mid
    else if array[mid] < value : lo mid + 1
    else if array[mid] > value : hi = mid -1

return -1 // value not found

```

_Finding all subsets of a set - ==O(2^n)==_
_Finding all permutations of string - ==O(n!)==_
_Sorting using merge sort - ==O(nlog(n))==_
_Iterating over all the cell in a matrix of size n by m - ==O(nm)_
