package com.test.training.next_step.level_1;

/**
 * level 1 - 숫자 문자열과 영단어
 */
public class Pg_0808_1 {
    public static void main(String[] args) {
        System.out.println(solution("one4seveneight"));
    }

    public static int solution(String s) {
        String temp[] = {"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};
        for (int i = 0; i < temp.length; i++) {
            if (s.contains(temp[i])) {
                s = s.replace(temp[i], String.valueOf(i));
            }
        }
        return Integer.parseInt(s);
    }
}