package mymath
import (
    "errors"
)
/*
小写字母开头的函数只在本包内可见，大写字母开头的函数才
能被其他包使用。
这个规则也适用于类型和变量的可见性。
*/

func Add(a int, b int) (ret int, err error) {
    if a < 0 || b < 0 { // 不支持负数
        err = errors.New("should be non-negative numbers!")
        return 
    }
//    return a + b, nil; // same as the three flowing lines
    ret = a + b
    err = nil
    return
}
