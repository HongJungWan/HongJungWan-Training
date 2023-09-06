package com.test.training.next_step.level_0.stream_practice;

import java.util.Arrays;

/**
 * level 0 - 조건에 맞게 수열 변환하기 3
 */
public class Pg_0824_2 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{1, 2, 3, 100, 99, 98}, 3));
        System.out.println(solution(new int[]{1, 2, 3, 100, 99, 98}, 2));
    }

    public static int[] solution(int[] arr, int k) {
        if (k % 2 == 0) {
            return Arrays.stream(arr)
                    .map(n -> n + k)
                    .toArray();
        } else {
            return Arrays.stream(arr)
                    .map(n -> n * k)
                    .toArray();
        }
    }
}