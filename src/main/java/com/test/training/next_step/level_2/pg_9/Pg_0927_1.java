package com.test.training.next_step.level_2.pg_9;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * level 2 - 기능개발
 */
public class Pg_0927_1 {
    public static void main(String[] args) {
        for (int answer : solution(new int[]{93, 30, 55}, new int[]{1, 30, 5})) {
            System.out.print(answer + ", ");
        }

        System.out.println();
        System.out.println();

        for (int answer : solution(new int[]{95, 90, 99, 99, 80, 99}, new int[]{1, 1, 1, 1, 1, 1})) {
            System.out.print(answer + ", ");
        }
    }

    public static int[] solution(int[] progresses, int[] speeds) {
        Queue<Integer> releaseDays = new LinkedList<>();
        List<Integer> answerList = new ArrayList<>();

        // progresses 마다 배포 날짜 구하기.
        for (int i = 0; i < progresses.length; i++) {
            int remaining = 100 - progresses[i];
            // a를 b로 나누었을 때 올림 공식 == (a + b - 1) / b
            int day = (remaining + speeds[i] - 1) / speeds[i];
            releaseDays.add(day);
        }

        // 배포날에 한꺼번에 배포되는 수
        while (!releaseDays.isEmpty()) {
            int releaseDay = releaseDays.poll();
            int count = 1;

            // 배포 날짜가 같거나 더 이른 기능들을 모두 포함하여 카운트
            while (!releaseDays.isEmpty() && releaseDays.peek() <= releaseDay) {
                count++;
                releaseDays.poll();
            }
            answerList.add(count);
        }

        return answerList.stream()
                .mapToInt(e -> e)
                .toArray();
    }
}