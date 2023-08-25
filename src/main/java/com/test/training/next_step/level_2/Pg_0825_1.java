package com.test.training.next_step.level_2;

/**
 * level 2 - 멀리 뛰기
 */
public class Pg_0825_1 {
    public static void main(String[] args) {
        System.out.println(solution(4));
        System.out.println(solution(3));
    }

    public static long solution(int n) {
        if (n == 1) return 1;
        if (n == 2) return 2;
        return fibonacci(n);
    }

    public static long fibonacci(int n) {
        long[] dp = new long[n];
        initializeDp(dp);

        for (int i = 2; i < n; i++) {
            dp[i] = (dp[i - 1] + dp[i - 2]) % 1234567;
        }
        return dp[n - 1];
    }

    public static void initializeDp(long[] dp) {
        dp[0] = 1;
        dp[1] = 2;
    }
}

// n번째 칸에 도달하는 방법의 수를 저장하는 동적 프로그래밍 테이블 dp를 사용하여 해결