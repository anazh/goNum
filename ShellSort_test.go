// ShellSort_test
/*
------------------------------------------------------
作者   : Black Ghost
日期   : 2019-03-05
版本   : 0.0.0
------------------------------------------------------
    希尔（Shell）排序法
理论：
    时间复杂度: O(n^1.3)
    最好情况  : O(n)
    最坏情况  : O(n^2)
    空间复杂度: O(1)
    稳定性    : 不稳定
------------------------------------------------------
输入   :
    in      输入矩阵, 1xn
输出   :
    sol     排序结果
    err     解出标志：false-未解出或达到步数上限；
                     true-全部解出
------------------------------------------------------
*/

package goNum_test

import (
	"testing"

	"github.com/anazh/goNum"
)

// ShellSort 希尔（Shell）排序法
func ShellSort(in goNum.Matrix) (goNum.Matrix, bool) {
	/*
	      希尔（Shell）排序法
	   输入   :
	       in      输入矩阵, 1xn
	   输出   :
	       sol     排序结果
	       err     解出标志：false-未解出或达到步数上限；
	                        true-全部解出
	*/
	//判断初值维数
	if in.Rows != 1 {
		panic("Error in goNum.ShellSort: Input Matrix error")
	}
	if in.Columns < 1 {
		panic("Error in goNum.ShellSort: Empty input Matrix")
	} else if in.Columns == 1 {
		return in, true
	}

	n := in.Columns
	sol := goNum.ZeroMatrix(1, n)
	subn := n / 2
	var err bool = false

	//初始化sol
	for i := 0; i < n; i++ {
		sol.Data[i] = in.Data[i]
	}
	//排序开始
	for ; subn > 0; subn = subn / 2 { //循环到增量减小为1
		for i := subn; i < n; i++ {
			temp := sol.Data[i]
			j := i - subn
			for ; j >= 0 && sol.Data[j] > temp; j -= subn {
				sol.Data[j+subn] = sol.Data[j]
			}
			sol.Data[j+subn] = temp
		}
	}

	err = true
	return sol, err
}

func BenchmarkShellSort(b *testing.B) {
	x61 := goNum.NewMatrix(1, 5, []float64{5.0, 3.2, 1.8, 2.4, 0.1})
	for i := 0; i < b.N; i++ {
		goNum.ShellSort(x61)
	}
}
