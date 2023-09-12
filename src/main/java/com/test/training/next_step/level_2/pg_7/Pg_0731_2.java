package com.test.training.next_step.level_2.pg_7;

/**
 * level 2 - 피보나치 수
 */
public class Pg_0731_2 {
    public static void main(String[] args) {
        System.out.println(solution(3));
    }

    public static int solution(int n) {
        if (n == 0) return 0;
        if (n == 1) return 1;

        int arr[] = new int[n + 1];
        arr[0] = 0;
        arr[1] = 1;
        for (int i = 2; i < n + 1; i++) {
            arr[i] = (arr[i - 2] + arr[i - 1]) % 1234567;
        }
        return arr[n];
    }
}
// 0 1 1 2 3 5