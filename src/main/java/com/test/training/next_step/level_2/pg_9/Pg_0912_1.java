package com.test.training.next_step.level_2.pg_9;

import java.util.Stack;

/**
 * level 2 - 괄호 회전하기
 */
public class Pg_0912_1 {
    public static void main(String[] args) {
        System.out.println(solution("[](){}")); // 3
        System.out.println(solution("}]()[{")); // 2
        System.out.println(solution("[)(]"));   // 0
    }

    static public int solution(String s) {
        int answer = 0;

        for (int i = 0; i < s.length(); i++) {

            // 문자열 회전 로직
            String str = s.substring(i, s.length()) + s.substring(0, i);
            Stack<Character> stack = new Stack<>();

            for (int j = 0; j < str.length(); j++) {
                char c = str.charAt(j);
                if (stack.isEmpty()) {
                    stack.push(c);
                } else if (c == ')' && stack.peek() == '(') {
                    stack.pop();
                } else if (c == '}' && stack.peek() == '{') {
                    stack.pop();
                } else if (c == ']' && stack.peek() == '[') {
                    stack.pop();
                } else {
                    stack.push(c);
                }
            }
            if (stack.isEmpty()) {
                answer++;
            }
        }
        return answer;
    }
}