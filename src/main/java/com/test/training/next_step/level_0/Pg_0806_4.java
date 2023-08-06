package com.test.training.next_step.level_0;

/**
 * level 0 - 접두사인지 확인하기
 */
public class Pg_0806_4 {
    public static void main(String[] args) {
        System.out.println(solution("banana", "ban"));
    }

    public static int solution(String my_string, String is_prefix) {
        if (is_prefix.length() > my_string.length()) return 0;
        for (int i = 0; i < is_prefix.length(); i++) {
            if (my_string.charAt(i) != is_prefix.charAt(i)) {
                return 0;
            }
        }
        return 1;
    }
}