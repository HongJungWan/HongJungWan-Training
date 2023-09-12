package com.test.training.next_step.level_2.pg_8;

/**
 * level 2 - 예상 대진표
 */
public class Pg_0814_1 {
    public static void main(String[] args) {
        System.out.println(solution(2, 1, 2));
    }

    public static int solution(int n, int a, int b) {
        int answer = 0;
        if (a == b) return answer;

        while (a != b) {
            a = (a + 1) / 2;
            b = (b + 1) / 2;
            answer++;
        }
        return answer;
    }
}