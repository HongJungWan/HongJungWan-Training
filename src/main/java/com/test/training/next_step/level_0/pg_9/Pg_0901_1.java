package com.test.training.next_step.level_0.pg_9;

/**
 * level 0 - 조건에 맞게 수열 변환하기 1
 */
public class Pg_0901_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{1, 2, 3, 100, 99, 98}));
    }

    public static int[] solution(int[] arr) {
        int[] answer = new int[arr.length];
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] < 50 && arr[i] % 2 != 0) {
                arr[i] = arr[i] * 2;
            } else if (arr[i] >= 50 && arr[i] % 2 == 0) {
                arr[i] = arr[i] / 2;
            }
            answer[i] = arr[i];
        }
        return answer;
    }
}