package com.test.training.next_step.level_0;

/**
 * level 0 - 문자열 정수의 합
 */
public class Pg_0821_2 {
    public static void main(String[] args) {
        System.out.println(solution("123456789"));
        System.out.println(solution("1000000"));
    }

    public static int solution(String num_str) {
        return num_str.chars()
                .map(Character::getNumericValue)
                .sum();
    }
}
//Character.getNumericValue(char ch) 메서드는 주어진 문자(char ch)의 숫자 값(인트형 값)을 반환