package com.test.training.next_step.level_0;

/**
 * level 0 - 부분 문자열
 */
public class Pg_0824_1 {
    public static void main(String[] args) {
        System.out.println(solution("abc", "aabcc"));
        System.out.println(solution("tbt", "tbbttb"));
    }

    public static int solution(String str1, String str2) {
        if (str2.contains(str1)) return 1;
        return 0;
    }
}