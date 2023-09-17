package com.test.training.next_step.level_2.pg_9;

import java.util.Arrays;

/**
 * level 2 - 가장 큰 수
 */
public class Pg_0917_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{6, 10, 2}));
    }

    public static String solution(int[] numbers) {
        String[] temp = new String[numbers.length];

        for (int i = 0; i < numbers.length; i++) {
            temp[i] = String.valueOf(numbers[i]);
        }

        /**
         * (o2 + o1).compareTo(o1 + o2)
         * o2 + o1과 o1 + o2를 사전 순서로 비교하여 내림차순으로 정렬하는 역할
         */
        Arrays.sort(temp, (o1, o2) -> (o2 + o1).compareTo(o1 + o2));
        if (temp[0].equals("0")) return "0";

        StringBuilder stringBuilder = new StringBuilder();
        for (int i = 0; i < numbers.length; i++) {
            stringBuilder.append(temp[i]);
        }
        return stringBuilder.toString();
    }
}
// Comparator는 Java에서 객체들을 정렬하는 데 사용되는 인터페이스
// 1. 외부에서 정렬 방법을 정의하고 싶을 때 사용
// 2. int compare(T o1, T o2)
// 3. 특정 조건에 따른 사용자 지정 정렬
// 4. 람다 표현식 or 별도의 클래스로 현

// Comparator.compareTo() 메소드를 오버라이드하여 원하는 정렬 기준을 지정

// (o2 + o1).compareTo(o1 + o2) : 내림차순
// (o1 + o2).compareTo(o2 + o1) : 오름차순