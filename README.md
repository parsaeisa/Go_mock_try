# Go_mock_try

For writing unit tests I had to mock the dependencies of a method which I was writing test for . 
go mock was very usefull , it generates mocks of golang interface's easily just in one line terminal code . 

```bash
mockgen -source={the file included interface} -destination={the file you want to put the mock in }
```

you can download and install it with : 

```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

also gomock has very efficient tools to write expectaions of a method 
( test inputs , outputs and test how many times this method has been called ) 
for example :


```go
func TestFoo(t *testing.T) {
  ctrl := gomock.NewController(t)

  // Assert that Bar() is invoked.
  defer ctrl.Finish()

  m := NewMockFoo(ctrl)

  // Asserts that the first and only call to Bar() is passed 99.
  // Anything else will fail.
  m.
    EXPECT().
    Bar(gomock.Eq(99)).
    Return(101).
    Times(2)

  SUT(m)
}
```

You can read more in :   https://github.com/golang/mock


