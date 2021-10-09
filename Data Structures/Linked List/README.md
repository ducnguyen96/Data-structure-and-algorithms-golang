## Linked List là gì
Data trong linked-list thì được lưu dưới dạng 1 chuỗi các nodes. 

Thường được so sánh với arrays tuy nhiên thay vì giá trị được lưu trong các box có đánh index giá trị trong linked-list được đặt trong các nodes liên kết với nhau(mỗi node chứa address trỏ đến node tiếp theo).

Với tính chất trên muốn thay đổi giá trị của 1 phần tử trong list, không giống như array có thể trỏ ngay đến phần từ cần update `arr[i]=3` thì trong linked-list, bạn phải từ `head` của list, dùng address của node tiếp theo trên head để truy cập đến `next node`, đến node tiếp theo lại `next node`, ...Vậy để update 1 node trong linked-list trường hợp xấu nhất sẽ mất O(n) time complexity.
## Tại sao
Vậy tại sao lại dùng linked-list.

Cũng với tính chất trên thì việc Add hoặc remove giá trị của node đầu tiên của rất dễ dàng chỉ mất O(1), còn đối với array thì nó phải shift toàn bộ các elements về sau nên mất đến O(n)
## Doubly linked list
Là linked list nhưng mỗi node chứa đồng thời address của node tiếp theo và address của node trước nó. 

Nguồn:
[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/1S0_-VxPLJo/0.jpg)](https://www.youtube.com/watch?v=1S0_-VxPLJo)