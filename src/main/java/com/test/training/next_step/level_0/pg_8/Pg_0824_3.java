package com.test.training.next_step.level_0.pg_8;

/**
 * level 0 - 원하는 문자열 찾기
 */
public class Pg_0824_3 {
    public static void main(String[] args) {
        System.out.println(solution("AbCdEfG", "aBc"));
        System.out.println(solution("aaAA", "aaaaa"));
    }

    public static int solution(String myString, String pat) {
        if (myString.toLowerCase().contains(pat.toLowerCase())) return 1;
        return 0;
    }
}