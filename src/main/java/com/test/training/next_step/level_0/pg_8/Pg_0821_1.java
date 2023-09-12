package com.test.training.next_step.level_0.pg_8;

/**
 * level 0 - 부분 문자열인지 확인하기
 */
public class Pg_0821_1 {
    public static void main(String[] args) {
        System.out.println(solution("banana", "ana"));
        System.out.println(solution("banana", "wxyz"));
    }

    public static int solution(String my_string, String target) {
        if (my_string.contains(target)) return 1;
        return 0;
    }
}