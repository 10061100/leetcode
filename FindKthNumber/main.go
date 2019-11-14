package main

func findKthNumber(n int, k int) int {

    if k > n {
        return -1
    }


    for prefix := 1; prefix <= 9; prefix++ {

        for t := 1; t <= n; {
            // 说明k在这个范围内
            if k <= t {

            } else {
                k = k - t
            }

            t = t * 10
        }
    }



}
