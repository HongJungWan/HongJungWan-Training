package com.test.training.next_step.level_0.pg_10;

import java.util.Arrays;

public class Pg_1003_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{49, 13}, new int[]{70, 11, 2}));
        System.out.println(solution(new int[]{100, 17, 84, 1}, new int[]{55, 12, 65, 36}));
        System.out.println(solution(new int[]{1, 2, 3, 4, 5}, new int[]{3, 3, 3, 3, 3}));
    }

    public static int solution(int[] arr1, int[] arr2) {
        if (arr1.length > arr2.length) return 1;
        if (arr2.length > arr1.length) return -1;

        if (arr1.length == arr2.length) {
            int tempSum1 = 0;
            int tempSum2 = 0;
            tempSum1 = Arrays.stream(arr1)
                    .sum();
            tempSum2 = Arrays.stream(arr2)
                    .sum();
            if (tempSum1 > tempSum2) return 1;
            if (tempSum1 < tempSum2) return -1;
            if (tempSum1 == tempSum2) return 0;
        }
        return 0;
    }
}