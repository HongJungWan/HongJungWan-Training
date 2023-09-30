package com.test.training.next_step.level_2.pg_9;

import java.util.Collections;
import java.util.PriorityQueue;

/**
 * level 2 - 프로세스 (재도전)
 */
public class Pg_0930_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{2, 1, 3, 2}, 2));
        System.out.println(solution(new int[]{1, 1, 9, 1, 1, 1}, 0));
    }

    public static int solution(int[] priorities, int location) {
        int count = 0;

        PriorityQueue<Integer> priorityQueue = new PriorityQueue(Collections.reverseOrder());
        for (int item : priorities) {
            priorityQueue.add(item);
        }

        while (!priorityQueue.isEmpty()) {
            for (int i = 0; i < priorities.length; i++) {
                if (priorities[i] == priorityQueue.peek()) {
                    priorityQueue.poll();
                    count++;
                    if (location == i) {
                        return count;
                    }
                }
            }
        }
        return count;
    }
}