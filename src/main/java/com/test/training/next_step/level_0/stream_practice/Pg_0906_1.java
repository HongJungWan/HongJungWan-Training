package com.test.training.next_step.level_0.stream_practice;

import java.util.stream.IntStream;

/**
 * level 0 - 배열에서 문자열 대소문자 변환하기
 */
public class Pg_0906_1 {
    public static void main(String[] args) {
        for (String s : solution(new String[]{"aBc", "AbC"})) {
            System.out.print(s + " ");
        }
    }

    public static String[] solution(String[] strArr) {
        return IntStream.range(0, strArr.length)
                .mapToObj(i -> i % 2 == 0 ? strArr[i].toLowerCase() : strArr[i].toUpperCase())
                .toArray(value -> new String[value]);
    }
}