package com.test.training.next_step.level_0.stream_practice;

import java.util.stream.Collectors;

/**
 * level 0 - rny_string
 */
public class Pg_0821_3 {
    public static void main(String[] args) {
        System.out.println(solution("masterpiece"));
        System.out.println(solution("programmers"));
        System.out.println(solution("jerry"));
        System.out.println(solution("burn"));
    }

    public static String solution(String rny_string) {
        return rny_string.chars()
                .mapToObj(c -> c == 'm' ? "rn" : String.valueOf((char) c))
                .collect(Collectors.joining());
    }
}
// 삼항 연산자 : 조건 ? 값1 : 값2