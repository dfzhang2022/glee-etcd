package strategy

import (
    //"fmt"
)

//type Strategy struct {
//    name string
//}
// func (s Strategy) GetName() string {
//     return s.name
// }
// func (s *Strategy) SetName(newName string) {
//     s.name = newName
// }



// 定义策略接口
type Strategy interface {
    Calculate(matrix [][]int) int
}

/* 
Input:
    - Ti_DM: Delay matrix at time i
    - LN: Leader number

Output:
    - IEFlag: Initial exception flag

Variables:
    - IEFlag: Boolean (initialize to False)
    - n: Integer
    - diff_1, diff_2, LNSum, LNSum2: Real numbers
    - T_i_minus_1_DM, T_i_minus_2_DM: Delay matrices

n = CalculateMatrixRank(Ti_DM)
IEFlag = False
(T_i_minus_1_DM, T_i_minus_2_DM) = GetPreviousDelayMatrices()

For col = 1 to n:
    If col ≠ LN:
        diff_1 = diff_1 + (Ti_DM[LN, col] - T_i_minus_1_DM[LN, col])
        LNSum = LNSum + T_i_minus_1_DM[LN, col]

R_1 = diff_1 / LNSum

If R_1 > 0.3:
    R_2 = 0
    For col = 1 to n:
        If col ≠ LN:
        diff_2 = diff_2 + (T_i_minus_1_DM[LN, col] - T_i_minus_2_DM[LN, col])
        LNSum2 = LNSum2 + T_i_minus_2_DM[LN, col]
    R_2 = diff_2 / LNSum2
    If R_2 > 0.3:
        IEFlag = True

Return IEFlag

*/

// 具体策略1：计算二维矩阵中所有元素的和
type SumStrategy struct{}

func (s SumStrategy) Calculate(matrix [][]int) int {
    sum := 0
    for _, row := range matrix {
        for _, num := range row {
            sum += num
        }
    }
    return sum
}

func CalculateIEFlag(Ti_DM [][]int, T1_DM [][]int, T2_DM [][]int, LN int) bool {
    n := len(Ti_DM)
    IEFlag := false
    diff_1 := 0
    diff_2 := 0
    LNSum1 := 0
    LNSum2 := 0

    for col := 0; col < n; col++ {
        if col != LN {
            diff_1 += Ti_DM[LN][col] - T1_DM[LN][col]
            LNSum1 += T1_DM[LN][col]
//          fmt.Println(diff_1)
        }
    }
   
    R_1 := float64(diff_1) / float64(LNSum1)
//  fmt.Println(R_1)
    if R_1 > 0.3 {
        for col := 0; col < n; col++ {
            if col != LN {
                diff_2 += T1_DM[LN][col] - T2_DM[LN][col]
                LNSum2 += T2_DM[LN][col]
            }
        }

        R_2 := float64(diff_2) / float64(LNSum2)
//      fmt.Println(R_2)
        if R_2 > 0.3 {
            IEFlag = true
        }
    }
 
    return IEFlag
}

// 具体策略2：计算二维矩阵中所有元素的平均值
type AverageStrategy struct{}

func (s AverageStrategy) Calculate(matrix [][]int) int {
    sum := 0
    count := 0
    for _, row := range matrix {
        for _, num := range row {
            sum += num
            count++
        }
    }
    if count > 0 {
        return sum / count
    }
    return 0
}

// 上下文，包含策略接口
type Context struct {
    strategy Strategy
}

// 上下文的成员函数，用于设置策略
func (c *Context) SetStrategy(strategy Strategy) {
    c.strategy = strategy
}

// 上下文的成员函数，用于执行策略
func (c Context) ExecuteStrategy(matrix [][]int) int {
    return c.strategy.Calculate(matrix)
}