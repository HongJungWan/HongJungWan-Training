package com.test.training.next_step.level_0;

import java.util.stream.IntStream;

/**
 * level 0 - 첫 번째로 나오는 음수
 */
public class Pg_0729_1 {
    public static void main(String[] args) {
        int arr[] = {12, 4, 15, 46, 38, -2, 15};
        System.out.println(solution(arr));
    }

    public static int solution(int[] num_list) {
        return IntStream.range(0, num_list.length)
                .filter(i -> num_list[i] < 0)
                .findFirst()
                .orElse(-1);
    }
}