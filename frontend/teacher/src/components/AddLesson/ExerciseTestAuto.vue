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
    <div v-if='exerciseData.task.correctAnswer != null' class='content-test'>
      <ul class='correct-answer-list'>
        <li v-for='(item, idx) in exerciseData.task.correctAnswer' :key='"exercise_" + item.id'>
          <span>{{ idx + 1 }})</span>
          <div class='contenteditable' contenteditable='true' @input='setContentCorrectAnswer($event, item)'>
            <p v-html='textData.task.correctAnswer[idx].phrase'></p>
          </div>
          <button v-if='item.answer' class='btn-correct' @click='item.answer = false'>Верный</button>
          <button v-else class='btn-incorrect' @click='item.answer = true'>Не верный</button>
          <button v-if='exerciseData.task.correctAnswer.length > 1' class='span-plus'
                  @click='removePhraseFromCorrectAnswer(idx)'>-
          </button>
        </li>
        <li>
          <button class='span-plus' @click='addPhraseToCorrectAnswer'>+</button>
        </li>
      </ul>
    </div>
    <div v-if='exerciseData.task.expandBrackets != null' class='content-test'>
      <div class='contenteditable' contenteditable='true'>
        <p v-html='textData.task.expandBrackets.text'></p>
      </div>
      <button class='btn-blue' @click='addInputToExpandBrackets'>Добавить поле</button>
      <ul class='correct-answer-list'>
        <li v-for='(item, idx) in exerciseData.task.expandBrackets.answerItems' :key='item.id'>
          <span>{{ idx + 1 }}</span>
          <div class='contenteditable' contenteditable='true' @input='setAnswerToExpandBrackets($event, item)'>
            Правильный вариант ответа
          </div>
        </li>
      </ul>
    </div>
    <ul class='select-test-list'>
      <li>
        <button class='btn-card-select-test' @click='selectTest("correctAnswer")'>Правильный вариант ответа</button>
      </li>
      <li>
        <button class='btn-card-select-test' @click='selectTest("expandBrackets")'>Раскрыть скобки</button>
      </li>
    </ul>
  </div>
</template>

<script>
import { useCourse } from '../../stores/useCourse';

const store = useCourse();
export default {
  name: 'ExerciseTestAuto',
  props: {
    exerciseNumber: Number,
    lessonNumber: Number,
    isEditCourse: Boolean
  },
  data() {
    return {
      exerciseData: {
        type: 'testAuto',
        name: '',
        taskDescription: '',
        task: {
          type: '',
          correctAnswer: null,
          expandBrackets: null
        }
      },
      textData: {
        name: 'Название задания',
        taskDescription: 'Описание задания',
        task: {
          type: '',
          correctAnswer: null,
          expandBrackets: null
        }
      },
      timeoutWatchExerciseData: null
    };
  },
  watch: {
    isEditCourse: {
      immediate: true,
      handler(n) {
        if (n) {
          Object.keys(this.exerciseData).forEach(key => {
            this.exerciseData[key] = store.lessonItems[this.lessonNumber - 1].exerciseItems[this.exerciseNumber - 1][key];

            switch (key) {
              case 'name':
                this.textData[key] = this.exerciseData[key];
                return;
              case 'taskDescription':
                this.textData[key] = this.exerciseData[key];
                return;
              case 'task':
                this.textData[key] = this.exerciseData[key];
                console.log(this.textData[key]);
                return;
            }
          });
        }
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
    setSubTitle(e) {
      this.exerciseData.name = e.target.innerText;
    },
    setTaskDescription(e) {
      this.exerciseData.taskDescription = e.target.innerText;
    },
    setContentCorrectAnswer(e, item) {
      item.phrase = e.target.innerText;
    },
    setAnswerToExpandBrackets(e, item) {
      item.answer = e.target.innerText;
    },
    selectTest(type) {
      this.exerciseData.task.type = type;
      switch (type) {
        case 'correctAnswer':
          this.exerciseData.task.expandBrackets = null;
          this.exerciseData.task.correctAnswer = [
            {
              id: 0,
              phrase: '',
              answer: false
            }
          ];
          break;
        case 'expandBrackets':
          this.exerciseData.task.correctAnswer = null;
          this.exerciseData.task.expandBrackets = {
            text: '',
            answerItems: []
          };
      }
    },
    addPhraseToCorrectAnswer() {
      this.exerciseData.task.correctAnswer.push({
        id: this.exerciseData.task.correctAnswer.length - 1,
        phrase: '',
        answer: false
      });
    },
    removePhraseFromCorrectAnswer(idx) {
      this.exerciseData.task.correctAnswer.splice(idx, 1);
    },
    addInputToExpandBrackets() {
      this.exerciseData.task.expandBrackets.answerItems.push({
        id: this.exerciseData.task.expandBrackets.answerItems.length - 1,
        answer: ''
      });
    },
    removeExercise() {
      store.lessonItems[this.lessonNumber - 1].exerciseItems.splice(this.exerciseNumber - 1, 1);
    }
  }
};
</script>

<style scoped>
.select-test-list {
  display: flex;
  margin: 20px 0;
}

.select-test-list li:not(:last-child) {
  padding-right: 10px;
}

.correct-answer-list > li {
  display: flex;
  align-items: center;
}

.correct-answer-list > li:not(:last-child) {
  padding-bottom: 5px;
}

.correct-answer-list > li > span {
  padding-right: 10px;
}

.correct-answer-list > li > button {
  margin: 0 5px;
}
</style>