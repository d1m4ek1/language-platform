import { defineStore } from 'pinia';

export const useCourse = defineStore('course', {
  state: () => ({
    courseId: '',
    name: '',
    language: '',
    description: '',
    price: 0,
    free: false,
    hidden: false,
    lessonItems: [],
    deleteLessonFileItems: [],
    deleteLessonItems: [],
  }),
  actions: {
    validateMainInfoCourse() {
      if (this.name === '') return  { error: 'Укажите название курса' };
      if (this.language === '') return  { error: 'Укажите язык курса' };
      if (!this.free && this.price === 0) return  { error: 'Укажите цену курса или установите его бесплатным' };
    },
    async sendSaveCourse(deleteFiles) {
      let error = this.validateMainInfoCourse()
      if (error) throw error

      for (let i = 0; i < this.lessonItems.length; i++) {
        const item = this.lessonItems[i];

        if (item.title === '') throw { error: `Укажите название для урока #${item.lessonNumber}` };
        if (item.exerciseItems.length === 0) throw { error: `Создайте задания для урока #${item.lessonNumber}` };
      }

      return await fetch("/teacher/save-course", {
        method: "POST",
        body: JSON.stringify({
          deleteFiles,
          main: {
            courseId: this.courseId,
            name: this.name,
            language: this.language,
            hidden: this.hidden,
            description: this.description,
            price: this.free ? 0 : this.price
          }
        })
      })
    },
    async sendSaveLessonForSavedCourse(courseId = 0) {
      for (let i = 0; i < this.lessonItems.length; i++) {
        if (this.lessonItems[i].id !== undefined && this.lessonItems[i].id !== 0) {
          this.lessonItems[i].lessonNumber = i+1
        }
      }

      return await fetch(`/teacher/save-lessons`, {
        method: 'POST',
        body: JSON.stringify({
          courseId,
          lessonItems: this.lessonItems,
          deleteLessonFileItems: this.deleteLessonFileItems,
          deleteLessonItems: this.deleteLessonItems
        })
      });
    },
    async sendCreateCourse() {
      let error = this.validateMainInfoCourse()
      if (error) throw error

      for (let i = 0; i < this.lessonItems.length; i++) {
        const item = this.lessonItems[i];

        if (item.title === '') throw { error: `Укажите название для урока #${item.lessonNumber}` };
        if (item.exerciseItems.length === 0) throw { error: `Создайте задания для урока #${item.lessonNumber}` };
      }

      return await fetch('/teacher/create-course', {
        method: 'POST',
        body: JSON.stringify({
          name: this.name,
          language: this.language,
          hidden: this.hidden,
          description: this.description,
          price: this.free ? 0 : this.price
        })
      });
    },
    async sendAddLessonsToCreatedCourse(courseId = 0) {
      return await fetch(`/teacher/add-lessons`, {
        method: 'POST',
        body: JSON.stringify({
          courseId,
          lessonItems: this.lessonItems
        })
      });
    },
    async sendGetCourseByID() {
      try {
        return await fetch(`/teacher/get-course?id=${this.courseId}`)
          .then(response => response.json());
      } catch (err) {
        console.error(err);
      }
    }
  }
});