package com.test.training.next_step.level_0;

/**
 * level 0 - 원소들의 곱과 합
 */
public class Pg_0813_2 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{3, 4, 5, 2, 1}));
        System.out.println(solution(new int[]{5, 7, 8, 3}));
    }

    public static int solution(int[] num_list) {
        int total = num_list[0]; //곱
        int sum = num_list[0]; //합

        for (int i = 1; i < num_list.length; i++) {
            total *= num_list[i];
            sum += num_list[i];
        }

        if (total < sum * sum) return 1;
        return 0;
    }
}