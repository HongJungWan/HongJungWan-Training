package com.test.training.next_step.level_0.pg_9;

import java.util.ArrayList;
import java.util.List;

/**
 * level 0 - 공백으로 구분하기 1
 */
public class Pg_0920_1 {
    public static void main(String[] args) {
        System.out.println(solution("i love you"));
    }

    public static String[] solution(String my_string) {
        List<String> words = new ArrayList<>();
        int startIdx = 0;
        for (int i = 0; i < my_string.length(); i++) {
            if (my_string.charAt(i) == ' ') {
                words.add(my_string.substring(startIdx, i));
                startIdx = i + 1;
            }
        }
        words.add(my_string.substring(startIdx));
        return words.toArray(new String[0]);
    }
}