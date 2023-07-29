package com.test.training.next_step.level_0;

import java.util.ArrayList;
import java.util.List;

/**
 * level 0 - n 번째 원소까지
 */
public class Pg_0729_4 {
    public static void main(String[] args) {
        int[] arr = {2, 1, 6};
        System.out.println(solution(arr, 1));
    }

    public static int[] solution(int[] num_list, int n) {
        List answer = new ArrayList();

        for (int i = 0; i < n; i++) {
            answer.add(num_list[i]);
        }

        return answer.stream()
                .mapToInt(i -> (int) i)
                .toArray();
    }
}