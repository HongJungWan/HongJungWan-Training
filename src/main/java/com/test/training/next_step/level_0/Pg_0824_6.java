package com.test.training.next_step.level_0;

import java.util.Arrays;

/**
 * level 0 - 뒤에서 5등 위로
 */
public class Pg_0824_6 {
    public static void main(String[] args) {
        System.out.println(new int[]{12, 4, 15, 46, 38, 1, 14, 56, 32, 10});
    }

    public static int[] solution(int[] num_list) {
        return Arrays.stream(num_list)
                .sorted()
                .skip(5)
                .toArray();
    }
}