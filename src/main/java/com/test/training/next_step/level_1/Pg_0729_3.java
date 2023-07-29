package com.test.training.next_step.level_1;

/**
 * level 1 - 3진법 뒤집기
 */
public class Pg_0729_3 {
    public static void main(String[] args) {
        System.out.println(solution(45));
        System.out.println(solution(125));
    }

    public static int solution(int n) {
        String answer = "";
        StringBuilder sb = new StringBuilder();

        // 10진법 -> 3진법으로 변환
        answer = Integer.toString(n, 3);

        // 앞뒤 뒤집기
        answer = sb.append(answer)
                .reverse()
                .toString();

        // 3진법 -> 10진법으로 변환
        return Integer.parseInt(answer, 3);
    }
}