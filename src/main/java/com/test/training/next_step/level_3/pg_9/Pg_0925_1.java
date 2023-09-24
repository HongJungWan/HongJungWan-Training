package com.test.training.next_step.level_3.pg_9;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

/**
 * level 3 - 베스트앨범
 * <p>
 * 다시 풀기.
 */
public class Pg_0925_1 {
    public static void main(String[] args) {
        System.out.println(solution(
                new String[]{"classic", "pop", "classic", "classic", "pop"},
                new int[]{500, 600, 150, 800, 2500}));
    }

    static class Song {
        int index;
        int play;

        public Song(int index, int play) {
            this.index = index;
            this.play = play;
        }
    }

    public static int[] solution(String[] genres, int[] plays) {
        Map<String, Integer> genrePlayCount = new HashMap<>();
        Map<String, List<Song>> genreSongs = new HashMap<>();

        for (int i = 0; i < genres.length; i++) {

            // 1. 장르별 총 재생 횟수를 계산
            genrePlayCount.put(genres[i], genrePlayCount.getOrDefault(genres[i], 0) + plays[i]);

            // 2. 각 장르에 속한 노래들을 재생 횟수와 고유 번호에 따라 저장
            List<Song> songs = genreSongs.getOrDefault(genres[i], new ArrayList<>());
            songs.add(new Song(i, plays[i]));
            genreSongs.put(genres[i], songs);
        }

        // 3. 각 장르에서 상위 두 개의 노래를 선택
        List<Integer> answerList = createAnswerList(genrePlayCount, genreSongs);

        return answerList.stream()
                .mapToInt(Integer::intValue)
                .toArray();
    }

    private static List<Integer> createAnswerList(Map<String, Integer> genrePlayCount, Map<String, List<Song>> genreSongs) {
        return genrePlayCount.entrySet().stream()
                .sorted((a, b) -> b.getValue().compareTo(a.getValue())) // 장르별 총 재생 횟수에 따라 내림차순 정렬
                .flatMap(entry -> genreSongs.get(entry.getKey()).stream()
                        .sorted((a, b) -> b.play == a.play ? a.index - b.index : b.play - a.play) // 노래를 재생 횟수와 고유 번호에 따라 정렬
                        .limit(2) // 상위 두 개의 노래만 선택
                        .map(song -> song.index))
                .collect(Collectors.toList());
    }
}

// flatMap은 스트림의 각 요소를 여러 개의 객체로 변환할 수 있는 중간 연산
//
// map과 유사하지만, map은 스트림의 각 요소를 하나의 객체로만 변환
// flatMap은 스트림의 각 요소를 여러 개의 객체로 변환한 후, 이를 하나의 스트림으로 연결