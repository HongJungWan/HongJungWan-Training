package com.test.training.next_step.level_2.pg_7;

/**
 * level 2 - 숫자의 표현
 */
public class Pg_0730_1 {
    public static void main(String[] args) {
        System.out.println(solution(15));
    }

    public static int solution(int n) {
        int answer = 1, sum = 0;

        for (int j = 1; j <= n / 2; j++) {
            sum = 0;
            for (int i = j; i <= (n / 2) + 1; i++) {
                sum += i;
                if (sum == n) {
                    answer++;
                    break;
                }
                if (sum > n) {
                    break;
                }
            }
        }
        return answer;
    }
}
/**
 * 야매로 효율성 통과한 느낌..
 */