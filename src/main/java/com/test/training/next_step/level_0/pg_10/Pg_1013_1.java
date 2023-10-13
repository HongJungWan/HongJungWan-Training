package com.test.training.next_step.level_0.pg_10;

import java.util.stream.IntStream;

public class Pg_1013_1 {
    public static void main(String[] args) {
        for (int item : solution(10, 3)) {
            System.out.print(item + " ");
        }
    }

    public static int[] solution(int start, int end_num) {
        return IntStream.rangeClosed(end_num, start)
                .boxed()
                .sorted((a, b) -> b.compareTo(a))
                .mapToInt(e -> e)
                .toArray();
    }
}