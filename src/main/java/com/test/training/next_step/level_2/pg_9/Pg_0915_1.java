package com.test.training.next_step.level_2.pg_9;

import java.util.Arrays;

/**
 * level 2 - H-Index
 */
public class Pg_0915_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{3, 0, 6, 1, 5}));
    }

    public static int solution(int[] citations) {
        Arrays.sort(citations);

        int hIndex = 0;
        for (int i = citations.length - 1; i >= 0; i--) {
            if (citations[i] > hIndex) {
                hIndex++;
            } else {
                break;
            }
        }
        return hIndex;
    }
}