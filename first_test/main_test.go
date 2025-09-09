package main

import "testing"

func TestFirstWay(t *testing.T){
   tests := []struct {
        a, b  float64
        expected int
    }{
        {6, 3, 18},
        {0, 5, 0},
        {7, 0, 0},
        {0, 0, 0},
        {-6, 3, -18},
        {6, -3, -18},
        {-6, -3, 18},
        {1, 5, 5},
        {5, 1, 5},
        {1, -1, -1},
        {-1, 1, -1},
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
        {-1, -1, 1},
    }

    for _, test := range tests {
        actual, err := firstWay(test.a, test.b)
        
        if err != nil {
            t.Errorf("firstWay(%f, %f) unexpected error: %v", test.a, test.b, err)
            continue
        }
        if actual != test.expected {
            t.Errorf("firstWay(%f, %f) = %d, want: %d", test.a, test.b, actual, test.expected)
        }
    }
}
func TestSecondWay(t *testing.T){
  tests := []struct {
        a, b  float32
        expected int
    }{
        {6, 3, 18},
        {0, 5, 0},
        {7, 0, 0},
        {0, 0, 0},
        {-6, 3, -18},
        {6, -3, -18},
        {-6, -3, 18},
        {1, 5, 5},
        {5, 1, 5},
        {1, -1, -1},
        {-1, 1, -1},
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
        {-1, -1, 1},
    }

    for _, test := range tests {
        actual, err := secondWay(test.a, test.b)
        
        if err != nil {
            t.Errorf("secondWay(%f, %f) unexpected error: %v", test.a, test.b, err)
            continue
        }
        if actual != test.expected {
            t.Errorf("secondWay(%f, %f) = %d, want: %d", test.a, test.b, actual, test.expected)
        }
    }
}
func TestThirdWay(t *testing.T){
    tests := []struct {
        a, b  float64
        expected int
    }{
        {6, 3, 18},
        {0, 5, 0},
        {7, 0, 0},
        {0, 0, 0},
        {-6, 3, -18},
        {6, -3, -18},
        {-6, -3, 18},
        {1, 5, 5},
        {5, 1, 5},
        {1, -1, -1},
        {-1, 1, -1},
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
        {-1, -1, 1},
    }

    for _, test := range tests {
        actual, err := thirdWay(test.a, test.b)
        
        if err != nil {
            t.Errorf("thirdWay(%f, %f) unexpected error: %v", test.a, test.b, err)
            continue
        }
        if actual != test.expected {
            t.Errorf("thirdWay(%f, %f) = %d, want: %d", test.a, test.b, actual, test.expected)
        }
    }
}
func TestFourtWay(t *testing.T){
   tests := []struct {
        a, b  float64
        expected int
    }{
        {6, 3, 18},
        {0, 5, 0},
        {7, 0, 0},
        {0, 0, 0},
        {-6, 3, -18},
        {6, -3, -18},
        {-6, -3, 18},
        {1, 5, 5},
        {5, 1, 5},
        {1, -1, -1},
        {-1, 1, -1},
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
        {-1, -1, 1},
    }

    for _, test := range tests {
        actual, err := fourthWay(test.a, test.b)
        
        if err != nil {
            t.Errorf("fourthWay(%f, %f) unexpected error: %v", test.a, test.b, err)
            continue
        }
        if actual != test.expected {
            t.Errorf("fourthWay(%f, %f) = %d, want: %d", test.a, test.b, actual, test.expected)
        }
    }
}
func TestFifthWay(t *testing.T) {
    tests := []struct {
        a, b     float64
        expected int
    }{
        {6, 3, 18},
        {0, 5, 0},
        {7, 0, 0},
        {0, 0, 0},
        {-6, 3, -18},
        {6, -3, -18},
        {-6, -3, 18},
        {1, 5, 5},
        {5, 1, 5},
        {1, -1, -1},
        {-1, 1, -1},
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
        {-1, -1, 1},
    }

    for _, test := range tests {
        actual, err := fifthWay(test.a, test.b)
        
        if err != nil {
            t.Errorf("fifthWay(%f, %f) unexpected error: %v", test.a, test.b, err)
            continue
        }
        if actual != test.expected {
            t.Errorf("fifthWay(%f, %f) = %d, want: %d", test.a, test.b, actual, test.expected)
        }
    }
}
func TestSixthWay(t *testing.T){
 tests := []struct {
        a, b     float64
        expected int
    }{
        {6, 3, 18},
        {0, 5, 0},
        {7, 0, 0},
        {0, 0, 0},
        {-6, 3, -18},
        {6, -3, -18},
        {-6, -3, 18},
        {1, 5, 5},
        {5, 1, 5},
        {1, -1, -1},
        {-1, 1, -1},
        {1, 1, 1},
        {1, 0, 0},
        {0, 1, 0},
        {-1, -1, 1},
    }

    for _, test := range tests {
        actual, err := sixthWay(test.a, test.b)
        
        if err != nil {
            t.Errorf("sixthWay(%f, %f) unexpected error: %v", test.a, test.b, err)
            continue
        }
        if actual != test.expected {
            t.Errorf("sixthWay(%f, %f) = %d, want: %d", test.a, test.b, actual, test.expected)
        }
    }
}