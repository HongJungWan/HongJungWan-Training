package com.test.training.next_step.level_2.pg_8;

import java.util.Arrays;

/**
 * level 2 - 구명보트
 */
public class Pg_0813_3 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{70, 50, 80, 50}, 100));
        System.out.println(solution(new int[]{70, 80, 50}, 100));
    }

    public static int solution(int[] people, int limit) {
        int answer = 0;
        int left = 0;
        int right = people.length - 1;

        Arrays.sort(people);

        while (left <= right) {
            if ((people[left] + people[right]) <= limit) {
                left++;
                right--;
                answer++;
            } else {
                right--;
                answer++;
            }
        }
        return answer;
    }
}