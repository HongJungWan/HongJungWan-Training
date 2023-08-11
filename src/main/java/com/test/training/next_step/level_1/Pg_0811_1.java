package com.test.training.next_step.level_1;

import java.util.Arrays;
import java.util.Set;
import java.util.TreeSet;

/**
 * level 1 - 두 개 뽑아서 더하기
 */
public class Pg_0811_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{2, 1, 3, 4, 1}));
        System.out.println(solution(new int[]{5, 0, 2, 7}));
    }

    public static int[] solution(int[] numbers) {
        Set set = new TreeSet();
        Arrays.sort(numbers);

        for (int i = 0; i < numbers.length; i++) {
            for (int j = i + 1; j < numbers.length; j++) {
                if (i != j) {
                    set.add(numbers[i] + numbers[j]);
                }
            }
        }
        return set.stream()
                .mapToInt(e -> (int) e)
                .toArray();
    }
}