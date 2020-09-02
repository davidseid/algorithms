package simplifyPath

import (
	"strings"
	"testing"
)

/*
71. Simplify Path
https://leetcode.com/problems/simplify-path/

Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level.

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.



Example 1:

Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
Example 2:

Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
Example 3:

Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
Example 4:

Input: "/a/./b/../../c/"
Output: "/c"
Example 5:

Input: "/a/../../b/../c//.//"
Output: "/c"
Example 6:

Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
*/

/*
Rationale:
- Remove ./
- On ../, remove and remove prior directory (unless it is the root)
- On consecutive /, collapse to 1
- Remote trailing slashes
*/

func TestSimplifyPath(t *testing.T) {
	input := "/home/"

	expected := "/home"

	actual := simplifyPath(input)

	if actual != expected {
		t.Errorf("Got %s, want %s", actual, expected)
	}
}

func TestSimplifyPath2(t *testing.T) {
	input := "/../"

	expected := "/"

	actual := simplifyPath(input)

	if actual != expected {
		t.Errorf("Got %s, want %s", actual, expected)
	}
}

func TestSimplifyPath3(t *testing.T) {
	input := "/home//foo/"

	expected := "/home/foo"

	actual := simplifyPath(input)

	if actual != expected {
		t.Errorf("Got %s, want %s", actual, expected)
	}
}

func TestSimplifyPath4(t *testing.T) {
	input := "/a/./b/../../c/"

	expected := "/c"

	actual := simplifyPath(input)

	if actual != expected {
		t.Errorf("Got %s, want %s", actual, expected)
	}
}

func TestSimplifyPath5(t *testing.T) {
	input := "/a/../../b/../c//.//"

	expected := "/c"

	actual := simplifyPath(input)

	if actual != expected {
		t.Errorf("Got %s, want %s", actual, expected)
	}
}

func TestSimplifyPath6(t *testing.T) {
	input := "/a//b////c/d//././/.."

	expected := "/a/b/c"

	actual := simplifyPath(input)

	if actual != expected {
		t.Errorf("Got %s, want %s", actual, expected)
	}
}

type stack struct {
	Values []string
}

func (s *stack) IsEmpty() bool {
	if len(s.Values) > 0 {
		return false
	}
	return true
}

func (s *stack) Push(dir string) {
	s.Values = append(s.Values, dir)
}

func (s *stack) Pop() string {
	if len(s.Values) > 0 {
		popped := s.Values[len(s.Values)-1]
		s.Values = s.Values[:len(s.Values)-1]
		return popped
	} else {
		return ""
	}
}

func simplifyPath(path string) string {
	dirStack := stack{}
	skip := map[string]bool{
		".":  true,
		"..": true,
		"":   true,
	}

	for _, dir := range strings.Split(path, "/") {
		if dir == ".." && !dirStack.IsEmpty() {
			dirStack.Pop()
			continue
		}

		if _, ok := skip[dir]; !ok {
			dirStack.Push(dir)
			continue
		}
	}

	result := ""

	for !dirStack.IsEmpty() {
		result = "/" + dirStack.Pop() + result
	}

	if len(result) == 0 {
		return "/"
	}

	return result
}
