package com.test.training.next_step.level_2;

/**
 * level 2 - 점프와 순간 이동
 */
public class Pg_0815_1 {
    public static void main(String[] args) {
        System.out.println(solution(5));
        System.out.println(solution(6));
        System.out.println(solution(5000));
    }

    public static int solution(int n) {
        int ans = 0;

        while (n > 0) {
            if (n % 2 == 1) {
                ans++;
            }
            n /= 2;
        }
        return ans;
    }
}