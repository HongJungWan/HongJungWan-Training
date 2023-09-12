package com.test.training.next_step.level_0.pg_8;

/**
 * level 0 - 길이에 따른 연산
 */
public class Pg_0813_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{3, 4, 5, 2, 5, 4, 6, 7, 3, 7, 2, 2, 1}));
        System.out.println(solution(new int[]{2, 3, 4, 5}));
    }

    public static int solution(int[] num_list) {
        int answer = num_list[0];
        if (num_list.length >= 11) {
            for (int i = 1; i < num_list.length; i++) {
                answer += num_list[i];
            }
        } else if (num_list.length <= 10) {
            for (int i = 1; i < num_list.length; i++) {
                answer *= num_list[i];
            }
        }
        return answer;
    }
}