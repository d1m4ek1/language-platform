<template>
  <div class='task'>
    <div class='task-header'>
      <p>Задание #{{ exerciseNumber }}</p>
      <button class='btn-red' @click='removeExercise'>Удалить задание</button>
    </div>
    <div class='contenteditable sub-title' contenteditable='true' @input='setSubTitle'>
      <h2 v-html='textData.name'></h2>
    </div>
    <div class='contenteditable description-task' contenteditable='true' @input='setTaskDescription'>
      <h3 v-html='textData.taskDescription'></h3>
    </div>
    <div class='content-test'>
      <FileLoader :edit-preview='editPreview'
                  :id-for='`file_loader_lesson_${lessonNumber}_exercise_${exerciseNumber}`' :name-form-data='`exercise_${lessonNumber}`'
                  :type-file='"audio"' @on-upload-file='onUploadFile'></FileLoader>
    </div>
  </div>
</template>

<script>
import FileLoader from '../FileLoader.vue';
import { useCourse } from '../../stores/useCourse';
const store = useCourse();

export default {
  name: "ExerciseAudio",
  components: { FileLoader },
  props: {
    exerciseNumber: Number,
    lessonNumber: Number,
    path: String,
    isEditCourse: Boolean
  },
  data() {
    return {
      exerciseData: {
        type: "audio",
        name: '',
        taskDescription: '',
        path: ''
      },
      textData: {
        name: "Название задания",
        taskDescription: "Описание задания",
      },
      editPreview: ""
    };
  },
  watch: {
    isEditCourse: {
      immediate: true,
      handler(n) {
        if (n) {
          Object.keys(this.exerciseData).forEach(key => {
            this.exerciseData[key] = store.lessonItems[this.lessonNumber - 1].exerciseItems[this.exerciseNumber - 1][key]

            if (key === "path") this.editPreview = this.exerciseData[key]
            if (key === "name") this.textData[key] = this.exerciseData[key]
            if (key === "taskDescription") this.textData[key] = this.exerciseData[key]
          })
        }
      }
    },
    path(n) {
      if (n !== "" && n !== undefined) {
        this.exerciseData.path = n;
      }
    },
    exerciseData: {
      immediate: true,
      deep: true,
      handler(n) {
        clearTimeout(this.timeoutWatchExerciseData);

        this.timeoutWatchExerciseData = setTimeout(() => {
          store.$patch((state) => {
            state.lessonItems[this.lessonNumber - 1].exerciseItems[this.exerciseNumber - 1] =
              { ...state.lessonItems[this.lessonNumber - 1].exerciseItems[this.exerciseNumber - 1], ...n };
          });
        }, 1000);
      }
    }
  },
  methods: {
    onUploadFile(data = {}) {
      if (data.remove) {
        this.$emit('onRemoveFile', {
          exerciseNumber: this.exerciseNumber
        });
      }

      if (!data.success) {
        return;
      }

      this.$emit('onSaveFile', {
        exerciseNumber: this.exerciseNumber,
        file: data.file
      });
    },
    setSubTitle(e) {
      this.exerciseData.name = e.target.innerText;
    },
    setTaskDescription(e) {
      this.exerciseData.taskDescription = e.target.innerText;
    },
    removeExercise() {
      store.lessonItems[this.lessonNumber - 1].exerciseItems.splice(this.exerciseNumber - 1, 1)
    }
  }
};
</script>
