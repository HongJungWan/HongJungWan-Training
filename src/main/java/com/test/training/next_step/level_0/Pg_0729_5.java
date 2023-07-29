package com.test.training.next_step.level_0;

/**
 * level 0 - n보다 커질 때까지 더하기
 */
public class Pg_0729_5 {
    public static void main(String[] args) {
        int arr[] = {34, 5, 71, 29, 100, 34};
        System.out.println(solution(arr, 123));
    }

    public static int solution(int[] numbers, int n) {
        int answer = 0;
        for (int i = 0; i < numbers.length; i++) {
            if (answer > n) {
                return answer;
            }
            answer += numbers[i];
        }
        return answer;
    }
}