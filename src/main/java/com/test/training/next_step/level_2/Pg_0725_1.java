package com.test.training.next_step.level_2;

import java.util.Stack;

/**
 * level 2 - 올바른 괄호
 */
public class Pg_0725_1 {
    public static void main(String[] args) {
        boolean answer = solution("()()");
        System.out.println(answer);
    }

    static boolean solution(String s) {
        Stack<String> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '(') {
                stack.push(s);
            } else if (!stack.isEmpty() && s.charAt(i) == ')') {
                stack.pop();
            } else {
                return false;
            }
        }
        return stack.isEmpty();
    }
}