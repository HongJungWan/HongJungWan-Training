package com.test.training.next_step.level_2.pg_9;

/**
 * level 2 - n^2 배열 자르기
 */
public class Pg_0913_1 {
    public static void main(String[] args) {
        System.out.println(solution(3, 2, 5));  // [3,2,2,3]
        System.out.println(solution(4, 7, 14)); // [4,3,3,3,4,4,4,4]
    }

    public static int[] solution(int n, long left, long right) {
        int len = (int) right - (int) left;
        int[] answer = new int[len + 1];

        int idx = 0;
        for (long i = left; i <= right; i++) {
            long row = i / n;
            long col = i % n;
            answer[idx++] = Math.max((int) row, (int) col) + 1;
        }
        return answer;
    }
}