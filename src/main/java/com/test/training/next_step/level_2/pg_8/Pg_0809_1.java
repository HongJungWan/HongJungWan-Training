package com.test.training.next_step.level_2.pg_8;

import java.util.ArrayList;
import java.util.List;

/**
 * level 2 - 영어 끝말잇기
 */
public class Pg_0809_1 {
    public static void main(String[] args) {
        String[] wordsOne = {"tank", "kick", "know", "wheel", "land", "dream", "mother", "robot", "tank"};
        for (int answer : solution(3, wordsOne)) {
            System.out.print(answer + " ");
        }

        System.out.println();

        String[] wordsTwo = {"hello", "one", "even", "never", "now", "world", "draw"};
        for (int answer : solution(2, wordsTwo)) {
            System.out.print(answer + " ");
        }
    }

    public static int[] solution(int n, String[] words) {
        int[] answer = new int[2];
        List<String> list = new ArrayList<>();

        for (int i = 0; i < words.length; i++) {
            if (!list.isEmpty() && list.contains(words[i])) {
                answer[0] = (i % n) + 1; // 몇번 째, 사람인지
                answer[1] = (i / n) + 1; // 몇번 째, 차례인지
                return answer;
            }
            if (!list.isEmpty() && words[i - 1].charAt(words[i - 1].length() - 1) != words[i].charAt(0)) {
                answer[0] = (i % n) + 1;
                answer[1] = (i / n) + 1;
                return answer;
            }
            list.add(words[i]);
        }
        return new int[]{0, 0};
    }
}