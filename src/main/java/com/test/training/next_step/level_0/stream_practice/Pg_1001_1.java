package com.test.training.next_step.level_0.stream_practice;

import java.util.stream.IntStream;

public class Pg_1001_1 {
    public static void main(String[] args) {
        for (int item : solution(10, 3)) {
            System.out.print(item + " ");
        }
    }

    public static int[] solution(int start, int end_num) {
        return IntStream.rangeClosed(end_num, start)
                .boxed()
                .sorted((a, b) -> b.compareTo(a))
                .mapToInt(e -> e)
                .toArray();
    }
}

// IntStream.rangeClosed(startInclusive, endInclusive) :
// startInclusive에서 시작하여 endInclusive에서 끝나는 연속된 정수 값들로 구성된 스트림을 생성


// boxed() : 원시 타입 스트림을 해당하는 박싱된 타입의 스트림으로 변환
// Java의 IntStream은 원시형 정수 스트림 -> 오름차순 정렬만 지원
// 내림차순 정렬을 위해 스트림의 원소를 박싱하여, Stream<Integer>로 변환한 후에 비교자(Comparator) 사용


// sorted((a, b) -> b.compareTo(a)) : 내림차순 정렬