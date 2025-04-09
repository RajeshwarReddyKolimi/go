# Go fundamentals

## Steps followed

- Created a folder fundamentals and created a main.go file and generated go.mod file.
- The main function calls a function for each exercise.

### Exercise 1: Calculator

- Declared 3 variables: n1 (First number), n2 (Second number), op (operator).
- Used fmt.Println() for guiding and fmt.Scanln() for taking user input.
- Then used switch to perform an operation based on the user input.
- Cases are handled for addition, subtraction, multiplication, division and default case (invalid input).
- For division case, handled division by zero by checking if the denominator is 0 and error is logged using fmt.Errorf().

### Exercise 2.1: Check Sublist

- Two arrays A and B of the same type are created.
- Based on their lengths:
  - If A and B are the same length, call checkEqual(A, B) to see if they're exactly equal.
  - If A is shorter, call checkSublist(A, B) to see if A is a sublist of B.
  - If A is longer, call checkSublist(B, A) to see if A is a superlist of B.
- checkEqual(A, B):
  - Compares each element in both arrays one by one.
  - If all elements match, it returns true.
  - If any element is different, it returns false.
- checkSublist(smaller, larger):
  - Looks for the smaller array as a continuous part (sublist) inside the larger array.
  - It checks every possible starting point in the larger array.
  - For each start, it compares elements with the smaller array.
  - If a full match is found, returns true. Otherwise, returns false.

### Exercise 2.2: Word count

- Taken a string as input.
- Used regexp package to create a regular expression.
- Used `\w+('\w+)?` which matches words with letters, digits, underscore and words with single apostrophe followed by some characters.
- Then created a freq map to count freq of each word matched using above regexp after converting the string to lowercase.
- Iterated over the matched words to increment there count in the freq map.

### Exercise 2.3: ETL

- Taken a map of int, array of string as input.
- It contains the alphabets point wise which we need to convert to alphabet wise points.
- Created a new map res of string, int.
- Looped through each point and in each point looped through each alphabet and set res[alphabet] = point

### Exercise 2.4: Grade school

- Created a struct `School` which has a `map[int][]string` with grade as key and array of students in each grade as values.
- Created a function `New()` which returns an empty `School` initialization.
- Created functionalities:
  - `addStudent(name string, grade int)`: Takes student name and grade as input and adds him to the respective grade.
  - `getStudentsOfAGrade(grade) []string` : Takes grade as input and returns the students in that grade.
  - `getAllStudentsSorted() []string`: Returns all the students sorted first by grade and then by alphabetical order.
  - `getAllStudentsWithGrades() map[int][]string`: Returns all the students of the school along with their grades.
  - `getGradeOfAStudent(student string) int`: Takes student name as input and returns his/her grade if present else returns -1.
  - `printGrade(grade int, name string)`: Helper function to log the grade of the student.
- `Exercise24()`: Creates a school instance, adds students, gets them in required manner

### Exercise 3: Interfaces

- Created 2 structs circle and rectangle.
- Created an area method for both of them and called them from main.
- Now instead of calling them separately we can club them using an interface and call with a loop.
- So created an interface shape with area method.
- Now any struct which implements area method is said to satisfy the interface.
- So the variables c, r can be assigned to type shape.
- Created a slice of type shape and called the area method for each shape.

### Exercise 4: File Reading

- Takes the path from user console.
- Used os package's os.Open() function to open the file.
- Used bufio package's bufio.NewScanner() to create a scanner that reads the file.
- Loops through each line of the file and reads it using sc.Scan() and returns true if there's more to read. sc.Text() gives the actual content of the current line.
- Handled errors which can occur during opening or reading the file.
