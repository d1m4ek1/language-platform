import { defineStore } from 'pinia';

export const useListCourses = defineStore('listCourses', {
  state: () => ({
    listOfCourses: []
  }),
  getters: {
    getListOfCourses: (state) => state.listOfCourses
  },
  actions: {
    async sendGetListOfCourses() {
      try {
        await fetch('/teacher/get-list-courses', {
          method: 'GET'
        })
          .then(response => response.json())
          .then(response => {
            this.listOfCourses = response.list || []
          });

        return true
      } catch (err) {
        console.log(err.error);
        return false
      }
    }
  }
});