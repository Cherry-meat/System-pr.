package main

import "math"

func main(){
}

func firstWay(a float64, b float64) (int, error) {
    if a == 0 || b == 0 {
        return 0, nil
    }

    sign := 1
    if (a < 0 && b > 0) || (a > 0 && b < 0) {
        sign = -1
    }

    absA := int(math.Abs(a))
    absB := int(math.Abs(b))

    result := 0
    for i := 0; i < absB; i++ {
        result += absA
    }
    
    return sign * result, nil
}
func secondWay(a float32, b float32)(int, error){
    if b == 0 {
        return 0, nil
    }
    if a == 0 {
        return 0, nil
    }
    result := int( a/(1/b))
    return result, nil
}
func thirdWay(a float64, b float64)(int, error){
    if a == 0 || b == 0 {
        return 0, nil
    }
    sign := 1
    if (a < 0 && b > 0) || (a > 0 && b < 0) {
        sign = -1
    }
    
    absA := math.Abs(a)
    absB := math.Abs(b)
    
    result := math.Pow(2, math.Log2(absA) + math.Log2(absB))
    
    return sign * int(math.Round(result)), nil
}
func fourthWay(a float64, b float64)(int, error){
     if a == 0 || b == 0 {
        return 0, nil
    }
    sign := 1
    if (a < 0 && b > 0) || (a > 0 && b < 0) {
        sign = -1
    }
    
    absA := math.Abs(a)
    absB := math.Abs(b)
    
    result := math.Round(1 - (absA+absB)/(math.Tan((math.Atan(absA)+ math.Atan(absB)))))
    return sign * int(math.Round(result)), nil
}

func fifthWay(a float64, b float64) (int, error) {
    if a == 0 || b == 0 {
        return 0, nil
    }
    
    sign := 1
    if (a < 0 && b > 0) || (a > 0 && b < 0) {
        sign = -1
    }

    absA := int(math.Abs(a))
    absB := int(math.Abs(b))
    
    result := 0
    for absB > 0 {
        if absB&1 == 1 {
            result += absA
        }
        absA <<= 1
        absB >>= 1
    }
    
    return sign * result, nil
}

func sixthWay(a float64, b float64) (int, error) {
    if a == 0 || b == 0 {
        return 0, nil
    }
    if b == 1 {
        return int(a), nil
    }
    if b == -1 {
        return int(-a), nil
    }

    sign := 1
    if (a < 0 && b > 0) || (a > 0 && b < 0) {
        sign = -1
    }

    absA := int(math.Abs(a))
    absB := int(math.Abs(b))

    halfB := absB / 2
    half, _ := sixthWay(float64(absA), float64(halfB))

    remainder := absB % 2
    var result int
    if remainder != 0 {
        result = half + half + absA
    } else {
        result = half + half
    }
    
    return sign * result, nil
}