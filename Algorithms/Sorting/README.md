## Benchmark

<center>

| Algorithms <div style="width:200px"></div> | Time(s) <div style="width:200px"></div> |
| :----------------------------------------- | :-------------------------------------: |
| Bubble Sort                                |                  0.008                  |
| Insertion Sort                             |                  0.004                  |
| Merge Sort                                 |                  0.001                  |
| Selection Sort                             |                  0.007                  |

</center>

- Most of the time we will use quick sort or merge sort
- Radix sort + counting sort only works for number because how it works on machine. they are non-comparision sorting algorithms.

## Sorting Interview: which algorithms to use ?

1. Sort 10 schools around your house by distince.

- **Answer**: Insertion Sort
- **Explain**: Only 10 items so it's fast, simple to code and space complexity of O(1)

2. eBay sorts listings by the current bid amount.

- **Answer**: Radix or Counting Sort
- **Explain**: Value is number from ~1-500$

3. Sort scores on ESPN

- **Answer**: Quick Sort
- **Explain**: We cannot use Radix or Counting because the number is not integer, we will use quick sort to save memory.

4. Massive database (can't fit all into memory) needs to sort though past year's user data

- **Answer**: Merge Sort
- **Explain**: The data can be very large and we want to avoid the worst scenior.

5. Almost sorted Udemy review data needs to update and add 2 new reviews

- **Answer**: Insertion Sort
- **Explain**: The data may be very large but it's almost sorted so insertion sort works very well on this.

6. Temperature records for the past 50 years in Canada

- **Answer**: Radix or Counting Sort if the data is contains only integer else we'll Quick Sort
- **Explain**: Save space

7. Large user name database needs to be sorted. Data is very random.

- **Answer**: Merge Sort || Quick Sort
- **Explain**

8. You want to teach sorting for the first time

- **Answer**: Bubble Sort || Selection Sort
- **Explain**
