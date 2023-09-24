package com.test.training.next_step.level_0.pg_9;

import java.util.ArrayList;
import java.util.List;

public class Pg_0925_2 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{4, 2, 6, 1, 7, 6}, 2));
    }

    public static int[] solution(int[] num_list, int n) {
        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < num_list.length; i += n) {
            list.add(num_list[i]);
        }

        return list.stream()
                .mapToInt(Integer::intValue)
                .toArray();
    }
}