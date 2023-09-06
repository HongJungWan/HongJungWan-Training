package com.test.training.next_step.level_0.stream_practice;

import java.util.ArrayList;
import java.util.List;

/**
 * level 0 - n 번째 원소부터
 */
public class Pg_0806_2 {
    public static void main(String[] args) {
        int[] num_list = {2, 1, 6};
        System.out.println(solution(num_list, 3));
    }

    public static int[] solution(int[] num_list, int n) {
        List list = new ArrayList<>();
        for (int i = n - 1; i < num_list.length; i++) {
            list.add(num_list[i]);
        }
        return list.stream()
                .mapToInt(e -> (int) e)
                .toArray();
    }
}