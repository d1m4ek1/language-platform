<template>
  <div class='lesson-component'>
    <div class='lesson-component-header'>
      <p>Урок #{{ lessonNumber }}</p>
      <button class='btn-red' @click='removeLesson'>Удалить урок</button>
    </div>
    <div class='lesson-exercise'>
      <div class='title'>
        <div class='contenteditable' contenteditable='true' @input='setLessonName'><h1 v-html='textData.name'></h1></div>
      </div>
      <div class='exercise-list'>
        <div v-for='(item, idx) in lessonData.exerciseItems' :key='`exercise_${idx}`' class='exercise-item'>
          <!--Test Auto-->
          <ExerciseTestAuto v-if='item.type === "testAuto"' :exercise-number='item.exerciseNumber'
                            :is-edit-course='isEditCourse' :lesson-number='lessonNumber'></ExerciseTestAuto>
          <!--Image-->
          <ExerciseImage v-if='item.type === "image"' :exercise-number='item.exerciseNumber'
                         :is-edit-course='isEditCourse' :lesson-number='lessonNumber'
                         :path='item.path'
                         @on-save-file='onSaveFile' @on-remove-file='onRemoveFile'></ExerciseImage>
          <!--Video-->
          <ExerciseVideo v-if='item.type === "video"' :exercise-number='item.exerciseNumber'
                         :is-edit-course='isEditCourse' :lesson-number='lessonNumber'
                         :path='item.path'
                         @on-save-file='onSaveFile' @on-remove-file='onRemoveFile'></ExerciseVideo>
          <!--Audio-->
          <ExerciseAudio v-if='item.type === "audio"' :exercise-number='item.exerciseNumber'
                         :is-edit-course='isEditCourse'
                         :lesson-number='lessonNumber' :path='item.path'
                         @on-save-file='onSaveFile' @on-remove-file='onRemoveFile'></ExerciseAudio>
        </div>
      </div>
      <div class='task-panel'>
        <button class='btn-blue'>Добавить задание ></button>
        <div class='task-list'>
          <button class='btn-blue' @click='addExercise("video")'>Видео-задание</button>
          <button class='btn-blue' @click='addExercise("audio")'>Аудио-задание</button>
          <button class='btn-blue' @click='addExercise("image")'>Задание с изображением</button>
          <button class='btn-blue' @click='addExercise("testAuto")'>Тест с авто-проверкой</button>
          <!--        <button class='btn-blue'>Тест с ручной проверкой</button>-->
        </div>
      </div>
    </div>
    <div class='lesson-task'>

    </div>
  </div>
</template>

<script>
import ExerciseTestAuto from './ExerciseTestAuto.vue';
import { useCourse } from '../../stores/useCourse';
import ExerciseImage from './ExerciseImage.vue';
import APIFileLoader from '../../models/fileLoader';
import ExerciseVideo from './ExerciseVideo.vue';
import ExerciseAudio from './ExerciseAudio.vue';

const store = useCourse();

