package com.test.training.next_step.level_0;

import java.util.Arrays;
import java.util.stream.Collectors;

/**
 * level 0 - 글자 이어 붙여 문자열 만들기
 */
public class Pg_0824_4 {
    public static void main(String[] args) {
        System.out.println(solution("cvsgiorszzzmrpaqpe", new int[]{16, 6, 5, 3, 12, 14, 11, 11, 17, 12, 7}));
        System.out.println(solution("zpiaz", new int[]{1, 2, 0, 0, 3}));
    }

    public static String solution(String my_string, int[] index_list) {
        return Arrays.stream(index_list)
                .mapToObj(index -> my_string.charAt(index))
                .map(obj -> String.valueOf(obj))
                .collect(Collectors.joining());
    }
}