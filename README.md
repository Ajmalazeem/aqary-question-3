# aqary-question-3
Assesment question number 3 given by aqary international group

// code ncontains only the solution for the given question. adding values to the table is done manually
// environment variables are not added in gitignore for refernce purpose

question number 3:

Table: Seat
+--------------------------+------------------+
| Column Name | Type |
+--------------------------+------------------+
| id | int |
| student | varchar |
+--------------------------+------------------+
id is the primary key (unique value) column for this table.
Each row of this table indicates the name and the ID of a student.
id is a continuous increment.
Write a solution to swap the seat id of every two consecutive students. If the number of students is odd, the id of the last student is not swapped.
Return the result table ordered by id in ascending order.
The result format is in the following example.
Example 1:
Input:
Seat table:
+--------+------------------+
| id | student |
+--------+------------------+
| 1 | Abbot |
| 2 | Doris |
| 3 | Emerson |
| 4 | Green |
| 5 | Jeames |
+--------+------------------+
Output:
+--------+------------------+
| id | student |
+--------+------------------+
| 1 | Doris |
| 2 | Abbot |
| 3 | Green |
| 4 | Emerson |
| 5 | Jeames |
+--------+------------------+


output:

for odd number of students

![alt text](image.png)

for even number of students

![alt text](image-1.png)
