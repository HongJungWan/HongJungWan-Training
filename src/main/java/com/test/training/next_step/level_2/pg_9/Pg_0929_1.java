package com.test.training.next_step.level_2.pg_9;

import java.util.Collections;
import java.util.PriorityQueue;

/**
 * level 2 - 프로세스
 */
public class Pg_0929_1 {
    public static void main(String[] args) {
        System.out.println(solution(new int[]{2, 1, 3, 2}, 2));
    }

    public static int solution(int[] priorities, int location) {
        int answer = 0;
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());

        for (int num : priorities) {
            pq.add(num);
        }

        System.out.println(pq);

        while (!pq.isEmpty()) {
            for (int i = 0; i < priorities.length; i++) {
                if (priorities[i] == pq.peek()) {
                    pq.poll();
                    answer++;
                    if (i == location)
                        return answer;
                }
            }
        }
        return answer;
    }
}

// Java의 PriorityQueue는 기본적으로 natural ordering에 따라 요소들이 정렬,
// 즉, 작은 숫자부터 큰 숫자 순으로, 문자열은 사전 순으로 정렬
// 그래서, Collections.reverseOrder() -> 큰 값을 기준으로 우선 순위를 높게 선정하기 위해.