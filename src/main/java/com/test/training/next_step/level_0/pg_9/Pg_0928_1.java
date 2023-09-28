package com.test.training.next_step.level_0.pg_9;

public class Pg_0928_1 {
    public static void main(String[] args) {
        System.out.println(solution(new String[]{"abc", "def", "ghi"}, "ef"));
    }

    public static String solution(String[] str_list, String ex) {
        String answer = "";
        for (String str : str_list) {
            if (!str.contains(ex)) {
                answer += str;
            }
        }
        return answer;
    }
}