export default {
  name: "AddLesson",
  components: { ExerciseAudio, ExerciseVideo, ExerciseImage, ExerciseTestAuto },
  props: {
    lessonNumber: Number,
    courseId: Number,
    isUploadLessonFiles: Boolean,
    isEditCourse: Boolean,
    lessonId: Number | undefined
  },
  data() {
    return {
      lessonData: {
        name: '',
        exerciseItems: [],
        deleteFiles: {}
      },
      textData: {
        name: "Заголовок",
      },
      lessonFiles: new FormData(),
      exerciseFilesIDS: {},
      timeoutWatchLessonItems: null,
    };
  },
  watch: {
    isEditCourse: {
      immediate: true,
      handler(n) {
        if (n) {
          Object.keys(store.lessonItems[this.lessonNumber-1]).forEach(key => {
            this.lessonData[key] = store.lessonItems[this.lessonNumber-1][key]

            if (key === "name") this.textData[key] = this.lessonData[key]
          })
        }
      }
    },
    lessonData: {
      immediate: true,
      deep: true,
      handler(n) {
        clearTimeout(this.timeoutWatchLessonItems);
        this.timeoutWatchLessonItems = setTimeout(() => {
          store.$patch((state) => {
            state.lessonItems[this.lessonNumber - 1] = { ...state.lessonItems[this.lessonNumber - 1], ...n };
          });
        }, 1000);
      }
    },
    async isUploadLessonFiles(n) {
      if (n) {
        const responseFileLoader = await APIFileLoader.fileLoad(this.lessonFiles, 'lesson_files',
          `lesson_id=${this.lessonNumber}`, `exercise_ids=${JSON.stringify(this.exerciseFilesIDS)}`, `course_id=${this.courseId}`);
        const jsonFileLoader = await responseFileLoader.json();
        if (!jsonFileLoader.success) return this.errorCreateCourse = jsonFileLoader;

        const exerciseFiles = JSON.parse(jsonFileLoader.exerciseFiles);

        for (const exerciseId of Object.keys(exerciseFiles)) {
          if (this.isEditCourse && this.lessonData.exerciseItems[+exerciseId - 1].path !== "") {
            this.lessonData.deleteFiles[exerciseId] = this.lessonData.exerciseItems[+exerciseId - 1].path
          }

          this.lessonData.exerciseItems[+exerciseId - 1].path = exerciseFiles[exerciseId];
        }
      }
    }
  },
  methods: {
    onSaveFile(data) {
      this.lessonFiles.append(`exercise_${data.exerciseNumber}`, data.file);
      this.exerciseFilesIDS[+data.exerciseNumber] = '';
    },
    onRemoveFile(data) {
      if (data.isRemoveFileFromSource) {
        this.lessonData.deleteFiles[+data.exerciseNumber-1] = data.path
        return
      }

      this.lessonFiles.delete(`exercise_${data.exerciseNumber}`);
      delete this.exerciseFilesIDS[data.exerciseNumber];
    },
    addExercise(type) {
      this.lessonData.exerciseItems.push({
        exerciseNumber: this.lessonData.exerciseItems.length + 1,
        type,
        name: 'Название задания',
        taskDescription: 'Описание задания'
      });
    },
    setLessonName(e) {
      this.lessonData.name = e.target.innerText;
    },
    removeLesson() {
      store.$patch((state) => {
        const deleteLessonFiles = {
          lessonId: this.lessonNumber-1,
          paths: []
        }
        for (let i = 0; i < this.lessonData.exerciseItems.length; i++) {
          const exercise = this.lessonData.exerciseItems[i]

          if (exercise.path !== undefined && exercise.path !== "") {
            deleteLessonFiles.paths.push(exercise.path)
          }
        }

        if (deleteLessonFiles.paths.length !== 0) {
          state.deleteLessonFileItems.push(deleteLessonFiles)
        }

        if (this.lessonId !== 0 && this.lessonId !== undefined) {
          state.deleteLessonItems.push(this.lessonId)
        }

        state.lessonItems.splice(this.lessonNumber - 1, 1);
      });
    }
  },
};
</script>

<style scoped>
.lesson-component {
  border: 1px solid black;
  border-radius: 10px;
  padding: 5px;
  margin-bottom: 10px;
}

.lesson-component-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.lesson-exercise .title {
  padding-bottom: 10px;
  border-bottom: 1px solid #b3b3b3;
  margin-bottom: 30px;
}

.contenteditable {
  width: max-content;
}

.contenteditable:focus {
  border-radius: 10px;
  padding: 5px;
  background-color: #dfdfdf;
}

.task-panel {
  display: flex;
}

.task-list {
  display: block;
  padding-left: 10px;
}

.task-list > button {
  margin: 0 5px;
}

.exercise-list {
  margin-bottom: 30px;
}

.exercise-item {
  padding-bottom: 5px;
  border-bottom: 1px solid black;
  margin-bottom: 10px;
}
</style>